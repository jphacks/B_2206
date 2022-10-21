package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type TagInteractor struct {
	TagRepository TagRepository
}

// ユーザーを作成するメソッド
func (interactor *TagInteractor) Add(u domain.Tag) (tag domain.Tag, err error) {
	tag, err = interactor.TagRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *TagInteractor) TagById(id int) (tag domain.Tag, err error) {
	tag, err = interactor.TagRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *TagInteractor) Tags() (tag domain.Tags, err error) {
	tag, err = interactor.TagRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *TagInteractor) Update(u domain.Tag) (tag domain.Tag, err error) {
	tag, err = interactor.TagRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *TagInteractor) DeleteByTag(tag domain.Tag) (err error) {
	err = interactor.TagRepository.DeleteById(tag)
	return
}
