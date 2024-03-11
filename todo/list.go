package todo

type List struct {
	ID    int64  `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

func (td *Todo) NewList(title string) (*List, error) {
	list := &List{
		Title: title,
	}

	const sqlStr = `INSERT INTO lists(title) VALUES (?);`

	r, err := td.db.Exec(sqlStr, list.Title)
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
	const sqlStr = `DELETE FROM lists WHERE id = ?;`
	_, err := td.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (td *Todo) UpdateList(list *List) error {
	const sqlStr = `UPDATE tasks SET title = ? WHERE id = ?`
	_, err := td.db.Exec(sqlStr, list.Title, list.ID)
	if err != nil {
		return err
	}
	return nil
}