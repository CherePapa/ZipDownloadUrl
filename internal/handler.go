package internal

import (
	"encoding/json"
	"net/http"
	"strings"
)

type TaskStatusResponse struct {
	TaskID      string   `json:"task_id"`
	Status      string   `json:"status"`
	Files       int      `json:"files,omitempty"`
	DownloadURL string   `json:"download_url,omitempty"`
	Errors      []string `json:"errors,omitempty"`
}

type TaskCreateResponse struct {
	TaskID string `json:"task_id"`
}

type AddFileRequest struct {
	TaskID string `json:"task_id"`
	URL    string `json:"url"`
}

type GenericMessageRespose struct {
	Message string `json:"message"`
}

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := createTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(TaskCreateResponse{TaskID: task.ID})
}

func HandleAddFile(w http.ResponseWriter, r *http.Request) {
	var data AddFileRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	if err := addFileToTask(data.TaskID, data.URL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := GenericMessageRespose{Message: "file added"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, err := getTask(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := TaskStatusResponse{
		TaskID: id,
		Status: string(task.Status),
		Files:  len(task.Files),
	}

	if task.Status == StatusDone {
		resp.DownloadURL = task.Results
	}

	if len(task.Error) > 0 {
		resp.Errors = task.Error
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HandleMultiStatus(w http.ResponseWriter, r *http.Request) {
	idsParam := r.URL.Query().Get("ids")
	if idsParam == "" {
		http.Error(w, "Нужен парметр ids", http.StatusBadRequest)
		return
	}

	ids := strings.Split(idsParam, ",")
	var results []TaskStatusResponse
	for _, id := range ids {
		id = strings.TrimSpace(id)
		task, err := getTask(id)
		if err != nil {
			results = append(results, TaskStatusResponse{
				TaskID: id,
				Status: "not_found",
			})
			continue
		}
		res := TaskStatusResponse{
			TaskID: id,
			Status: string(task.Status),
			Files:  len(task.Files),
		}
		if task.Status == StatusDone {
			res.DownloadURL = task.Results
		}
		if len(task.Error) > 0 {
			res.Errors = task.Error
		}
		results = append(results, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
