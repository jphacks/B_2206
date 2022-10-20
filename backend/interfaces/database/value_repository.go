package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type ValueRepository struct {
	SqlHandler
}

func (db *ValueRepository) Store(u domain.Value) (value domain.Value, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	value = u
	return
}

func (db *ValueRepository) FindById(id int) (value domain.Value, err error) {
	if err = db.FindBy(&value, id).Error; err != nil {
		return
	}
	return
}

func (db *ValueRepository) FindAll() (values domain.Values, err error) {
	if err = db.Find(&values).Error; err != nil {
		return
	}
	return
}

func (db *ValueRepository) Update(u domain.Value) (value domain.Value, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	value = u
	return
}

func (db *ValueRepository) DeleteById(value domain.Value) (err error) {
	if err = db.Delete(&value).Error; err != nil {
		return
	}
	return
}
