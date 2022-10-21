package usecase

import(
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
)

type watchListUseCase struct{
	rep rep.WatchListRepository
}

type WatchListUseCase interface {
	GetWatchList(context.Context) ([]domain.WatchList, error)
	GetWatchListByID(context.Context, string) (domain.WatchList, error)
	CreateWatchList(context.Context, string, string, string, string, string) error
	UpdateWatchList(context.Context, string, string, string, string, string, string) error
	DestroyWatchList(context.Context, string) error
}

func NewWatchListUseCase(rep rep.WatchListRepository) WatchListUseCase {
	return &watchListUseCase{rep}
}

//watchListsの取得(Gets)
func (s *watchListUseCase) GetWatchList(c context.Context) ([]domain.WatchList, error) {
	watchList := domain.WatchList{}
	var watchLists []domain.WatchList
	rows , err := s.rep.All(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
      &watchList.ID,
      &watchList.Name,
      &watchList.Tel,
      &watchList.Email,
      &watchList.Address,
      &watchList.Representative,
      &watchList.CreatedAt,
      &watchList.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		watchLists = append(watchLists, watchList)
	}
	return watchLists , nil
}

//watchListの取得(Get)
func (s *watchListUseCase) GetWatchListByID(c context.Context, id string) (domain.WatchList, error) {
	watchList := domain.WatchList{}
	row , err := s.rep.Find(c, id)
	err = row.Scan(
		&watchList.ID,
		&watchList.Name,
		&watchList.Tel,
		&watchList.Email,
		&watchList.Address,
		&watchList.Representative,
		&watchList.CreatedAt,
		&watchList.UpdatedAt,
	)
	if err != nil {
		return watchList ,err
	}
	return watchList, nil
}

//watchListの作成(create)
func (s *watchListUseCase) CreateWatchList(
	c context.Context,
	Name string,
	Tel string,
	Email string,
	Address string,
	Representative string,
)error {
	err := s.rep.Create(c, Name, Tel, Email, Address, Representative)
	return err
}

//watchListの編集(Update)
func (s *watchListUseCase) UpdateWatchList(
	c context.Context,
	id string,
	Name string,
	Tel string,
	Email string,
	Address string,
	Representative string,
)error {
	err := s.rep.Update(c, id, Name, Tel, Email, Address, Representative)
	return err
}

// //watchListの削除(delete)
func (s *watchListUseCase) DestroyWatchList(
	c context.Context,
	id string,
)error {
	err := s.rep.Delete(c, id)
	return err
}
