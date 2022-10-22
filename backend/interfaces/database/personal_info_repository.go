package database

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type PersonalInfoRepository struct {
	SqlHandler
}

func (db *PersonalInfoRepository) Store(p domain.PersonalInfo) (personalInfo domain.PersonalInfo, err error) {
	if err = db.Create(&p).Error; err != nil {
		return
	}
	personalInfo = p
	return
}

func (db *PersonalInfoRepository) FindById(id int) (personalInfo domain.PersonalInfo, err error) {
	if err = db.FindBy(&personalInfo,id).Error; err != nil {
		return
	}
	return
}

func (db *PersonalInfoRepository) FindAll() (personalInfos domain.PersonalInfos, err error) {
	if err = db.Find(&personalInfos).Error; err != nil {
		return
	}
	return
}

func (db *PersonalInfoRepository) Update(p domain.PersonalInfo) (personalInfo domain.PersonalInfo, err error) {
	if err = db.Save(&p).Error; err != nil {
		return
	}
	personalInfo = p
	return
}

func (db *PersonalInfoRepository) DeleteById(personalInfo domain.PersonalInfo) (err error) {
	if err = db.Delete(&personalInfo).Error; err != nil {
		return
	}
	return
}
