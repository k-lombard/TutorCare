package handler

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"main/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/nu7hatch/gouuid"
)

func signup(router chi.Router) {
	router.Post("/", signupPage)
}

func login(router chi.Router) {
	router.Post("/", loginPage)
}

var client *redis.Client

func init() {
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	// 	dsn = "localhost:6379"
	// }
	url, err2 := redis.ParseURL("redis://redis:6379")
	if err2 != nil {
		panic(err2)
	}
	client = redis.NewClient(url)
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "signup.html")
		return
	}
	user := &models.User{}

	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	isUnique := dbInstance.Signup(user)

	switch {
	case user.Email != "" && isUnique == true:
		userOut1, err := dbInstance.AddUser(user)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
		if err := render.Render(w, r, &userOut1); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		http.Redirect(w, r, "/login", 301)
	case isUnique == false:
		http.Error(w, "Server error, unable to create your account. User with email already exists", 500)
		return
	default:
		http.Redirect(w, r, "/", 301)
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "login.html")
		return
	}
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	userOut, isMatch := dbInstance.Login(user)
	if isMatch == false {
		render.Render(w, r, ErrNotFound)
		return
	}
	if err := render.Render(w, r, &userOut); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		http.Redirect(w, r, "/login", 301)
		return
	}
	h := fnv.New64a()
	h.Write([]byte(userOut.UserID.String()))
	summedUserID := h.Sum64()
	ts, err := CreateToken(summedUserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	saveErr := AuthFunc(summedUserID, ts)
	if saveErr != nil {
		render.Render(w, r, ErrorRenderer(err))
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	json.NewEncoder(w).Encode(tokens)
}

func CreateToken(userid uint64) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	newUuid, err3 := uuid.NewV4()
	if err3 != nil {
		fmt.Println("error creating v4 uuid: ", err3)
		return td, err3
	}
	td.AccessUuid = newUuid.String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshUuid, err2 := uuid.NewV4()
	if err2 != nil {
		fmt.Println("error creating v4 uuid: ", err2)
		return td, err2
	}
	td.RefreshUuid = refreshUuid.String()

	var err error
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_id"] = userid
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return td, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func AuthFunc(userid uint64, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(client.Context(), td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(client.Context(), td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
