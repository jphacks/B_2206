package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type TagRepository struct {
	SqlHandler
}

func (db *TagRepository) Store(u domain.Tag) (tag domain.Tag, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	tag = u
	return
}

func (db *TagRepository) FindById(id int) (tag domain.Tag, err error) {
	if err = db.FindBy(&tag, id).Error; err != nil {
		return
	}
	return
}

func (db *TagRepository) FindAll() (tags domain.Tags, err error) {
	if err = db.Find(&tags).Error; err != nil {
		return
	}
	return
}

func (db *TagRepository) Update(u domain.Tag) (tag domain.Tag, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	tag = u
	return
}

func (db *TagRepository) DeleteById(tag domain.Tag) (err error) {
	if err = db.Delete(&tag).Error; err != nil {
		return
	}
	return
}
