package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type WatchListInteractor struct {
	WatchListRepository WatchListRepository
}

// ユーザーを作成するメソッド
func (interactor *WatchListInteractor) Add(u domain.WatchList) (watchList domain.WatchList, err error) {
	watchList, err = interactor.WatchListRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *WatchListInteractor) WatchListById(id int) (watchList domain.WatchList, err error) {
	watchList, err = interactor.WatchListRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *WatchListInteractor) WatchLists() (watchList domain.WatchLists, err error) {
	watchList, err = interactor.WatchListRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *WatchListInteractor) Update(u domain.WatchList) (watchList domain.WatchList, err error) {
	watchList, err = interactor.WatchListRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *WatchListInteractor) DeleteByWatchList(watchList domain.WatchList) (err error) {
	err = interactor.WatchListRepository.DeleteById(watchList)
	return
}
