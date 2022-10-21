package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jphacks/B_2206/api/drivers/db"
	"github.com/pkg/errors"
)

type personalInfoReportRepository struct {
	client db.Client
}

type PersonalInfoReportRepository interface {
	All(context.Context) (*sql.Rows, error)
	Find(context.Context, string) (*sql.Row, error)
	Create(context.Context, string, string, string, string, string, string) error
	Update(context.Context, string, string, string, string, string, string, string) error
	Delete(context.Context, string) error
	AllWithOrderItem(context.Context) (*sql.Rows, error)
	FindWithOrderItem(context.Context, string) (*sql.Row, error)
	GetPersonalInfoItemByPersonalInfoOrderID(context.Context, string) (*sql.Rows, error)
	FindNewRecord(context.Context) (*sql.Row,error)
}


func NewPersonalInfoReportRepository(client db.Client) PersonalInfoReportRepository {
	return &personalInfoReportRepository{client}
}

//全件取得
func (prr *personalInfoReportRepository) All(c context.Context) (*sql.Rows, error) {
	query := "select * from personalInfo_reports"
	rows, err := prr.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connenct SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//1件取得
func (prr *personalInfoReportRepository) Find(c context.Context, id string) (*sql.Row, error) {
	query := "select * from personalInfo_reports where id =" + id
	row := prr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}

//作成
func (ppr *personalInfoReportRepository) Create(
	c context.Context,
	userId string,
	discount string,
	addition string,
	finance_check string,
	personalInfoOrderId string,
	remark string,
) error {
	var query = "insert into personalInfo_reports (user_id, discount, addition, finance_check, personalInfo_order_id, remark) values (" + userId + "," + discount + "," + addition + "," + finance_check + "," + personalInfoOrderId + ",'" + remark + "')"
	_, err := ppr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//編集
func (ppr *personalInfoReportRepository) Update(
	c context.Context,
	id string,
	userId string,
	discount string,
	addition string,
	finance_check string,
	personalInfoOrderId string,
	remark string,
) error {
	var query = "update personalInfo_reports set user_id =" + userId + ", discount =" + discount + ",addition =" + addition + ", finance_check =" + finance_check +", personalInfo_order_id =" + personalInfoOrderId + ", remark ='" + remark + "' where id = " + id 
	_, err := ppr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//削除
func (ppr *personalInfoReportRepository) Delete(
	c context.Context,
	id string,
) error {
	query := "Delete from personalInfo_reports where id =" + id
	_, err := ppr.client.DB().ExecContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return err
}

//PersonalInfo_reportに紐づく、PersonalInfo_orderからPersonalInfo_itemsの取得
func (ppr *personalInfoReportRepository) AllWithOrderItem(c context.Context) (*sql.Rows, error) {
	query := "select * from personalInfo_reports inner join users as report_user on personalInfo_reports.user_id = report_user.id inner join personalInfo_orders on personalInfo_reports.personalInfo_order_id = personalInfo_orders.id inner join users as order_user on personalInfo_orders.user_id = order_user.id"
	rows, err := ppr.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//idで選択しPersonalInfo_reportに紐づく、PersonalInfo_orderからPersonalInfo_itemsの取得
func (ppr *personalInfoReportRepository) FindWithOrderItem(c context.Context, id string) (*sql.Row, error) {
	query := "select * from personalInfo_reports inner join users as report_user on personalInfo_reports.user_id = report_user.id inner join personalInfo_orders on personalInfo_reports.personalInfo_order_id = personalInfo_orders.id inner join users as order_user on personalInfo_orders.user_id = order_user.id where personalInfo_reports.id = " + id
	row := ppr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}

//personalInfo_order_idに紐づいたpuchase_itemの取得
func (ppr *personalInfoReportRepository) GetPersonalInfoItemByPersonalInfoOrderID(c context.Context, personalInfoOrderID string) (*sql.Rows, error) {
	query := "select * from personalInfo_items where personalInfo_order_id = " + personalInfoOrderID
	rows, err := ppr.client.DB().QueryContext(c, query)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect SQL")
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return rows, nil
}

//personalInforeportの最新のレコード取得
func (ppr *personalInfoReportRepository) FindNewRecord(c context.Context) (*sql.Row, error) {
	query := "select * from personalInfo_reports order by id desc limit 1"
	row := ppr.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row, nil
}