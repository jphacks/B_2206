package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type ClassificationRepository struct {
	SqlHandler
}

func (db *ClassificationRepository) Store(u domain.Classification) (classification domain.Classification, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	classification = u
	return
}

func (db *ClassificationRepository) FindById(id int) (classification domain.Classification, err error) {
	if err = db.FindBy(&classification, id).Error; err != nil {
		return
	}
	return
}

func (db *ClassificationRepository) FindAll() (classifications domain.Classifications, err error) {
	if err = db.Find(&classifications).Error; err != nil {
		return
	}
	return
}

func (db *ClassificationRepository) Update(u domain.Classification) (classification domain.Classification, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	classification = u
	return
}

func (db *ClassificationRepository) DeleteById(classification domain.Classification) (err error) {
	if err = db.Delete(&classification).Error; err != nil {
		return
	}
	return
}
