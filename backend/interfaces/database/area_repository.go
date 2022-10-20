package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type AreaRepository struct {
	SqlHandler
}

func (db *AreaRepository) Store(ci domain.Area) (area domain.Area, err error) {
	if err = db.Create(&ci).Error; err != nil {
		return
	}
	area = ci
	return
}

func (db *AreaRepository) FindById(id int) (area domain.Area, err error) {
	if err = db.FindBy(&area,id).Error; err != nil {
		return
	}
	return
}

func (db *AreaRepository) FindAll() (areas domain.Areas, err error) {
	if err = db.Find(&areas).Error; err != nil {
		return
	}
	return
}

func (db *AreaRepository) Update(ci domain.Area) (area domain.Area, err error) {
	if err = db.Save(&ci).Error; err != nil {
		return
	}
	area = ci
	return
}

func (db *AreaRepository) DeleteById(area domain.Area) (err error) {
	if err = db.Delete(&area).Error; err != nil {
		return
	}
	return
}
