package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type RequestRepository struct {
	SqlHandler
}

func (db *RequestRepository) Store(ci domain.Request) (request domain.Request, err error) {
	if err = db.Create(&ci).Error; err != nil {
		return
	}
	request = ci
	return
}

func (db *RequestRepository) FindById(id int) (request domain.Request, err error) {
	if err = db.Find(&request).Error; err != nil {
		return
	}
	return
}

func (db *RequestRepository) FindAll() (requests domain.Requests, err error) {
	if err = db.Find(&requests).Error; err != nil {
		return
	}
	return
}

func (db *RequestRepository) Update(ci domain.Request) (request domain.Request, err error) {
	if err = db.Save(&ci).Error; err != nil {
		return
	}
	request = ci
	return
}

func (db *RequestRepository) DeleteById(request domain.Request) (err error) {
	if err = db.Delete(&request).Error; err != nil {
		return
	}
	return
}
