package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (db *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = db.FindBy(&user,id).Error; err != nil {
		return
	}
	return
}

func (db *UserRepository) FindAll() (users domain.Users, err error) {
	if err = db.Find(&users).Error; err != nil {
		return
	}
	return
}

func (db *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (db *UserRepository) DeleteById(user domain.User) (err error) {
	if err = db.Delete(&user).Error; err != nil {
		return
	}
	return
}
