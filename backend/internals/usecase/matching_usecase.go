package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type matchingUseCase struct {
	rep rep.MatchingRepository
}

type MatchingUseCase interface {
	GetMatching(context.Context) ([]domain.Matching, error)
	GetMatchingByID(context.Context, string) (domain.Matching, error)
	CreateMatching(context.Context, string, string, string, string, string, string, string) error
	UpdateMatching(context.Context, string, string, string, string, string, string, string, string) error
	DestroyMatching(context.Context, string) error
	GetMatchingWithPurchaseOrder(context.Context) ([]domain.MatchingWithOrder, error)
	GetMatchingWithPurchaseOrderByID(context.Context, string) (domain.MatchingWithOrder, error)
}

func NewMatchingUseCase(rep rep.MatchingRepository) MatchingUseCase {
	return &matchingUseCase{rep}
}

//Matchingsの取得(Gets)
func (p *matchingUseCase) GetMatching(c context.Context) ([]domain.Matching, error) {
	matching := domain.Matching{}
	var matchings []domain.Matching
	rows, err := p.rep.All(c)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		err := rows.Scan(
			&matching.ID,
			&matching.Item,
			&matching.Price,
			&matching.Quantity,
			&matching.Detail,
			&matching.Url,
			&matching.PurchaseOrderID,
			&matching.FinanceCheck,
			&matching.CreatedAt,
			&matching.UpdatedAt,
		)
		if err != nil {
			return nil , err
		}
		matchings = append(matchings, matching)
	}
	return matchings, nil
}

//matchingの取得(Get)
func(p *matchingUseCase) GetMatchingByID(c context.Context, id string) (domain.Matching, error){
	matching := domain.Matching{}
	row, err := p.rep.Find(c, id)
	err = row.Scan(
			&matching.ID,
			&matching.Item,
			&matching.Price,
			&matching.Quantity,
			&matching.Detail,
			&matching.Url,
			&matching.PurchaseOrderID,
			&matching.FinanceCheck,
			&matching.CreatedAt,
			&matching.UpdatedAt,
	)
	if err != nil {
		return matching, err
	}
	return matching, nil
}


//matchingの作成
func(p *matchingUseCase) CreateMatching(
	c context.Context,
	Item string,
	Price string,
	Quantity string,
	Detail string,
	Url string,
	PurchaseOrderID string,
	FinanceCheck string,
)error {
	err := p.rep.Create(c, Item, Price, Quantity, Detail, Url, PurchaseOrderID, FinanceCheck)
	return err
}

//matchingの修正(Update)
func(p *matchingUseCase) UpdateMatching(
	c context.Context,
	id string,
	Item string,
	Price string,
	Quantity string,
	Detail string,
	Url string,
	PurchaseOrderID string,
	FinanceCheck string,
)error {
	err := p.rep.Update(c, id, Item, Price, Quantity, Detail, Url, PurchaseOrderID, FinanceCheck)
	return err
}

//matchingの削除(delete)
func(p *matchingUseCase) DestroyMatching(
	c context.Context,
	id string,
)error {
	err := p.rep.Delete(c, id)
	return err
}

//purchaseOrderに紐づくMatchingの取得(Gets)
func (p *matchingUseCase) GetMatchingWithPurchaseOrder(c context.Context) ([]domain.MatchingWithOrder, error) {

	matchingwithpurchaseorder := domain.MatchingWithOrder{}
	var matchingwithpurchaseorders []domain.MatchingWithOrder

	//クエリ実行
	rows, err := p.rep.AllWithPurchaseOrder(c)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(
			&matchingwithpurchaseorder.ID,
			&matchingwithpurchaseorder.Item,
			&matchingwithpurchaseorder.Price,
			&matchingwithpurchaseorder.Quantity,
			&matchingwithpurchaseorder.Detail,
			&matchingwithpurchaseorder.Url,
			&matchingwithpurchaseorder.DeadLine,
			&matchingwithpurchaseorder.Name,
			&matchingwithpurchaseorder.FinanceCheck,
			&matchingwithpurchaseorder.CreatedAt,
			&matchingwithpurchaseorder.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}
		matchingwithpurchaseorders = append(matchingwithpurchaseorders, matchingwithpurchaseorder)
	}
	return matchingwithpurchaseorders, nil
}

//purchaseOrderに紐づくMatchingの取得(Get)
func (p *matchingUseCase) GetMatchingWithPurchaseOrderByID(c context.Context, id string) (domain.MatchingWithOrder, error){
	matchingwithpurchaseorder := domain.MatchingWithOrder{}

	row, err := p.rep.FindWithPurchaseOrder(c, id)
	err = row.Scan(
		&matchingwithpurchaseorder.ID,
		&matchingwithpurchaseorder.Item,
		&matchingwithpurchaseorder.Price,
		&matchingwithpurchaseorder.Quantity,
		&matchingwithpurchaseorder.Detail,
		&matchingwithpurchaseorder.Url,
		&matchingwithpurchaseorder.DeadLine,
		&matchingwithpurchaseorder.Name,
		&matchingwithpurchaseorder.FinanceCheck,
		&matchingwithpurchaseorder.CreatedAt,
		&matchingwithpurchaseorder.UpdatedAt,
	)
	if err != nil {
		return matchingwithpurchaseorder, err
	}
	return matchingwithpurchaseorder, nil

}

