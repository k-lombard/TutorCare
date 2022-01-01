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

func RouteHandler(db database.Database) *gin.Engine {
	dbInstance = db
	router.Use(ginBodyLogMiddleware)
	r := routes{
		router: gin.Default(),
	}
	users := r.router.Group("/users")
	r.users(users)
	signup := r.router.Group("/signup")
	r.signup(signup)
	login := r.router.Group("/login")
	r.login(login)
	logout := r.router.Group("/logout")
	r.logout(logout)
	token := r.router.Group("/token")
	r.token(token)
	return r.router
}
