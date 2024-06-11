package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	IsDone      bool      `json:"is_done"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Elapsed     int       `json:"elapsed"`
	Reward      float64   `json:"reward"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task
