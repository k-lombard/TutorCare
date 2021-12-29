package handler

import (
	"context"
	"fmt"
	"main/database"
	"net/http"

	"main/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var userIDKey = "userID"

func users(router chi.Router) {
	router.Get("/", getAllUsers)
	router.Post("/", addUser)
	router.Route("/{userId}", func(router chi.Router) {
		router.Use(UserContext)
		router.Get("/", getUserById)
		router.Put("/", updateUser)
		router.Delete("/", deleteUser)
	})
}
func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		if userId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("user ID is required")))
			return
		}
		id := uuid.MustParse(userId)
		// if err != nil {
		// 	render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid user ID")))
		// }
		ctx := context.WithValue(r.Context(), userIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	userOut, err := dbInstance.AddUser(user)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err2 := render.Render(w, r, user); err2 != nil {
		render.Render(w, r, ServerErrorRenderer(err2))
		return
	}
	w.Write([]byte(userOut.UserID.String()))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(uuid.UUID)
	user, err := dbInstance.GetUserById(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(uuid.UUID)
	err := dbInstance.DeleteUser(userId)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(uuid.UUID)
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	user, err := dbInstance.UpdateUser(userId, userData)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
