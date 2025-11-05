package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, desc string) Task {
	return Task{
		Title:       title,
		Description: desc,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	compTime := time.Now()
	t.Completed = true
	t.CompletedAt = &compTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
