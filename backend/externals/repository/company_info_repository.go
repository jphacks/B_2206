package repository

import(
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
	"fmt"
)

type companyInfoRepository struct {
	client db.Client
}

type CompanyInfoRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string, string, string) error
	Update(context.Context, string, string, string, string) error
	Delete(context.Context, string) error
}

func NewCompanyInfoRepository(client db.Client) CompanyInfoRepository{
	return &companyInfoRepository{client}
}

//全件取得
func (ssr *companyInfoRepository) All(c context.Context) (*sql.Rows, error){
	query :=  "select * from companyInfo_s"
	rows , err := ssr.client.DB().QueryContext(c,query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//１件取得
func (ssr *companyInfoRepository) Find(c context.Context, id string) (*sql.Row, error){
	query := "select *from companyInfo_s where id = " + id
	row := ssr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}

//作成
func (ssr *companyInfoRepository) Create(
	c context.Context,
	scale string,
	isColor string,
	price string,
)error {
	var query = "insert into companyInfo_s (scale, is_color, price) values ('" + scale + "'," + isColor + "," + price + ")"
	_, err := ssr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//編集
func (ssr *companyInfoRepository) Update(
	c context.Context,
	id string,
	scale string,
	isColor string,
	price string,
)error {
	var query = "update companyInfo_s set scale = '" + scale + "' , is_color = " + isColor + ", price = " + price + " where id = " + id
	_, err :=ssr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//削除
func (ssr *companyInfoRepository) Delete(
	c context.Context,
	id string,
)error {
	query := "Delete from companyInfo_s where id =" + id
	_, err := ssr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}