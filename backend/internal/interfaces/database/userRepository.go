package database

import "github.com/kairo913/tasclock/internal/domain/model"

type UserRepository struct {
	Sqlhandler
}

func (repo *UserRepository) Store(u model.User) (id int64, err error) {
	r, err := repo.Sqlhandler.Execute(
		"INSERT INTO users (firstname, lastname, email, password, salt) VALUES (?, ?, ?, ?, ?)", u.FirstName, u.LastName, u.Email, u.Password, u.Salt,
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

func (repo *UserRepository) Update(u model.User) (err error) {
	_, err = repo.Sqlhandler.Execute(
		"UPDATE users SET firstname = ?, lastname = ?, email = ?, password = ?, salt = ? WHERE id = ?;", u.FirstName, u.LastName, u.Email, u.Password, u.Salt, u.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *UserRepository) Delete(u model.User) (err error) {
	_, err = repo.Sqlhandler.Execute(
		"DELETE FROM users WHERE id = ?;", u.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *UserRepository) FindById(id int) (user model.User, err error) {
	row, err := repo.Sqlhandler.Query("SELECT * FROM users WHERE id = ?;", id)
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByEmail(email string) (user model.User, err error) {
	row, err := repo.Sqlhandler.Query("SELECT * FROM users WHERE email = ?;", email)
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT * FROM users;")
	if err != nil {
		return
	}

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}