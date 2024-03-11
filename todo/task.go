package todo

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int64  `json:"id"`          // Task id
	ListID      int64  `json:"list_id"`     // List id
	Title       string `json:"title"`       // Task title
	IsDone      bool   `json:"is_done"`     // Is task completed
	Starred     bool   `json:"starred"`     // Is task starred
	Description string `json:"description"` // Task description
	Deadline    string `json:"deadline"`    // Task deadline
	Reward      int64  `json:"reward"`      // Task reward
	Elapsed     int64  `json:"elapsed"`     // Elapsed time of task
	CreatedAt   string `json:"created_at"`  // Time task created
}

func (td *Todo) NewTask(list_id int64, title string, starred bool, description string, deadline string, reward int64) (*Task, error) {
	fmt.Println(starred)
	task := &Task{
		ListID:      list_id,
		Title:       title,
		IsDone:      false,
		Starred:     starred,
		Description: description,
		Deadline:    deadline,
		Reward:      reward,
		Elapsed:     0,
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	const sqlStr = `INSERT INTO tasks(list_id,title,starred,description,deadline,reward,created_at) VALUES (?,?,?,?,?,?,?);`

	r, err := td.db.Exec(sqlStr, task.ListID, task.Title, task.Starred, task.Description, task.Deadline, task.Reward, task.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	task.ID = id

	return task, nil
}

func (td *Todo) RemoveTask(id int64) error {
	const sqlStr = `DELETE FROM tasks WHERE id = ?;`

	_, err := td.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (td *Todo) UpdateTask(task *Task) error {
	const sqlStr = `UPDATE tasks SET list_id = ?, title = ?, is_done = ?, starred = ?, description = ?, deadline = ?, reward = ?, elapsed = ? WHERE id = ?`
	_, err := td.db.Exec(sqlStr, task.ListID, task.Title, task.IsDone, task.Starred, task.Description, task.Deadline, task.Reward, task.Elapsed, task.ID)
	if err != nil {
		return err
	}
	return nil
}
