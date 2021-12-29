package handler

import (
	"main/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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

	// username := r.FormValue("email")
	// password := r.FormValue("password")

	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	userOut, err := dbInstance.Signup(user)

	switch {
	case userOut.Email != "" && err == nil:
		if err := dbInstance.AddUser(user); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
		if err := render.Render(w, r, user); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		w.Write([]byte("User created!"))
		return
	case userOut.Email == "":
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
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		http.Redirect(w, r, "/login", 301)
		return
	}
	w.Write([]byte("Hello " + userOut.Email))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
