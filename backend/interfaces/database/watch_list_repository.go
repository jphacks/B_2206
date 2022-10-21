package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type WatchListRepository struct {
	SqlHandler
}

func (db *WatchListRepository) Store(u domain.WatchList) (watchList domain.WatchList, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	watchList = u
	return
}

func (db *WatchListRepository) FindById(id int) (watchList domain.WatchList, err error) {
	if err = db.FindBy(&watchList,id).Error; err != nil {
		return
	}
	return
}

func (db *WatchListRepository) FindAll() (watchLists domain.WatchLists, err error) {
	if err = db.Find(&watchLists).Error; err != nil {
		return
	}
	return
}

func (db *WatchListRepository) Update(u domain.WatchList) (watchList domain.WatchList, err error) {
	if err = db.Save(&u).Error; err != nil {
		return
	}
	watchList = u
	return
}

func (db *WatchListRepository) DeleteById(watchList domain.WatchList) (err error) {
	if err = db.Delete(&watchList).Error; err != nil {
		return
	}
	return
}
