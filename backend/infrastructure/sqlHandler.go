package infrastructure

import (
	"os"
	"fmt"

	"github.com/jphacks/B_2206/backend/interfaces/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

// データベースへの値の挿入するメソッド
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.db.Create(value)
}

// 指定された条件に一致するレコードを探すメソッド
func (handler *SqlHandler) Find(value interface{}) *gorm.DB {
	fmt.Println("aasasasas")
	fmt.Println(handler.db.Find(value))
	return handler.db.Find(value)
}

func (hander *SqlHandler) FindBy(value interface{}, id int) *gorm.DB {
	return hander.db.First(value, id)
}

// データベースの値を更新し、値に主キーがない場合は挿入するメソッド
func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.db.Updates(value)
}

// 指定された条件に合う値を削除する。値に主キーがある場合は、主キーを条件に含む。
func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.db.Delete(value)
}
