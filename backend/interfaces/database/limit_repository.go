package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type LimitRepository struct {
	SqlHandler
}

func (db *LimitRepository) Store(u domain.Limit) (limit domain.Limit, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	limit = u
	return
}

func (db *LimitRepository) FindById(id int) (limit domain.Limit, err error) {
	if err = db.FindBy(&limit,id).Error; err != nil {
		return
	}
	return
}

func (db *LimitRepository) FindAll() (limits domain.Limits, err error) {
	if err = db.Find(&limits).Error; err != nil {
		return
	}
	return
}

func (db *LimitRepository) Update(u domain.Limit) (limit domain.Limit, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	limit = u
	return
}

func (db *LimitRepository) DeleteById(limit domain.Limit) (err error) {
	if err = db.Delete(&limit).Error; err != nil {
		return
	}
	return
}
