package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type AreaInteractor struct {
	AreaRepository AreaRepository
}

func (interactor *AreaInteractor) Add(u domain.Area) (area domain.Area, err error) {
	area, err = interactor.AreaRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *AreaInteractor) AreaById(id int) (area domain.Area, err error) {
	area, err = interactor.AreaRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *AreaInteractor) Areas() (area domain.Areas, err error) {
	area, err = interactor.AreaRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *AreaInteractor) Update(u domain.Area) (area domain.Area, err error) {
	area, err = interactor.AreaRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *AreaInteractor) DeleteByArea(area domain.Area) (err error) {
	err = interactor.AreaRepository.DeleteById(area)
	return
}
