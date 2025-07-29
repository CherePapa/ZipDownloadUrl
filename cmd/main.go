package main

import (
	"log"
	"net/http"

	"github.com/CherePapa/ZipDownloadUrl/internal"
	"github.com/gorilla/mux"
)

func main() {
	internal.InitArchiveFolder()

	r := mux.NewRouter()

	r.HandleFunc("/task/create", internal.HandleCreateTask).Methods("POST")
	r.HandleFunc("/task/add", internal.HandleAddFile).Methods("POST")
	r.HandleFunc("/task/statusFile/{task_id}", internal.HandleStatus).Methods("GET")
	r.HandleFunc("/tasks", internal.HandleMultiStatus).Methods("GET")

	r.PathPrefix("/archive/").Handler(
		http.StripPrefix("/archive/", http.FileServer(http.Dir("./archives"))),
	)

	log.Println("[ZipDownloadUrl] Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
