package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"strconv"
)

type personalInfoUseCase struct {
	rep rep.PersonalInfoRepository
}

type PersonalInfoUseCase interface {
	GetPersonalInfos(context.Context) ([]domain.PersonalInfo, error)
	GetPersonalInfoByID(context.Context, string) (domain.PersonalInfo, error)
	CreatePersonalInfo(context.Context, string, string, string) error
	UpdatePersonalInfo(context.Context, string, string, string, string) error
	DestroyPersonalInfo(context.Context, string) error
	GetOrderWithUserItem(context.Context) ([]domain.OrderWithItemAndUser,error)
	GetOrderWithUserItemByID(context.Context, string) (domain.OrderWithItemAndUser,error)
	GetPersonalInfoNewRecord(context.Context, string, string,string) (domain.PersonalInfo,error)
	GetPersonalInfoEdit(context.Context,string,string,string,string) (domain.PersonalInfo,error)
}


func NewPersonalInfoUseCase(rep rep.PersonalInfoRepository) PersonalInfoUseCase {
	return &personalInfoUseCase{rep}
}

//PersonalInfosの取得(Gets)
func (p *personalInfoUseCase) GetPersonalInfos(c context.Context) ([]domain.PersonalInfo, error)	{
	personalInfo := domain.PersonalInfo{}
		var personalInfos []domain.PersonalInfo
		rows , err := p.rep.All(c)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(
				&personalInfo.ID,
				&personalInfo.DeadLine,
				&personalInfo.UserID,
				&personalInfo.FinanceCheck,
				&personalInfo.CreatedAt,
				&personalInfo.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			personalInfos = append(personalInfos, personalInfo)
		}
		return personalInfos, nil
}

//PersonalInfoの取得(Get)
func (p *personalInfoUseCase) GetPersonalInfoByID(c context.Context, id string) (domain.PersonalInfo, error){
	personalInfo := domain.PersonalInfo{}
	row, err := p.rep.Find(c, id)
	err = row.Scan(
		&personalInfo.ID,
		&personalInfo.DeadLine,
		&personalInfo.UserID,
		&personalInfo.FinanceCheck,
		&personalInfo.CreatedAt,
		&personalInfo.UpdatedAt,
	)
	if err != nil {
		return personalInfo, err
	}
	return personalInfo, nil
}

//PurcahseOrderの作成(create)
func (p *personalInfoUseCase) CreatePersonalInfo(
	c context.Context, 
	DeadLine string,
	userID string,
	FinanceCheck string,
) error {
	err := p.rep.Create(c, DeadLine, userID, FinanceCheck)
	return err
}

//PersonalInfoの修正(Update)
func (p *personalInfoUseCase) UpdatePersonalInfo(
	c context.Context,
	id string,
	DeadLine string,
	userID string,
	FinanceCheck string,
) error {
	err := p.rep.Update(c, id, DeadLine, userID, FinanceCheck)
	return err
}

//PersonalInfoの削除(delete)
func (p *personalInfoUseCase) DestroyPersonalInfo(
	c context.Context,
	id string,
)error {
	err := p.rep.Delete(c, id)
	return err
}

//Purchase_orderに紐づくUserとItemの取得(All)
func (p *personalInfoUseCase) GetOrderWithUserItem(c context.Context) ([]domain.OrderWithItemAndUser,error){
	orderWithUserAndItem := domain.OrderWithItemAndUser{}
	var orderWithUserAndItems []domain.OrderWithItemAndUser
	purchaseItem := domain.PurchaseItem{}
	var purchaseItems []domain.PurchaseItem
	rows, err := p.rep.AllOrderWithUser(c)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		err := rows.Scan(
			&orderWithUserAndItem.PersonalInfo.ID,
			&orderWithUserAndItem.PersonalInfo.DeadLine,
			&orderWithUserAndItem.PersonalInfo.UserID,
			&orderWithUserAndItem.PersonalInfo.FinanceCheck,
			&orderWithUserAndItem.PersonalInfo.CreatedAt,
			&orderWithUserAndItem.PersonalInfo.UpdatedAt,
			&orderWithUserAndItem.User.ID,
			&orderWithUserAndItem.User.Name,
			&orderWithUserAndItem.User.BureauID,
			&orderWithUserAndItem.User.RoleID,
			&orderWithUserAndItem.User.CreatedAt,
			&orderWithUserAndItem.User.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rows, err := p.rep.GetPurchaseItemByOrderId(c, strconv.Itoa(int(orderWithUserAndItem.PersonalInfo.ID)))
		for rows.Next(){
			err := rows.Scan(
				&purchaseItem.ID,
				&purchaseItem.Item,
				&purchaseItem.Price,
				&purchaseItem.Quantity,
				&purchaseItem.Detail,
				&purchaseItem.Url,
				&purchaseItem.PersonalInfoID,
				&purchaseItem.FinanceCheck,
				&purchaseItem.CreatedAt,
				&purchaseItem.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			purchaseItems = append(purchaseItems,purchaseItem)
		}
		orderWithUserAndItem.PurchaseItem = purchaseItems
		orderWithUserAndItems = append(orderWithUserAndItems,orderWithUserAndItem)
		purchaseItems = nil
	}
	return orderWithUserAndItems,nil
}

//Purchase_orderに紐づくUserとItemの取得(ByID)
func (p *personalInfoUseCase) GetOrderWithUserItemByID(c context.Context, id string) (domain.OrderWithItemAndUser,error) {
	orderWithUserAndItem := domain.OrderWithItemAndUser{}
	purchaseItem := domain.PurchaseItem{}
	var purchaseItems []domain.PurchaseItem
	row ,err := p.rep.FindWithOrderItem(c, id)
	err = row.Scan(
		  &orderWithUserAndItem.PersonalInfo.ID,
			&orderWithUserAndItem.PersonalInfo.DeadLine,
			&orderWithUserAndItem.PersonalInfo.UserID,
			&orderWithUserAndItem.PersonalInfo.FinanceCheck,
			&orderWithUserAndItem.PersonalInfo.CreatedAt,
			&orderWithUserAndItem.PersonalInfo.UpdatedAt,
			&orderWithUserAndItem.User.ID,
			&orderWithUserAndItem.User.Name,
			&orderWithUserAndItem.User.BureauID,
			&orderWithUserAndItem.User.RoleID,
			&orderWithUserAndItem.User.CreatedAt,
			&orderWithUserAndItem.User.UpdatedAt,
	)
	if err != nil {
		return orderWithUserAndItem ,nil
	}
	rows, err := p.rep.GetPurchaseItemByOrderId(c, strconv.Itoa(int(orderWithUserAndItem.PersonalInfo.ID)))
	for rows.Next() {
		err := rows.Scan(
			&purchaseItem.ID,
			&purchaseItem.Item,
			&purchaseItem.Price,
			&purchaseItem.Quantity,
			&purchaseItem.Detail,
			&purchaseItem.Url,
			&purchaseItem.PersonalInfoID,
			&purchaseItem.FinanceCheck,
			&purchaseItem.CreatedAt,
			&purchaseItem.UpdatedAt,
		)
		if err != nil {
			return orderWithUserAndItem ,nil
		}
		purchaseItems = append(purchaseItems,purchaseItem)
	}
	orderWithUserAndItem.PurchaseItem = purchaseItems
	return orderWithUserAndItem, nil
}

//postした時に作成されたレコードの取得
func (p *personalInfoUseCase) GetPersonalInfoNewRecord(
	c context.Context,
	deadLine string,
	userID string,
	finansuCheck string,
	)(domain.PersonalInfo,error) {
	personalInfo := domain.PersonalInfo{}
	p.rep.Create(c,deadLine,userID,finansuCheck)
	row, err := p.rep.FindNewRecord(c)
	err = row.Scan(
		&personalInfo.ID,
		&personalInfo.DeadLine,
		&personalInfo.UserID,
		&personalInfo.FinanceCheck,
		&personalInfo.CreatedAt,
		&personalInfo.UpdatedAt,
	)
	if err != nil {
		return personalInfo, err
	}	
	return personalInfo,nil
}

//putした時に更新されたレコードの取得
func (p *personalInfoUseCase) GetPersonalInfoEdit(
	c context.Context,
	id string,
	deadLine string,
	userID string,
	finansuCheck string,
)(domain.PersonalInfo,error){
	personalInfo :=domain.PersonalInfo{}
	p.rep.Update(c,id,deadLine,userID,finansuCheck)
	row, err := p.rep.Find(c, id)
	err = row.Scan(
		&personalInfo.ID,
		&personalInfo.DeadLine,
		&personalInfo.UserID,
		&personalInfo.FinanceCheck,
		&personalInfo.CreatedAt,
		&personalInfo.UpdatedAt,
	)
	if err != nil {
		return personalInfo, err
	}
	return personalInfo, nil
}
