package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Import environment variables
	godotenv.Load()

	port := os.Getenv("PORT")

	//Configure Server
	mux := http.NewServeMux()
	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
