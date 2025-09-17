package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"


	"gamecraft-backend/middlewares"
	db "gamecraft-backend/prisma/db"

	"github.com/golang-jwt/jwt/v5"
)




func GetUser(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		w.Header().Set("Context-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Method not allowed",
			Status: false,
		})
		return
	}




	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		http.Error(w, "failed to connect to server", http.StatusInternalServerError)
		return
	}
	defer client.Prisma.Disconnect()


	claims, ok := r.Context().Value(middlewares.UserKey).(jwt.MapClaims)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	str := fmt.Sprintf("%v", claims["user_id"])
	num, _ := strconv.Atoi(str)
	existing, err := client.User.FindUnique(
		db.User.ID.Equals(num),
	).Exec(context.Background())



	if err != nil || existing == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid credentials",
			Status:  false,
		})
		return
	}

	response := ResponseUserData{
		Username: existing.Username,
		FirstName: existing.FirstName,
		LastName: existing.LastName,
		Email: existing.Email,
	}


	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "user fetched successfully",
		Status:  true,
		Data: response,
	})
}