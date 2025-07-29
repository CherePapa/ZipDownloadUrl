package main

import (
	"log"
	"net/http"

	"github.com/CherePapa/ZipDownloadUrl/internal"
)

func main() {
	internal.InitArchiveFolder()

	http.HandleFunc("/task/create", internal.HandleCreateTask)
	http.HandleFunc("/task/add", internal.HandleAddFile)
	http.HandleFunc("/task/statusFile", internal.HandleStatus)
	http.HandleFunc("/task/statusesFile", internal.HandleMultiStatus)

	http.Handle("/archive/", http.StripPrefix("/archive/", http.FileServer(http.Dir("./archives"))))

	log.Println("[ZipDownloadUrl] Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
