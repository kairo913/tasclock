package model

type Task struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	IsDone      bool   `json:"is_done"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Elapsed     int    `json:"elapsed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Tasks []Task
