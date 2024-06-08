package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kairo913/tasclock/internal/env"
)

type SqlHandler struct {
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

func NewSqlHandler() *SqlHandler {
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

	return &SqlHandler{Conn: db}
}

func (handler *SqlHandler) Execute(stmt string, args ...interface{}) (sql.Result, error) {
	res := SqlResult{}

	result, err := handler.Conn.Exec(stmt, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(stmt string, args ...interface{}) (SqlRow, error) {
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