package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type MatchingInteractor struct {
	MatchingRepository MatchingRepository
}

// ユーザーを作成するメソッド
func (interactor *MatchingInteractor) Add(u domain.Matching) (matching domain.Matching, err error) {
	matching, err = interactor.MatchingRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *MatchingInteractor) MatchingById(id int) (matching domain.Matching, err error) {
	matching, err = interactor.MatchingRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *MatchingInteractor) Matchings() (matching domain.Matchings, err error) {
	matching, err = interactor.MatchingRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *MatchingInteractor) Update(u domain.Matching) (matching domain.Matching, err error) {
	matching, err = interactor.MatchingRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *MatchingInteractor) DeleteByMatching(matching domain.Matching) (err error) {
	err = interactor.MatchingRepository.DeleteById(matching)
	return
}
