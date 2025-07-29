package internal

import (
	"encoding/json"
	"net/http"
	"strings"
)

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := createTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"task_id": task.ID})
}

func HandleAddFile(w http.ResponseWriter, r *http.Request) {
	var data struct {
		TaskID string `json:"task_id"`
		URL    string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&data)
	err := addFileToTask(data.TaskID, data.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "file added"})
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, err := getTask(id)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	resp := map[string]interface{}{
		"status": task.Status,
		"files":  len(task.Files),
	}

	if task.Status == StatusDone {
		resp["download_url"] = task.Results
	}

	if len(task.Error) > 0 {
		resp["errors"] = task.Error
	}
	json.NewEncoder(w).Encode(resp)
}

func HandleMultiStatus(w http.ResponseWriter, r *http.Request) {
	idsParam := r.URL.Query().Get("ids")
	if idsParam == "" {
		http.Error(w, "Нужен парметр ids", http.StatusBadRequest)
		return
	}

	ids := strings.Split(idsParam, ",")
	var results []map[string]interface{}
	for _, id := range ids {
		id = strings.TrimSpace(id)
		task, err := getTask(id)
		if err != nil {
			results = append(results, map[string]interface{}{
				"task_id": id,
				"status":  "not_found",
			})
			continue
		}
		res := map[string]interface{}{
			"task_id": task.ID,
			"status":  task.Status,
			"files":   len(task.Files),
		}
		if task.Status == StatusDone {
			res["download_url"] = task.Results
		}
		if len(task.Error) > 0 {
			res["errors"] = task.Error
		}
		results = append(results, res)
	}
	json.NewEncoder(w).Encode(results)
}
