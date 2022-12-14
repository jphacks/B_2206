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

type CompanyInfoRepository interface {
	Store(domain.CompanyInfo) (domain.CompanyInfo, error)
	FindById(int) (domain.CompanyInfo, error)
	FindAll() (domain.CompanyInfos, error)
	Update(domain.CompanyInfo) (domain.CompanyInfo, error)
	DeleteById(domain.CompanyInfo) error
}

type RequestRepository interface {
	Store(domain.Request) (domain.Request, error)
	FindById(int) (domain.Request, error)
	FindAll() (domain.Requests, error)
	Update(domain.Request) (domain.Request, error)
	DeleteById(domain.Request) error
}

type AreaRepository interface {
	Store(domain.Area) (domain.Area, error)
	FindById(int) (domain.Area, error)
	FindAll() (domain.Areas, error)
	Update(domain.Area) (domain.Area, error)
	DeleteById(domain.Area) error
}

type ClassificationRepository interface {
	Store(domain.Classification) (domain.Classification, error)
	FindById(int) (domain.Classification, error)
	FindAll() (domain.Classifications, error)
	Update(domain.Classification) (domain.Classification, error)
	DeleteById(domain.Classification) error
}

type TagRepository interface {
	Store(domain.Tag) (domain.Tag, error)
	FindById(int) (domain.Tag, error)
	FindAll() (domain.Tags, error)
	Update(domain.Tag) (domain.Tag, error)
	DeleteById(domain.Tag) error
}

type MatchingRepository interface {
	Store(domain.Matching) (domain.Matching, error)
	FindById(int) (domain.Matching, error)
	FindAll() (domain.Matchings, error)
	Update(domain.Matching) (domain.Matching, error)
	DeleteById(domain.Matching) error
}

type LimitRepository interface {
	Store(domain.Limit) (domain.Limit, error)
	FindById(int) (domain.Limit, error)
	FindAll() (domain.Limits, error)
	Update(domain.Limit) (domain.Limit, error)
	DeleteById(domain.Limit) error
}

type ValueRepository interface {
	Store(domain.Value) (domain.Value, error)
	FindById(int) (domain.Value, error)
	FindAll() (domain.Values, error)
	Update(domain.Value) (domain.Value, error)
	DeleteById(domain.Value) error
}
