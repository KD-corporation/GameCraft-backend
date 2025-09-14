package routes

import (
	"fmt"
	"gamecraft-backend/controllers/auths"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/signup", auths.SignUp)
	mux.HandleFunc("/login", auths.Login)
	mux.HandleFunc("/logout", auths.Logout)
}

// only get request are Alloweed to this Function
func RegisterRouterGet(mux *http.ServeMux) {
	//here you can map the get requests to their respective handler functions

	mux.HandleFunc("/getuser", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "send some data without writing any function"}`))
	})

	fmt.Println("GET /getuser route registered")
}

// similar you can do for the put and delete reuest
