package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
)

type valueUseCase struct{
	rep rep.ValueRepository
}

type ValueUseCase interface {
	GetValues(context.Context) ([]domain.Value, error)
	GetValuesByID(context.Context, string) (domain.Value, error)
	CreateValue(context.Context, string, string, string) error
	UpdateValue(context.Context, string, string, string, string) error
	DestroyValue(context.Context, string) error
}

func NewValueUseCase(rep rep.ValueRepository) ValueUseCase {
	return &valueUseCase{rep}
}

//Sponsorstylesの取得(Gets)
func (s *valueUseCase) GetValues(c context.Context) ([]domain.Value, error) {
	value :=domain.Value{}
	var values []domain.Value
	rows, err := s.rep.All(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&value.ID,
			&value.Scale,
			&value.IsColor,
			&value.Price,
			&value.CreatedAt,
			&value.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

//Valueの取得(Get)
func (s *valueUseCase) GetValuesByID(c context.Context, id string) (domain.Value, error){
	value := domain.Value{}
	row , err := s.rep.Find(c, id)
	err = row.Scan(
		&value.ID,
		&value.Scale,
		&value.IsColor,
		&value.Price,
		&value.CreatedAt,
		&value.UpdatedAt,
	)
	if err != nil {
		return value, err
	}
	return value, nil
}

//Valueの作成(Create)
func (s *valueUseCase) CreateValue(
	c context.Context, 
	Scale string,
	IsColor string,
	Price string,
)error {
	err := s.rep.Create(c, Scale, IsColor, Price)
	return err
}
//Valueの編集(Update)
func (s *valueUseCase) UpdateValue(
	c context.Context,
	id string,
	Scale string,
	IsColor string,
	Price string,
)error {
	err := s.rep.Update(c, id, Scale, IsColor, Price)
	return err
}

//Valueの削除(Delate)
func (s *valueUseCase) DestroyValue(
	c context.Context,
	id string,
)error{
	err := s.rep.Delete(c, id)
	return err
}