package main

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/infrastructure"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var (
	db  *gorm.DB
	err error
	dsn = os.Getenv("DSN")
)

func main() {
	dbinit()
	e := echo.New()
	infrastructure.Init()
	e.GET("/", healthCheck)
	e.Logger.Fatal(e.Start(":1323"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check")
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.Migrator().CreateTable(domain.Area{})
	db.Migrator().CreateTable(domain.Classification{})
	db.Migrator().CreateTable(domain.CompanyInfo{})
	db.Migrator().CreateTable(domain.Detail{})
	db.Migrator().CreateTable(domain.DetailsRange{})
	db.Migrator().CreateTable(domain.DetailsValue{})
	db.Migrator().CreateTable(domain.Matching{})
	db.Migrator().CreateTable(domain.PersonalInfo{})
	db.Migrator().CreateTable(domain.Range{})
	db.Migrator().CreateTable(domain.Request{})
	db.Migrator().CreateTable(domain.Tag{})
	db.Migrator().CreateTable(domain.User{})
	db.Migrator().CreateTable(domain.Value{})
	db.Migrator().CreateTable(domain.WatchList{})
}
