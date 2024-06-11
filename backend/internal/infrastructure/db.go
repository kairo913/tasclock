package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kairo913/tasclock/internal/env"
	"github.com/kairo913/tasclock/internal/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
)

type Sqlhandler struct {
	Conn *sql.DB
}

func checkSqlConnect(db *sql.DB, count int) error {
	if err := db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		if count == 0 {
			return fmt.Errorf("db connection failed")
		}
		checkSqlConnect(db, count)
	}
	return nil
}

func NewSqlhandler() *Sqlhandler {
	username := env.GetEnvAsStringOrFallback("MYSQL_USER", "root")
	password := env.GetEnvAsStringOrFallback("MYSQL_PASSWORD", "password")
	port := env.GetEnvAsStringOrFallback("MYSQL_PORT", "3306")
	db_name := env.GetEnvAsStringOrFallback("MYSQL_DATABASE", "tasclock")
	connStr := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8&parseTime=true", username, password, port, db_name)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil
	}

	retry, err := env.GetEnvAsIntOrFallback("CONN_RETRY_COUNT", 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = checkSqlConnect(db, retry)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &Sqlhandler{Conn: db}
}

func CreateTable(handler *Sqlhandler) (err error) {
	_, err = handler.Execute("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, firstname VARCHAR(255) NOT NULL, lastname VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, salt VARCHAR(255) NOT NULL, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);")
	if err != nil {
		return
	}

	_, err = handler.Execute("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, user_id INT DEFAULT NULL, title VARCHAR(255) NOT NULL, is_done TINYINT(1) NOT NULL, description TEXT, deadline DATETIME, elapsed INT NOT NULL, reward FLOAT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL ON UPDATE CASCADE);")
	if err != nil {
		return
	}

	return
}

func (handler *Sqlhandler) Execute(stmt string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}

	result, err := handler.Conn.Exec(stmt, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *Sqlhandler) Query(stmt string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(stmt, args...)
	if err != nil {
		return SqlRow{}, err
	}
	return SqlRow{Row: rows}, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Row *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Row.Next()
}

func (r SqlRow) Close() error {
	return r.Row.Close()
}