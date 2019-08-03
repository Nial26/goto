package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBEnv struct{
	Db *sql.DB
}


func InitDB(dataStoreName string) (*DBEnv, error) {
	db, err := sql.Open("mysql", dataStoreName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	env := &DBEnv{Db: db}
	return env, nil
}