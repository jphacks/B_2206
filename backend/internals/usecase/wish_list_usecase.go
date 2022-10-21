package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type wishListUseCase struct {
	rep rep.WishListRepository
}

type WishListUseCase interface {
	GetWishLists(context.Context) ([]domain.WishList, error)
	GetWishListByID(context.Context, string) (domain.WishList, error)
	CreateWishList(context.Context, string) error
	UpdateWishList(context.Context, string, string) error
	DestroyWishList(context.Context, string) error
}

func NewWishListUseCase(rep rep.WishListRepository) WishListUseCase {
	return &wishListUseCase{rep}
}

func (y *wishListUseCase) GetWishLists(c context.Context) ([]domain.WishList, error) {

	wishList := domain.WishList{}
	var wishLists []domain.WishList

	// クエリー実行
	rows, err := y.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&wishList.ID,
			&wishList.WishList,
			&wishList.CreatedAt,
			&wishList.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}

		wishLists = append(wishLists, wishList)
	}
	return wishLists, nil
}

func (y *wishListUseCase) GetWishListByID(c context.Context, id string) (domain.WishList, error) {
	var wishList domain.WishList

	row, err := y.rep.Find(c, id)
	err = row.Scan(
		&wishList.ID,
		&wishList.WishList,
		&wishList.CreatedAt,
		&wishList.UpdatedAt,
	)

	if err != nil {
		return wishList, err
	}

	return wishList, nil
}

func (y *wishListUseCase) CreateWishList(c context.Context, wishList string) error {
	err := y.rep.Create(c, wishList)
	return err
}

func (y *wishListUseCase) UpdateWishList(c context.Context, id string, wishList string) error {
	err := y.rep.Update(c, id, wishList)
	return err
}

func (y *wishListUseCase) DestroyWishList(c context.Context, id string) error {
	err := y.rep.Destroy(c, id)
	return err
}
