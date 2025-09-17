package auths

import (
	"encoding/json"
	"net/http"
	"time"


)




func Logout(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		w.Header().Set("Context-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Method not allowed",
			Status: false,
		})
		return
	}




	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,               // JS can’t read it
		Secure:   false,              // set true if using HTTPS
		Path:     "/",
		Expires:  time.Unix(0,0),
		MaxAge: -1,
	})


	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "logout successful",
		Status:  true,
	})
}