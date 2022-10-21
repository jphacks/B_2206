package repository

import (
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
	"fmt"
)

type wishListRepository struct {
	client db.Client
}

type WishListRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string) error
	Update(context.Context, string, string) error
	Destroy(context.Context, string) error
}

func NewWishListRepository(client db.Client) WishListRepository {
	return &wishListRepository{client}
}

// 全件取得
func (dr *wishListRepository) All(c context.Context) (*sql.Rows, error) {
	query := "select * from wishLists"
	rows, err := dr.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

// 1件取得
func (dr *wishListRepository) Find(c context.Context, id string) (*sql.Row, error) {
	query := "select * from wishLists where id = "+id
	row := dr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}

// 作成
func (dr *wishListRepository) Create(c context.Context, name string) error {
	query := "insert into wishLists (name) values ('"+name+"')"
	_, err := dr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 編集
func (dr *wishListRepository) Update(c context.Context, id string, name string) error {
	query := "update wishLists set name = '"+name+"' where id = "+id
	_, err := dr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 削除
func (dr *wishListRepository) Destroy(c context.Context, id string) error {
	query := "delete from wishLists where id = "+id
	_, err := dr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}
