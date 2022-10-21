package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type MatchingRepository struct {
	SqlHandler
}

func (db *MatchingRepository) Store(u domain.Matching) (matching domain.Matching, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	matching = u
	return
}

func (db *MatchingRepository) FindById(id int) (matching domain.Matching, err error) {
	if err = db.FindBy(&matching,id).Error; err != nil {
		return
	}
	return
}

func (db *MatchingRepository) FindAll() (matchings domain.Matchings, err error) {
	if err = db.Find(&matchings).Error; err != nil {
		return
	}
	return
}

func (db *MatchingRepository) Update(u domain.Matching) (matching domain.Matching, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	matching = u
	return
}

func (db *MatchingRepository) DeleteById(matching domain.Matching) (err error) {
	if err = db.Delete(&matching).Error; err != nil {
		return
	}
	return
}
