package main

import (
	"github.com/jphacks/B_2206/api/internals/di"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	client := di.InitializeServer()
	defer client.CloseDB()
}
