package handler

import (
	"bytes"
	"fmt"
	"main/database"

	"github.com/gin-gonic/gin"
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

func RouteHandler(db database.Database) *gin.Engine {
	dbInstance = db
	router.Use(ginBodyLogMiddleware)
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(CORSMiddleware())
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
	applications := api.Group("/applications")
	r.applications(applications)
	return r.router
}
