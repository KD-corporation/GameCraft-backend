package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "gamecraft-backend/prisma/db"
)

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
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



	var otp OtpController

	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid request body",
			Status:  false,
		})
		return
	}
	fmt.Println("user data:", otp)

	otpRecord, err := client.Otp.FindFirst(
		db.Otp.Email.Equals(otp.Email),
		db.Otp.Otp.Equals(otp.Otp),
		db.Otp.ExpiresAt.Gt(time.Now()), 
	).Exec(context.Background())

	if err != nil || otpRecord == nil {
		_, err := client.Otp.FindMany(
			db.Otp.ExpiresAt.Lt(time.Now()),
		).Delete().Exec(context.Background())
		if err != nil {
			fmt.Println("Cleanup error:", err)
		}


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Message: "otp not found",
			Status:  false,
		})
		return
	}



	newUser, err := client.User.CreateOne(
		db.User.Username.Set(otpRecord.Username),
		db.User.FirstName.Set(otpRecord.FirstName),
		db.User.LastName.Set(otpRecord.LastName),
		db.User.Email.Set(otpRecord.Email),
		db.User.Password.Set(otpRecord.Password),
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
		Username: newUser.Username,
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		Email: newUser.Email,
	}

	fmt.Println("new user response:", newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "otp verified successfully",
		Status:  true,
		Data:    responseUserData,
	})

}



