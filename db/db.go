package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sikozonpc/ecom/configs"
)

func NewStorage() (*sql.DB, error) {
	switch configs.Envs.DBDriver {
	case "mysql":
		return NewMySQLStorage()
	default:
		return NewSqliteStorage()
	}
}

func NewMySQLStorage() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func NewSqliteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
