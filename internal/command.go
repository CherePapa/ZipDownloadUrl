package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func CreateCommandTerminal() {
	resp, err := http.Post("http://localhost:8080/task/create", "application/json", bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal("Ошибка при использовании команды и создания задачи", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Ошибка сервера", err)
	}

	var data TaskCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal("Ошибка при использовании команды и создания задачи", err)
	}

	fmt.Println("Задача создана! task_id:", data.TaskID)
}

func AddFileCommandTerminal(taskID, url string) {
	data := AddFileRequest{
		TaskID: taskID,
		URL:    url,
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Ошибка запроса", err)
	}

	resp, err := http.Post("http://localhost:8080/task/add", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Ошибка при запросе", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := new(strings.Builder)
		_, _ = io.Copy(respBody, resp.Body)
		log.Fatal("Ошибка сервера: %s\n%s", resp.Status, respBody)
	}

	results := GenericMessageRespose{
		Message: "файл добавлен",
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Fatal("Ошибка при запросе", err)
	}

	fmt.Println(results.Message)
}

func CheckTaskStatusCommandTerminal(taskID string) {
	url := fmt.Sprintf("http://localhost:8080/task/statusFile/%s", taskID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка при запросе", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := new(strings.Builder)
		_, _ = io.Copy(respBody, resp.Body)
		log.Fatal("Ошибка сервера: %s\n%s", resp.Status, respBody)
	}

	var data TaskStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal("Ошибка при запросе", err)
	}
	fmt.Printf("Задача %s: \n", data.TaskID)
	fmt.Println("  Статус:     ", data.Status)
	fmt.Println("  Файлов:     ", data.Files)
	if data.DownloadURL != "" {
		fmt.Println("  Ссылка:     ", data.DownloadURL)
	}
	if len(data.Errors) > 0 {
		fmt.Println("  Ошибки:     ", data.Errors)
	}
}

func CheckAllTaskStatusCommand() {
	resp, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		log.Fatal("Ошибка при запросе", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := new(strings.Builder)
		_, _ = io.Copy(respBody, resp.Body)
		log.Fatal("Ошибка сервера: %s\n%s", resp.Status, respBody)
	}

	var list []TaskStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		log.Fatal("Ошибка при запросе", err)
	}

	if len(list) == 0 {
		fmt.Println("Нет активынх задач")
		return
	}

	fmt.Println("Список задач:")
	for _, task := range list {
		fmt.Printf("- [%s] %s (%d файлов)\n", task.TaskID, task.Status, task.Files)
		if task.DownloadURL != "" {
			fmt.Printf("  Ссылка: %s\n", task.DownloadURL)
		}
		if len(task.Errors) > 0 {
			fmt.Printf("  Ошибки: %s\n", task.Errors)
		}
	}
}
