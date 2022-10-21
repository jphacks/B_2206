package repository

import (
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
	"fmt"
)

type requestRepository struct {
	client db.Client
}

type RequestRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string, string, string) error
	Update(context.Context, string, string, string, string) error
	Delete(context.Context, string) error
	AllOrderWithUser(context.Context) (*sql.Rows,error)
	FindWithOrderItem(context.Context,string) (*sql.Row,error)
	GetPurchaseItemByOrderId(context.Context, string) (*sql.Rows,error)
	FindNewRecord(context.Context) (*sql.Row,error)
}

func NewRequestRepository(client db.Client) RequestRepository {
	return &requestRepository{client}
}

//全件取得
func (por *requestRepository) All(c context.Context) (*sql.Rows, error) {
	query := "select * from purchase_orders"
	rows, err := por.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//1件取得
func (por * requestRepository) Find(c context.Context, id string) (*sql.Row, error) {
	query := "select * from purchase_orders where id = "+ id
	row := por.client.DB().QueryRowContext(c ,query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
} 

//作成
func (por * requestRepository) Create(
	c context.Context,
	deadLine string,
	userId string,
	financeCheck string,
) error {
		var query = "insert into purchase_orders (deadline, user_id, finance_check) values ('" + deadLine + "'," + userId + "," + financeCheck +")"
		_, err := por.client.DB().ExecContext(c, query)
		fmt.Printf("\x1b[36m%s\n", query)
		return err
}

//編集
func (por * requestRepository) Update(
	c context.Context,
	id string,
	deadLine string,
	userId string,
	financeCheck string,
) error {
	var query = "update purchase_orders set deadline ='" + deadLine + "', user_id = " + userId + ",finance_check = " + financeCheck	+ " where id = " + id 
	_, err := por.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err 
}
//削除
func (por * requestRepository) Delete(
	c context.Context,
	id string,
)error {
	query := "Delete from purchase_orders where id =" + id
	_, err := por.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//orderに紐づくuserの取得(All)
func (p *requestRepository) AllOrderWithUser(c context.Context) (*sql.Rows, error){
	query := "select * from purchase_orders inner join users on purchase_orders.user_id = users.id;"
	rows , err := p.client.DB().QueryContext(c,query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//orderに紐づくuserの取得(byID)
func (p *requestRepository) FindWithOrderItem(c context.Context, id string) (*sql.Row, error) {
	query := " select * from purchase_orders inner join users on purchase_orders.user_id = users.id where purchase_orders.id =" +id
	row := p.client.DB().QueryRowContext(c,query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row,nil	
}

//指定したorder_idのitemを取得する
func (p *requestRepository) GetPurchaseItemByOrderId(c context.Context, requestID string) (*sql.Rows,error) {
	query := "select * from purchase_items where purchase_items.purchase_order_id ="+ requestID
	rows, err := p.client.DB().QueryContext(c, query)
	if err!= nil {
		return nil, errors.Wrapf(err, "cannot connect SQL") 
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//最新のレコードを取得
func (por *requestRepository) FindNewRecord(c context.Context) (*sql.Row,error) {
	query := "select * from purchase_orders order by id desc limit 1"
	row:= por.client.DB().QueryRowContext(c,query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row ,nil
}