package task

type Task struct {
	ID          int    `json:"id" yaml:"id"`
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Status      string `json:"status" yaml:"status"`
	CreatedAt   string `json:"created_at" yaml:"created_at"`
	UpdatedAt   string `json:"updated_at" yaml:"updated_at"`
}
