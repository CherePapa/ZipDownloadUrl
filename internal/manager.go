package internal

import (
	"errors"
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
		return nil, errors.New("сервер занят, подождите загрузки файлов!")
	}

	id := generateID()
	task := &Task{
		ID:     id,
		Files:  []string{},
		Status: StatusWaiting,
	}

	return task, nil
}

func getTask(id string) (*Task, error) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	task, ok := taskStorage[id]

	if !ok {
		return nil, errors.New("Задача не найдена")
	}

	return task, nil
}

func addFileToTask(id, url string) error {
	taskMutex.Lock()
	taskMutex.Unlock()

	task, ok := taskStorage[id]

	if !ok {
		return errors.New("Задача не найдена!")
	}

	if task.Status != StatusWaiting {
		return errors.New("Здача уже в работе")
	}

	if len(task.Files) >= 3 {
		return errors.New("В задаче уже максимум файлов!")
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
	taskMutex.Unlock()
	activeTasks--
}
