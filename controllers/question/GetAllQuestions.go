package question


import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "gamecraft-backend/prisma/db"
)

func GetAllQustion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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





	questions, err := client.Question.FindMany().
		OrderBy(
			db.Question.ID.Order(db.SortOrderAsc),
		).
		Exec(context.Background())

	if err != nil {
		fmt.Println("Error fetching questions: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Message:  "failed to fetch questions",
			Status:   false,
			TryLater: "please try again later",
		})
		return
	}

	// Map into a lightweight DTO
	var result []QuestionResponse
	for _, q := range questions {
		result = append(result, QuestionResponse{
			Id:          q.ID,
			Title:       q.Title,
			Description: q.Description,
		})
	}



	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "questioned fetched successfully",
		Status:  true,
		Data:    result,
	})

}

