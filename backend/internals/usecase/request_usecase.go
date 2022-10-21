package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"strconv"
)

type requestUseCase struct {
	rep rep.RequestRepository
}

type RequestUseCase interface {
	GetRequests(context.Context) ([]domain.Request, error)
	GetRequestByID(context.Context, string) (domain.Request, error)
	CreateRequest(context.Context, string, string, string, string, string,string) error
	UpdateRequest(context.Context, string, string, string, string, string, string,string) error
	DestroyRequest(context.Context, string) error
	GetRequestsWithOrderItem(context.Context) ([]domain.RequestWithOrderItem, error)
	GetRequestWithOrderItemByID(context.Context, string) (domain.RequestWithOrderItem, error)
	GetRequestPostRecord(context.Context, string, string, string, string, string, string) (domain.Request, error)
	GetRequestPutRecord(context.Context, string, string, string, string, string, string, string) (domain.Request, error)
}

func NewRequestUseCase(rep rep.RequestRepository) RequestUseCase {
	return &requestUseCase{rep}
}

//Requestsの取得(Gets)
func (p *requestUseCase) GetRequests(c context.Context) ([]domain.Request, error) {
	request := domain.Request{}
	var requests []domain.Request
	rows, err := p.rep.All(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&request.ID,
			&request.UserID,
			&request.Discount,
			&request.Addition,
			&request.FinanceCheck,
			&request.PurchaseOrderID,
			&request.Remark,
			&request.CreatedAt,
			&request.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

//requestの取得(Get)
func (p *requestUseCase) GetRequestByID(c context.Context, id string) (domain.Request, error) {
	request := domain.Request{}
	row, err := p.rep.Find(c, id)
	err = row.Scan(
		&request.ID,
		&request.UserID,
		&request.Discount,
		&request.Addition,
		&request.FinanceCheck,
		&request.PurchaseOrderID,
		&request.Remark,
		&request.CreatedAt,
		&request.UpdatedAt,
	)
	if err != nil {
		return request, err
	}
	return request, nil
}

//Requestの作成(create)
func (p *requestUseCase) CreateRequest(
	c context.Context,
	UserID string,
	Discount string,
	Addition string,
	FinanceCheck string,
	PurchaseOrderID string,
	Remark string,
) error {
	err := p.rep.Create(c, UserID ,Discount, Addition, FinanceCheck, PurchaseOrderID,Remark)
	return err
}

//Requestの修正(update)
func (p *requestUseCase) UpdateRequest(
	c context.Context,
	id string,
	UserID string,
	Discount string,
	Addition string,
	FinanceCheck string,
	PurchaseOrderID string,
	Remark string,
) error {
	err := p.rep.Update(c, id, UserID ,Discount, Addition, FinanceCheck, PurchaseOrderID,Remark)
	return err
}

//Requestの削除(delate)
func (p *requestUseCase) DestroyRequest(
	c context.Context,
	id string,
) error {
	err := p.rep.Delete(c, id)
	return err
}

//Purchase_reportに紐づく、Purchase_orderからPurchase_itemsの取得(GETS)
func (p *requestUseCase) GetRequestsWithOrderItem(c context.Context) ([]domain.RequestWithOrderItem, error) {
	requestwithorderitem := domain.RequestWithOrderItem{}
	var requestwithorderitems []domain.RequestWithOrderItem
	purchaseItem := domain.PurchaseItem{}
	var purchaseItems []domain.PurchaseItem
	rows, err := p.rep.AllWithOrderItem(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&requestwithorderitem.Request.ID,
			&requestwithorderitem.Request.UserID,
			&requestwithorderitem.Request.Discount,
			&requestwithorderitem.Request.Addition,
			&requestwithorderitem.Request.FinanceCheck,
			&requestwithorderitem.Request.PurchaseOrderID,
			&requestwithorderitem.Request.Remark,
			&requestwithorderitem.Request.CreatedAt,
			&requestwithorderitem.Request.UpdatedAt,
			&requestwithorderitem.ReportUser.ID,
			&requestwithorderitem.ReportUser.Name,
			&requestwithorderitem.ReportUser.BureauID,
			&requestwithorderitem.ReportUser.RoleID,
			&requestwithorderitem.ReportUser.CreatedAt,
			&requestwithorderitem.ReportUser.UpdatedAt,
			&requestwithorderitem.PurchaseOrder.ID,
			&requestwithorderitem.PurchaseOrder.DeadLine,
			&requestwithorderitem.PurchaseOrder.UserID,
			&requestwithorderitem.PurchaseOrder.FinanceCheck,
			&requestwithorderitem.PurchaseOrder.CreatedAt,
			&requestwithorderitem.PurchaseOrder.UpdatedAt,
			&requestwithorderitem.OrderUser.ID,
			&requestwithorderitem.OrderUser.Name,
			&requestwithorderitem.OrderUser.BureauID,
			&requestwithorderitem.OrderUser.RoleID,
			&requestwithorderitem.OrderUser.CreatedAt,
			&requestwithorderitem.OrderUser.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rows, err := p.rep.GetPurchaseItemByPurchaseOrderID(c, strconv.Itoa(int(requestwithorderitem.Request.PurchaseOrderID)))
		for rows.Next() {
			err := rows.Scan(
				&purchaseItem.ID,
				&purchaseItem.Item,
				&purchaseItem.Price,
				&purchaseItem.Quantity,
				&purchaseItem.Detail,
				&purchaseItem.Url,
				&purchaseItem.PurchaseOrderID,
				&purchaseItem.FinanceCheck,
				&purchaseItem.CreatedAt,
				&purchaseItem.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			purchaseItems = append(purchaseItems, purchaseItem)
		}
		requestwithorderitem.PurchaseItems = purchaseItems
		requestwithorderitems = append(requestwithorderitems, requestwithorderitem)
		purchaseItems = nil
	}
	return requestwithorderitems, nil
}

//idで選択しPurchase_reportに紐づく、Purchase_orderからPurchase_itemsの取得(GETS)
func (p *requestUseCase) GetRequestWithOrderItemByID(c context.Context, id string) (domain.RequestWithOrderItem, error) {
	requestwithorderitem := domain.RequestWithOrderItem{}
	purchaseItem := domain.PurchaseItem{}
	var purchaseItems []domain.PurchaseItem
	row, err := p.rep.FindWithOrderItem(c, id)
	err = row.Scan(
		&requestwithorderitem.Request.ID,
		&requestwithorderitem.Request.UserID,
		&requestwithorderitem.Request.Discount,
		&requestwithorderitem.Request.Addition,
		&requestwithorderitem.Request.FinanceCheck,
		&requestwithorderitem.Request.PurchaseOrderID,
		&requestwithorderitem.Request.Remark,
		&requestwithorderitem.Request.CreatedAt,
		&requestwithorderitem.Request.UpdatedAt,
		&requestwithorderitem.ReportUser.ID,
		&requestwithorderitem.ReportUser.Name,
		&requestwithorderitem.ReportUser.BureauID,
		&requestwithorderitem.ReportUser.RoleID,
		&requestwithorderitem.ReportUser.CreatedAt,
		&requestwithorderitem.ReportUser.UpdatedAt,
		&requestwithorderitem.PurchaseOrder.ID,
		&requestwithorderitem.PurchaseOrder.DeadLine,
		&requestwithorderitem.PurchaseOrder.UserID,
		&requestwithorderitem.PurchaseOrder.FinanceCheck,
		&requestwithorderitem.PurchaseOrder.CreatedAt,
		&requestwithorderitem.PurchaseOrder.UpdatedAt,
		&requestwithorderitem.OrderUser.ID,
		&requestwithorderitem.OrderUser.Name,
		&requestwithorderitem.OrderUser.BureauID,
		&requestwithorderitem.OrderUser.RoleID,
		&requestwithorderitem.OrderUser.CreatedAt,
		&requestwithorderitem.OrderUser.UpdatedAt,
	)
	if err != nil {
		return requestwithorderitem, err
	}
	rows, err := p.rep.GetPurchaseItemByPurchaseOrderID(c, strconv.Itoa(int(requestwithorderitem.Request.PurchaseOrderID)))
	for rows.Next() {
		err := rows.Scan(
			&purchaseItem.ID,
			&purchaseItem.Item,
			&purchaseItem.Price,
			&purchaseItem.Quantity,
			&purchaseItem.Detail,
			&purchaseItem.Url,
			&purchaseItem.PurchaseOrderID,
			&purchaseItem.FinanceCheck,
			&purchaseItem.CreatedAt,
			&purchaseItem.UpdatedAt,
		)
		if err != nil {
			return requestwithorderitem, err
		}
		purchaseItems = append(purchaseItems, purchaseItem)
	}
	requestwithorderitem.PurchaseItems = purchaseItems
	return requestwithorderitem, nil
}

//postした際に作成されるレコードの取得
func (p *requestUseCase) GetRequestPostRecord(
	c context.Context,
	UserID string,
	Discount string,
	Addition string,
	FinanceCheck string,
	PurchaseOrderID string,
	Remark string,
)(domain.Request,error){
	p.rep.Create(c, UserID ,Discount, Addition, FinanceCheck, PurchaseOrderID,Remark)
	request := domain.Request{}
	row, err := p.rep.FindNewRecord(c)
	err = row.Scan(
		&request.ID,
		&request.UserID,
		&request.Discount,
		&request.Addition,
		&request.FinanceCheck,
		&request.PurchaseOrderID,
		&request.Remark,
		&request.CreatedAt,
		&request.UpdatedAt,
	)
	if err != nil {
		return request, err
	}
	return request, nil
}

//putした際に更新されるレコードの取得
func (p *requestUseCase) GetRequestPutRecord(
	c context.Context,
	id string,
	UserID string,
	Discount string,
	Addition string,
	FinanceCheck string,
	PurchaseOrderID string,
	Remark string,
)(domain.Request, error) {
	p.rep.Update(c,id, UserID, Discount, Addition, FinanceCheck, PurchaseOrderID, Remark)
	request := domain.Request{}
	row, err := p.rep.Find(c, id)
	err = row.Scan(
		&request.ID,
		&request.UserID,
		&request.Discount,
		&request.Addition,
		&request.FinanceCheck,
		&request.PurchaseOrderID,
		&request.Remark,
		&request.CreatedAt,
		&request.UpdatedAt,
	)
	if err != nil {
		return request, err
	}
	return request, nil
}