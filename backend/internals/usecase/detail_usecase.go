package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type detailUseCase struct {
	rep rep.DetailRepository
}

type DetailUseCase interface {
	GetDetails(context.Context) ([]domain.Detail, error)
	GetDetailByID(context.Context, string) (domain.Detail, error)
	CreateDetail(context.Context, string) error
	UpdateDetail(context.Context, string, string) error
	DestroyDetail(context.Context, string) error
}

func NewDetailUseCase(rep rep.DetailRepository) DetailUseCase {
	return &detailUseCase{rep}
}

func (d *detailUseCase) GetDetails(c context.Context) ([]domain.Detail, error) {

	detail := domain.Detail{}
	var details []domain.Detail

	// クエリー実行
	rows, err := d.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&detail.ID,
			&detail.Name,
			&detail.CreatedAt,
			&detail.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}

		details = append(details, detail)
	}
	return details, nil
}

func (d *detailUseCase) GetDetailByID(c context.Context, id string) (domain.Detail, error) {
	var detail domain.Detail

	row, err := d.rep.Find(c, id)
	err = row.Scan(
		&detail.ID,
		&detail.Name,
		&detail.CreatedAt,
		&detail.UpdatedAt,
	)

	if err != nil {
		return detail, err
	}

	return detail, nil
}

func (d *detailUseCase) CreateDetail(c context.Context, name string) error {
	err := d.rep.Create(c, name)
	return err
}

func (d *detailUseCase) UpdateDetail(c context.Context, id string, name string) error {
	err := d.rep.Update(c, id, name)
	return err
}

func (d *detailUseCase) DestroyDetail(c context.Context, id string) error {
	err := d.rep.Destroy(c, id)
	return err
}
