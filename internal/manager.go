package internal

import (
	"errors"
	"log"
	"path"
	"strings"
	"sync"
)

var (
	taskStorage = make(map[string]*Task)
	taskMutex   sync.Mutex
	activeTasks = 0
)

const MaxActiveTasks = 3

func createTask() (*Task, error) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	if activeTasks >= MaxActiveTasks {
		return nil, errors.New("сервер занят, подождите загрузки файлов")
	}

	id := generateID()
	task := &Task{
		ID:     id,
		Files:  []string{},
		Status: StatusWaiting,
	}

	taskStorage[id] = task
	activeTasks++

	return task, nil
}

func getTask(id string) (*Task, error) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	task, ok := taskStorage[id]

	if !ok {
		return nil, errors.New("задача не найдена")
	}

	return task, nil
}

func addFileToTask(id, url string) error {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	task, ok := taskStorage[id]
	if !ok {
		return errors.New("задача не найдена")
	}

	if task.Status != StatusWaiting {
		return errors.New("задача уже в работе")
	}

	if len(task.Files) >= 3 {
		return errors.New("в задаче уже максимум файлов")
	}

	ext := strings.ToLower(path.Ext(url))
	if ext != ".pdf" && ext != ".jpeg" && ext != ".jpg" {
		return errors.New("разрешены только .pdf, .jpeg, .jpg")
	}

	task.Files = append(task.Files, url)

	if len(task.Files) == 3 {
		task.Status = StatusProcessing
		go runTask(task)
	}

	return nil
}

func finishTask(task *Task) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	activeTasks--
	log.Printf("Задача завершена: %s, статус: %s", task.ID, task.Status)
}
