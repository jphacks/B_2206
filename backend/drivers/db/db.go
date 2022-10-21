package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type client struct {
	db *sql.DB
}

type Client interface {
	DB() *sql.DB
	CloseDB()
}

func ConnectMySQL() (client, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))

	if err != nil {
		return client{}, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("[Failed] Not Connect to MySQL") // 失敗
		return client{}, err
	} else {
		fmt.Println("[Success] Connect to MySQL") // 成功
		return client{db}, nil
	}
}

func (c client) CloseDB() {
	if c.db != nil {
		c.db.Close()
	}
}

func (c client) DB() *sql.DB {
	return c.db
}
