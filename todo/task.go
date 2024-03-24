package todo

import (
	"log"
	"time"
)

type Task struct {
	ID          int64  `json:"id"`
	ListID      int64  `json:"list_id"`
	Title       string `json:"title"`
	IsDone      bool   `json:"is_done"`
	Starred     bool   `json:"starred"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Reward      int64  `json:"reward"`
	Elapsed     int64  `json:"elapsed"`
	CreatedAt   string `json:"created_at"`
}

func (td *Todo) NewTask(list_id int64, title string, starred bool, description string, deadline string, reward int64) (*Task, error) {
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

	stmt, err := td.db.Prepare("INSERT INTO tasks(list_id,title,starred,description,deadline,reward,created_at) VALUES (?,?,?,?,?,?,?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(task.ListID, task.Title, task.Starred, task.Description, task.Deadline, task.Reward, task.CreatedAt)

	if err != nil {
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
	stmt, err := td.db.Prepare("DELETE FROM tasks WHERE id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (td *Todo) UpdateTask(task *Task) error {
	stmt, err := td.db.Prepare("UPDATE tasks SET list_id = ?, title = ?, is_done = ?, starred = ?, description = ?, deadline = ?, reward = ?, elapsed = ? WHERE id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.ListID, task.Title, task.IsDone, task.Starred, task.Description, task.Deadline, task.Reward, task.Elapsed, task.ID)

	if err != nil {
		return err
	}

	return nil
}
