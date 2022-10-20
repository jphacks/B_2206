package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type DetailRepository struct {
	SqlHandler
}

func (db *DetailRepository) Store(u domain.Detail) (detail domain.Detail, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	detail = u
	return
}

func (db *DetailRepository) FindById(id int) (detail domain.Detail, err error) {
	if err = db.FindBy(&detail,id).Error; err != nil {
		return
	}
	return
}

func (db *DetailRepository) FindAll() (details domain.Details, err error) {
	if err = db.Find(&details).Error; err != nil {
		return
	}
	return
}

func (db *DetailRepository) Update(u domain.Detail) (detail domain.Detail, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	detail = u
	return
}

func (db *DetailRepository) DeleteById(detail domain.Detail) (err error) {
	if err = db.Delete(&detail).Error; err != nil {
		return
	}
	return
}
