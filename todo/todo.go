package todo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	db  *sql.DB
}

func NewTodo() *Todo {
	return &Todo{}
}

func (td *Todo) Startup() error {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	td.db = db

	if err := td.initDB(); err != nil {
		return err
	}

	return nil
}

func (td *Todo) Shutdown() {
	td.db.Close()
}

func (td *Todo) initDB() error {
	sqlStrs := []string{
		`CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY,
			list_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			is_done BOOLEAN NOT NULL DEFAULT 0,
			starred BOOLEAN NOT NULL,
			description TEXT NOT NULL,
			deadline TEXT NOT NULL,
			reward INTEGER NOT NULL,
			elapsed INTEGER NOT NULL DEFAULT 0,
			created_at TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS lists(
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL
		);`,
	}

	for _, sqlStr := range sqlStrs {
		if _, err := td.db.Exec(sqlStr); err != nil {
			return err
		}
	}

	return nil
}

func (td *Todo) Tasks(limit int64) ([]*Task, error) {
	stmt, err := td.db.Prepare("SELECT id, list_id, is_done, starred, title, description, deadline, reward, elapsed, created_at FROM tasks LIMIT ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ts []*Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.ListID, &t.IsDone, &t.Starred, &t.Title, &t.Description, &t.Deadline, &t.Reward, &t.Elapsed, &t.CreatedAt)

		if err != nil {
			return nil, err
		}

		ts = append(ts, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}

func (td *Todo) Lists(limit int64) ([]*List, error) {
	stmt, err := td.db.Prepare("SELECT id, title FROM lists LIMIT ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ls []*List
	for rows.Next() {
		var l List
		err := rows.Scan(&l.ID, &l.Title)

		if err != nil {
			return nil, err
		}

		ls = append(ls, &l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ls, nil
}