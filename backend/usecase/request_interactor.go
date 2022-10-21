package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type RequestInteractor struct {
	RequestRepository RequestRepository
}

// ユーザーを作成するメソッド
func (interactor *RequestInteractor) Add(u domain.Request) (request domain.Request, err error) {
	request, err = interactor.RequestRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *RequestInteractor) RequestById(id int) (request domain.Request, err error) {
	request, err = interactor.RequestRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *RequestInteractor) Requests() (request domain.Requests, err error) {
	request, err = interactor.RequestRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *RequestInteractor) Update(u domain.Request) (request domain.Request, err error) {
	request, err = interactor.RequestRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *RequestInteractor) DeleteByRequest(request domain.Request) (err error) {
	err = interactor.RequestRepository.DeleteById(request)
	return
}
