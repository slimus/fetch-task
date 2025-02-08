package main

import (
	"log"
	"net/http"
	"os"

	"github.com/slimus/fetch-task/internal/api/receipts"
	"github.com/slimus/fetch-task/internal/storage/inmemory"
)

func main() {
	port := os.Getenv("PORT")

	storage := inmemory.NewDB()
	app := receipts.NewApp(storage)

	//TODO: we could use /api/v1/ to version the API (improvements)
	http.HandleFunc("/receipts/process", app.HandleRecieptsProcess)
	http.HandleFunc("/receipts/{id}/points", app.HandleRecieptsPoints)

	log.Printf("Server running on port %s ...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	//TODO: we have to handle OS signals to help us gracefully shutdown the server
}
