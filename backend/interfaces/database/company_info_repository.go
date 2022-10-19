package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type CompanyInfoRepository struct {
	SqlHandler
}

func (db *CompanyInfoRepository) Store(ci domain.CompanyInfo) (companyInfo domain.CompanyInfo, err error) {
	if err = db.Create(&ci).Error; err != nil {
		return
	}
	companyInfo = ci
	return
}

func (db *CompanyInfoRepository) FindById(id int) (companyInfo domain.CompanyInfo, err error) {
	if err = db.Find(&companyInfo).Error; err != nil {
		return
	}
	return
}

func (db *CompanyInfoRepository) FindAll() (companyInfos domain.CompanyInfos, err error) {
	if err = db.Find(&companyInfos).Error; err != nil {
		return
	}
	return
}

func (db *CompanyInfoRepository) Update(ci domain.CompanyInfo) (companyInfo domain.CompanyInfo, err error) {
	if err = db.Save(&ci).Error; err != nil {
		return
	}
	companyInfo = ci
	return
}

func (db *CompanyInfoRepository) DeleteById(companyInfo domain.CompanyInfo) (err error) {
	if err = db.Delete(&companyInfo).Error; err != nil {
		return
	}
	return
}
