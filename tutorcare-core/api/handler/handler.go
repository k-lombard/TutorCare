package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/database"
	"main/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

var dbInstance database.Database
var (
	router = gin.Default()
)

type routes struct {
	router *gin.Engine
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 400 {
		fmt.Println("Response body: " + blw.body.String())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RouteHandler(db database.Database, m *melody.Melody) *gin.Engine {
	dbInstance = db
	router.Use(ginBodyLogMiddleware)
	r := routes{
		router: gin.Default(),
	}
	// r.router.Use(CORSMiddleware())
	r.router.Use(cors.Default())
	api := r.router.Group("/api", CORSMiddleware())
	geolocationpositions := api.Group("/geolocationpositions")
	r.geolocationpositions(geolocationpositions)
	users := api.Group("/users")
	r.users(users)
	signup := api.Group("/signup")
	r.signup(signup)
	login := api.Group("/login")
	r.login(login)
	logout := api.Group("/logout", TokenAuthMiddleware())
	r.logout(logout)
	token := api.Group("/token")
	r.token(token)
	profile := api.Group("/profile", TokenAuthMiddleware())
	r.profile(profile)
	posts := api.Group("/posts")
	r.posts(posts)
	postCodes := api.Group("/post_codes")
	r.post_codes(postCodes)
	applications := api.Group("/applications", TokenAuthMiddleware())
	r.applications(applications)
	chatrooms := api.Group("/chatrooms", TokenAuthMiddleware())
	r.chatrooms(chatrooms)
	messages := api.Group("/messages", TokenAuthMiddleware())
	r.messages(messages)
	api.GET("/:chatroomid/ws", func(c *gin.Context) {
		// websocketToken := models.WebsocketToken{}
		// if err := render.Bind(c.Request, &websocketToken); err != nil {
		// 	c.JSON(http.StatusBadRequest, "Error: Bad request")
		// 	return
		// }
		// useridOut, err := Client.Get(ctx, websocketToken.Token).Result()
		// if err != nil {
		// 	c.JSON(http.StatusForbidden, "Error: Forbidden; please re-login.")
		// }
		// if useridOut == websocketToken.UserID.String() {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		val, found := s.Get("sent")
		fmt.Println(found)
		if val != nil && found == true {
			m.BroadcastFilter(msg, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
		} else {
			var tokenObj models.WebsocketToken
			if err2 := json.Unmarshal(msg, &tokenObj); err2 != nil {
				s.CloseWithMsg(melody.FormatCloseMessage(403, "Error: please re-login."))
			}
			fmt.Println(tokenObj.UserID, tokenObj.Token)
			useridOut, err4 := Client.Get(ctx, tokenObj.Token).Result()
			if err4 != nil {
				s.CloseWithMsg(melody.FormatCloseMessage(403, "Error: please re-login."))
			}
			fmt.Println(useridOut)
			if useridOut == tokenObj.UserID {
				fmt.Println("success")
				s.Set("sent", true)
			} else {
				fmt.Println("equality error")
				s.CloseWithMsg(melody.FormatCloseMessage(403, "Error: please re-login."))
			}
		}
	})
	return r.router
}
