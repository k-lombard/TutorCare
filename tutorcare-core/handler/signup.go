package handler

import (
	"hash/fnv"
	"main/models"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v4"
)

func signup(router chi.Router) {
	router.Post("/", signupPage)
}

func login(router chi.Router) {
	router.Post("/", loginPage)
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
	token, err := CreateToken(h.Sum64())
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	w.Write([]byte(token))
}

func CreateToken(userid uint64) (string, error) {
	var err error
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_id"] = userid
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
