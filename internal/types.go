package internal

type TaskStatus string

const (
	StatusWaiting    TaskStatus = "waiting"
	StatusProcessing TaskStatus = "processing"
	StatusDone       TaskStatus = "done"
	StatusError      TaskStatus = "error"
)

type Task struct {
	ID      string
	Files   []string
	Status  TaskStatus
	Results string
	Error   []string
}
