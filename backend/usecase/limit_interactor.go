package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type LimitInteractor struct {
	LimitRepository LimitRepository
}

// ユーザーを作成するメソッド
func (interactor *LimitInteractor) Add(u domain.Limit) (limit domain.Limit, err error) {
	limit, err = interactor.LimitRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *LimitInteractor) LimitById(id int) (limit domain.Limit, err error) {
	limit, err = interactor.LimitRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *LimitInteractor) Limits() (limit domain.Limits, err error) {
	limit, err = interactor.LimitRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *LimitInteractor) Update(u domain.Limit) (limit domain.Limit, err error) {
	limit, err = interactor.LimitRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *LimitInteractor) DeleteByLimit(limit domain.Limit) (err error) {
	err = interactor.LimitRepository.DeleteById(limit)
	return
}
