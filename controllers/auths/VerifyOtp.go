package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "gamecraft-backend/prisma/db"
)

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer client.Prisma.Disconnect()

	var otp OtpController
	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	email := strings.TrimSpace(strings.ToLower(otp.Email))
	otpCode := strings.TrimSpace(otp.Otp)
	now := time.Now().UTC() // ensure consistent timezone

	fmt.Println("Verifying OTP for email:", email, "with code:", otpCode, "at", now)

	otpRecord, err := client.Otp.FindFirst(
		db.Otp.Email.Equals(email),
		db.Otp.Otp.Equals(otpCode),
		db.Otp.ExpiresAt.Gt(time.Now().UTC()),
	).Exec(context.Background())

	if err != nil {
		fmt.Println("Error fetching OTP record:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if otpRecord == nil {
		// Optional cleanup of expired OTPs
		if _, err := client.Otp.FindMany(db.Otp.ExpiresAt.Lt(now)).Delete().Exec(context.Background()); err != nil {
			fmt.Println("Cleanup error:", err)
		}
		http.Error(w, "OTP not found or expired", http.StatusUnauthorized)
		return
	}

	// Create new user
	newUser, err := client.User.CreateOne(
		db.User.FirstName.Set(otpRecord.FirstName),
		db.User.LastName.Set(otpRecord.LastName),
		db.User.Email.Set(otpRecord.Email),
		db.User.Password.Set(otpRecord.Password),
		db.User.Username.Set(otpRecord.Username),
	).Exec(context.Background())

	if err != nil {
		fmt.Println("Error creating user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	fmt.Println("New user created:", newUser.Username)

	responseUserData := ResponseUserData{
		Username:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message: "OTP verified successfully",
		Status:  true,
		Data:    responseUserData,
	})
}
