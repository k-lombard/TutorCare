package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func authentication(router chi.Router) {
	router.Post("/", authFunc)
}

func authFunc(w http.ResponseWriter, r *http.Request) {

}
