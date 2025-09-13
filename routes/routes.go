package routes

import (
	"gamecraft-backend/controllers/auths"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/signup", auths.SignUp)
}