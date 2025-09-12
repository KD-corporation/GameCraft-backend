package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	// PostgreSQL connection string
	dsn := "postgres://gagan:gagan@localhost:5432/mydb"

	var err error
	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Simple test route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go backend is running 🚀"))
	})

	// Example: fetch users
	http.HandleFunc("/users", getUsers)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
	}
}
