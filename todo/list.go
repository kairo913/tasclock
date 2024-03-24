package todo

import "log"

type List struct {
	ID    int64  `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

func (td *Todo) NewList(title string) (*List, error) {
	list := &List{
		Title: title,
	}

	stmt, err := td.db.Prepare("INSERT INTO lists(title) VALUES (?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(list.Title)

	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		return nil, err
	}

	list.ID = id

	return list, nil
}

func (td *Todo) RemoveList(id int64) error {
	stmt, err := td.db.Prepare("DELETE FROM lists WHERE id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (td *Todo) UpdateList(list *List) error {
	stmt, err := td.db.Prepare("UPDATE tasks SET title = ? WHERE id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(list.Title, list.ID)

	if err != nil {
		return err
	}

	return nil
}