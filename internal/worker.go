package internal

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func runTask(task *Task) {
	zipPath := "./archives/" + task.ID + ".zip"

	file, err := os.Create(zipPath)

	if err != nil {
		task.Status = StatusError
		task.Error = append(task.Error, "неудалось создать архив")
		finishTask(task)
		return
	}

	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, url := range task.Files {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			task.Error = append(task.Error, "ошибка скачивания: "+url)
			continue
		}
		body := resp.Body

		name := filepath.Base(url)
		w, err := zipWriter.Create(name)
		if err != nil {
			task.Error = append(task.Error, "Ошибка zip: "+url)
			body.Close()
			continue
		}
		io.Copy(w, resp.Body)
		body.Close()
	}

	if len(task.Error) > 0 {
		task.Error = append(task.Error, "Ошибка скачивания")
	} else {
		task.Status = StatusDone
		task.Results = "/archive/" + task.ID + ".zip"
	}
	finishTask(task)
}
