package repository

import (
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
	"fmt"
)

type limitRepository struct {
	client db.Client
}

type LimitRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string) error
	Update(context.Context, string, string) error
	Destroy(context.Context, string) error
}

func NewLimitRepository(client db.Client) LimitRepository {
	return &limitRepository{client}
}

// 全件取得
func (sr *limitRepository) All(c context.Context) (*sql.Rows, error) {
	query := "select * from limits"
	rows, err := sr.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

// 1件取得
func (sr *limitRepository) Find(c context.Context, id string) (*sql.Row, error) {
	query := "select * from limits where id = "+id
	row := sr.client.DB().QueryRowContext(c, query)
	return row, nil
}

// 作成
func (sr *limitRepository) Create(c context.Context, name string) error {
	query := "insert into limits (name) values ('"+name+"')"
	_, err := sr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 編集
func (sr *limitRepository) Update(c context.Context, id string, name string) error {
	query := "update limits set name = '"+name+"' where id = "+id
	_, err := sr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 削除
func (sr *limitRepository) Destroy(c context.Context, id string) error {
	query := "delete from limits where id = "+id
	_, err := sr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}
