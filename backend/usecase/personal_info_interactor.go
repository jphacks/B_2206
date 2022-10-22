package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type PersonalInfoInteractor struct {
	PersonalInfoRepository PersonalInfoRepository
}

func (interactor *PersonalInfoInteractor) Add(u domain.PersonalInfo) (personalInfo domain.PersonalInfo, err error) {
	personalInfo, err = interactor.PersonalInfoRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *PersonalInfoInteractor) PersonalInfoById(id int) (personalInfo domain.PersonalInfo, err error) {
	personalInfo, err = interactor.PersonalInfoRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *PersonalInfoInteractor) PersonalInfos() (personalInfo domain.PersonalInfos, err error) {
	personalInfo, err = interactor.PersonalInfoRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *PersonalInfoInteractor) Update(u domain.PersonalInfo) (personalInfo domain.PersonalInfo, err error) {
	personalInfo, err = interactor.PersonalInfoRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *PersonalInfoInteractor) DeleteByPersonalInfo(personalInfo domain.PersonalInfo) (err error) {
	err = interactor.PersonalInfoRepository.DeleteById(personalInfo)
	return
}
