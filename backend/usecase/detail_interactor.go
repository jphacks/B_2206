package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type DetailInteractor struct {
	DetailRepository DetailRepository
}

// ユーザーを作成するメソッド
func (interactor *DetailInteractor) Add(u domain.Detail) (detail domain.Detail, err error) {
	detail, err = interactor.DetailRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *DetailInteractor) DetailById(id int) (detail domain.Detail, err error) {
	detail, err = interactor.DetailRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *DetailInteractor) Details() (detail domain.Details, err error) {
	detail, err = interactor.DetailRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *DetailInteractor) Update(u domain.Detail) (detail domain.Detail, err error) {
	detail, err = interactor.DetailRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *DetailInteractor) DeleteByDetail(detail domain.Detail) (err error) {
	err = interactor.DetailRepository.DeleteById(detail)
	return
}
