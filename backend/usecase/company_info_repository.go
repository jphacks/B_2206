package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type CompanyInfoRepository interface {
	Store(domain.CompanyInfo) (domain.CompanyInfo, error)
	FindById(int) (domain.CompanyInfo, error)
	FindAll() (domain.CompanyInfos, error)
	Update(domain.CompanyInfo) (domain.CompanyInfo, error)
	DeleteById(domain.CompanyInfo) error
}
