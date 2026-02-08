package task

var List []Task

type Task struct {
	ID   int
	Text string
	Done bool
}

type CreateTaskRequest struct {
	Text string `json:"text"`
}

type UpdateTaskRequest struct {
	ID   int
	Text string
}
