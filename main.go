package main

import (
	"fmt"
	"gamecraft-backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	apiMux := http.NewServeMux()
	getMux := http.NewServeMux()

	// Register different routes in different muxes
	routes.RegisterRouter(apiMux)    // for POST routes like /signup
	routes.RegisterRouterGet(getMux) // for GET routes like /getuser

	// Mount them under prefixes
	http.Handle("/api/", http.StripPrefix("/api", apiMux))
	http.Handle("/getapis/", http.StripPrefix("/getapis", getMux))


	

	if err := godotenv.Load(); err != nil {
		log.Println(" No .env file found, using system environment variables")
	}
	// Ensure DATABASE_URL is set
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal(" DATABASE_URL environment variable is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("❓ Unhandled request: %s %s", r.Method, r.URL.Path)
		http.NotFound(w, r)
	})

	fmt.Println("🚀 Server running at http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", nil))

}
