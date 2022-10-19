package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}

type PersonalInfoRepository interface {
	Store(domain.PersonalInfo) (domain.PersonalInfo, error)
	FindById(int) (domain.PersonalInfo, error)
	FindAll() (domain.PersonalInfos, error)
	Update(domain.PersonalInfo) (domain.PersonalInfo, error)
	DeleteById(domain.PersonalInfo) error
}
