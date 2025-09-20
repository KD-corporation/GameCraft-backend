package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gamecraft-backend/controllers/helpers"
	db "gamecraft-backend/prisma/db"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
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
	fmt.Println("user data:", user)


	existing, _ := client.User.FindUnique(
			db.User.Email.Equals(user.Email),
	).Exec(context.Background())

	if existing != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(Response{
			Message:  "user with email already exist",
			Status:   false,
		})
		return
	}

	existing, _ = client.User.FindUnique(
			db.User.Username.Equals(user.Username),
	).Exec(context.Background())

	if existing != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(Response{
			Message:  "user with username already exist",
			Status:   false,
		})
		return
	}



	// go helpers.SendEmail() // Send email in a separate goroutine
	to := []string{user.Email}
	otp := helpers.OptGenerate()
	expiresAt := time.Now().Add(5 * time.Minute)

	if helpers.SendEmail(to, otp) {
		newUser, err := client.Otp.CreateOne(
			db.Otp.Username.Set(user.Username),
			db.Otp.FirstName.Set(user.FirstName),
			db.Otp.LastName.Set(user.LastName),
			db.Otp.Email.Set(user.Email),
			db.Otp.Otp.Set(otp),
			db.Otp.ExpiresAt.Set(expiresAt),
			db.Otp.Password.Set(HashPassword(user.Password)),
		).Exec(context.Background())

		if err != nil {
			fmt.Println("Error creating user:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Message:  "failed to create user and send otp",
				Status:   false,
				TryLater: "please try again later",
			})
			return
		}

		responseUserData := ResponseUserData{
			Username: newUser.Username,
			FirstName: newUser.FirstName,
			LastName: newUser.LastName,
			Email: newUser.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: "user created successfully",
			Status:  true,
			Data:    responseUserData,
		})

		return;

	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{
		Message:  "failed to create user and send otp",
		Status:   false,
		TryLater: "please try again later",
	})
}


	// go helpers.SendEmail() // Send email in a separate goroutine
	// helpers.SendEmail();
	println("OTP for verification:", helpers.GenerateOTP())
}

