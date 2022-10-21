package repository

import (
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
	"fmt"
)

type areaRepository struct {
	client db.Client
}

type AreaRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string) error
	Update(context.Context, string, string) error
	Destroy(context.Context, string) error
}

func NewAreaRepository(client db.Client) AreaRepository {
	return &areaRepository{client}
}

// 全件取得
func (yr *areaRepository) All(c context.Context) (*sql.Rows, error) {
	query := "select * from areas"
	rows, err := yr.client.DB().QueryContext(c,query )
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

// 1件取得
func (yr *areaRepository) Find(c context.Context, id string) (*sql.Row, error) {
	query := "select * from areas where id = "+id
	row := yr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}

// 作成
func (yr *areaRepository) Create(c context.Context, area string) error {
	query := "insert into areas (area) values ('"+area+"')"
	_, err := yr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 編集
func (yr *areaRepository) Update(c context.Context, id string, area string) error {
	query := "update areas set area = '"+area+"' where id = "+id
	_, err := yr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

// 削除
func (yr *areaRepository) Destroy(c context.Context, id string) error {
	query :="delete from areas where id = "+id
	_, err := yr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}
