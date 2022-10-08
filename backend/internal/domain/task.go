package domain

const (
	TaskTypeGeneral = iota
	TaskTypePersonal
	TaskTypeCommand
)

const (
	TaskPriorityLow = iota
	TaskPriorityMiddle
	TaskPriorityHigh
	TaskPriorityEmergency
)

const (
	TaskStatusWaiting = iota
	TaskStatusProgress
	TaskStatusDone
)

type Task struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Body        string  `json:"body"`
	Revenue     float32 `json:"revenue"`
	Type        int     `json:"type"`
	Priority    int     `json:"priority"`
	Status      int     `json:"status"`
	Thumbnail   string  `json:"thumbnail"`
}
