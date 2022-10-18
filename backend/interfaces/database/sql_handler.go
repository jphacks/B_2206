package database

import "gorm.io/gorm"

// database/user_repository.goでこれを呼び出す。(DIP: 依存関係逆転の原則)
type SqlHandler interface {
	Create(interface{}) *gorm.DB
	Find(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}
