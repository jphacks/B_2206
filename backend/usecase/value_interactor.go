package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type ValueInteractor struct {
	ValueRepository ValueRepository
}

// ユーザーを作成するメソッド
func (interactor *ValueInteractor) Add(u domain.Value) (value domain.Value, err error) {
	value, err = interactor.ValueRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *ValueInteractor) ValueById(id int) (value domain.Value, err error) {
	value, err = interactor.ValueRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *ValueInteractor) Values() (value domain.Values, err error) {
	value, err = interactor.ValueRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *ValueInteractor) Update(u domain.Value) (value domain.Value, err error) {
	value, err = interactor.ValueRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *ValueInteractor) DeleteByValue(value domain.Value) (err error) {
	err = interactor.ValueRepository.DeleteById(value)
	return
}
