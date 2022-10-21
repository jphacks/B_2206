package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormClient struct {
	db *gorm.DB
}

type GormClient interface {
	DB() *gorm.DB
}

func (c gormClient) DB() *gorm.DB {
	return c.db
}

func ConnectMySQLFromGorm() (GormClient, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		fmt.Println("[Failed] Not Connect to MySQL")
		return nil, err
	} else {
		fmt.Println("[Success] Connect to MySQL")
		return gormClient{db}, nil
	}
}
