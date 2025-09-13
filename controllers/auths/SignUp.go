package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "gamecraft-backend/prisma/db"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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



	var user SignUpController

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid request body",
			Status:  false,
		})
		return
	}


	newUser, err := client.User.CreateOne(
		db.User.FirstName.Set(user.FirstName),
		db.User.LastName.Set(user.LastName),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(HashPassword(user.Password)),
	).Exec(context.Background())


	if err != nil {
		fmt.Println("Error creating user:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Message:  "failed to create user",
			Status:   false,
			TryLater: "please try again later",
		})
		return
	}

	responseUserData := ResponseUserData{
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		Email: newUser.Email,
	}

	fmt.Println("new user response:", newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "user created successfully",
		Status:  true,
		Data:    responseUserData,
	})
}