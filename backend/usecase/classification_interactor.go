package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type ClassificationInteractor struct {
	ClassificationRepository ClassificationRepository
}

// ユーザーを作成するメソッド
func (interactor *ClassificationInteractor) Add(u domain.Classification) (classification domain.Classification, err error) {
	classification, err = interactor.ClassificationRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *ClassificationInteractor) ClassificationById(id int) (classification domain.Classification, err error) {
	classification, err = interactor.ClassificationRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *ClassificationInteractor) Classifications() (classification domain.Classifications, err error) {
	classification, err = interactor.ClassificationRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *ClassificationInteractor) Update(u domain.Classification) (classification domain.Classification, err error) {
	classification, err = interactor.ClassificationRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *ClassificationInteractor) DeleteByClassification(classification domain.Classification) (err error) {
	err = interactor.ClassificationRepository.DeleteById(classification)
	return
}
