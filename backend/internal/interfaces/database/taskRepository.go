package database

import "github.com/kairo913/tasclock/internal/domain/model"

type TaskRepository struct {
	Sqlhandler
}

func (repo *TaskRepository) Store(t model.Task) (id int64, err error) {
	r, err := repo.Sqlhandler.Execute(
		"INSERT INTO tasks (user_id, title, is_done, description, deadline, elapsed) VALUES (?, ?, ?, ?, ?, ?);", t.UserId, t.Title, t.IsDone, t.Description, t.Deadline, t.Elapsed,
	)

	if err != nil {
		return
	}

	id, err = r.LastInsertId()

	if err != nil {
		return -1, err
	}

	return
}

func (repo *TaskRepository) Update(t model.Task) (err error) {
	_, err = repo.Sqlhandler.Execute(
		"UPDATE tasks SET user_id = ?, title = ?, is_done = ?, description = ?, deadline = ?, elapsed = ? WHERE id = ?;", t.UserId, t.Title, t.IsDone, t.Description, t.Deadline, t.Elapsed, t.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *TaskRepository) Delete(t model.Task) (err error) {
	_, err = repo.Sqlhandler.Execute(
		"DELETE FROM tasks WHERE id = ?;", t.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *TaskRepository) FindById(id int) (task model.Task, err error) {
	row, err := repo.Sqlhandler.Query("SELECT * FROM tasks WHERE id = ?;", id)
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&task.Id, &task.UserId, &task.Title, &task.IsDone, &task.Description, &task.Deadline, &task.Elapsed, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return
	}

	return
}

func (repo *TaskRepository) FindByUserId(userId int) (tasks model.Tasks, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT * FROM tasks WHERE user_id = ?;", userId)
	if err != nil {
		return
	}

	for rows.Next() {
		var task model.Task
		if err = rows.Scan(&task.Id, &task.UserId, &task.Title, &task.IsDone, &task.Description, &task.Deadline, &task.Elapsed, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return
		}
		tasks = append(tasks, task)
	}

	return
}

func (repo *TaskRepository) FindAll() (tasks model.Tasks, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT * FROM tasks;")
	if err != nil {
		return
	}

	for rows.Next() {
		var task model.Task
		if err = rows.Scan(&task.Id, &task.UserId, &task.Title, &task.IsDone, &task.Description, &task.Deadline, &task.Elapsed, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return
		}
		tasks = append(tasks, task)
	}

	return
}