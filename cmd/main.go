package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CherePapa/ZipDownloadUrl/internal"
	"github.com/gorilla/mux"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "create":
			internal.CreateCommandTerminal()
			return
		case "add":
			if len(os.Args) != 4 {
				fmt.Println("ОШИБКА: Проверьте правильность написания task_id и url")
				return
			}
			internal.AddFileCommandTerminal(os.Args[2], os.Args[3])
			return
		case "status":
			if len(os.Args) != 3 {
				fmt.Println("ОШИБКА: Проверьте правильность написания task_id")
				return
			}
			internal.CheckTaskStatusCommandTerminal(os.Args[2])
			return
		case "status-all":
			internal.CheckAllTaskStatusCommand()
			return
		default:
			fmt.Println("ОШИБКА: Неизвестная команда")
			return
		}

	}
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
