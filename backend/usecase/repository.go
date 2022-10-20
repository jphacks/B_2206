package usecase

import (
	"github.com/jphacks/B_2206/backend/domain"
)

type PersonalInfoRepository interface {
	Store(domain.PersonalInfo) (domain.PersonalInfo, error)
	FindById(int) (domain.PersonalInfo, error)
	FindAll() (domain.PersonalInfos, error)
	Update(domain.PersonalInfo) (domain.PersonalInfo, error)
	DeleteById(domain.PersonalInfo) error
}

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}

type WatchListRepository interface {
	Store(domain.WatchList) (domain.WatchList, error)
	FindById(int) (domain.WatchList, error)
	FindAll() (domain.WatchLists, error)
	Update(domain.WatchList) (domain.WatchList, error)
	DeleteById(domain.WatchList) error
}

