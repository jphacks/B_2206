package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type CompanyInfoInteractor struct {
	CompanyInfoRepository CompanyInfoRepository
}

func (interactor *CompanyInfoInteractor) Add(u domain.CompanyInfo) (companyInfo domain.CompanyInfo, err error) {
	companyInfo, err = interactor.CompanyInfoRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *CompanyInfoInteractor) CompanyInfoById(id int) (companyInfo domain.CompanyInfo, err error) {
	companyInfo, err = interactor.CompanyInfoRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *CompanyInfoInteractor) CompanyInfos() (companyInfo domain.CompanyInfos, err error) {
	companyInfo, err = interactor.CompanyInfoRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *CompanyInfoInteractor) Update(u domain.CompanyInfo) (companyInfo domain.CompanyInfo, err error) {
	companyInfo, err = interactor.CompanyInfoRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *CompanyInfoInteractor) DeleteByCompanyInfo(companyInfo domain.CompanyInfo) (err error) {
	err = interactor.CompanyInfoRepository.DeleteById(companyInfo)
	return
}
