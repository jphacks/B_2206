package usecase

import (
	"context"

	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type areaUseCase struct {
	rep rep.AreaRepository
}

type AreaUseCase interface {
	GetActivities(context.Context) ([]domain.Area, error)
	GetAreaByID(context.Context, string) (domain.Area, error)
	CreateArea(context.Context, string, string, string, string, string) error
	UpdateArea(context.Context, string, string, string, string, string) error
	DestroyArea(context.Context, string) error
	//GetActivitiesWithSponsorAndStyle(context.Context) ([]domain.AreaForAdminView, error)
}

func NewAreaUseCase(rep rep.AreaRepository) AreaUseCase {
	return &areaUseCase{rep}
}

func (a *areaUseCase) GetActivities(c context.Context) ([]domain.Area, error) {

	area := domain.Area{}
	var activities []domain.Area

	// クエリー実行
	rows, err := a.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&area.ID,
			&area.CreatedAt,
			&area.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}

		activities = append(activities, area)
	}
	return activities, nil
}

func (a *areaUseCase) GetAreaByID(c context.Context, id string) (domain.Area, error) {
	var area domain.Area

	row, err := a.rep.Find(c, id)
	err = row.Scan(
		&area.ID,
		&area.PostCode,
		&area.Prefecture,
		&area.City,
		&area.AddressNumber,
		&area.BuildingName,
		&area.CreatedAt,
		&area.UpdatedAt,
	)

	if err != nil {
		return area, err
	}

	return area, nil
}

func (a *areaUseCase) CreateArea(c context.Context, PostCode string, Prefecture string, City string, AddressNumber string, BuildingName string) error {
	err := a.rep.Create(c, PostCode, Prefecture, City, AddressNumber, BuildingName)
	return err
}

func (a *areaUseCase) UpdateArea(c context.Context, id string, PostCode string, Prefecture string, City string, AddressNumber string, BuildingName string) error {
	err := a.rep.Update(c, id, PostCode, Prefecture, City, AddressNumber, BuildingName)
	return err
}

func (a *areaUseCase) DestroyArea(c context.Context, id string) error {
	err := a.rep.Destroy(c, id)
	return err
}

// func (a *areaUseCase) GetActivitiesWithSponsorAndStyle(c context.Context) ([]domain.AreaForAdminView, error) {

// 	area := domain.AreaForAdminView{}
// 	var activities []domain.AreaForAdminView

// 	// クエリー実行
// 	rows, err := a.rep.AllWithSponsor(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(
// 			&area.Area.ID,
// 			&area.Area.SponsorStyleID,
// 			&area.Area.UserID,
// 			&area.Area.IsDone,
// 			&area.Area.SponsorID,
// 			&area.Area.CreatedAt,
// 			&area.Area.UpdatedAt,
// 			&area.Sponsor.ID,
// 			&area.Sponsor.Name,
// 			&area.Sponsor.Tel,
// 			&area.Sponsor.Email,
// 			&area.Sponsor.Address,
// 			&area.Sponsor.Representative,
// 			&area.Sponsor.CreatedAt,
// 			&area.Sponsor.UpdatedAt,
// 			&area.SponsorStyle.ID,
// 			&area.SponsorStyle.Scale,
// 			&area.SponsorStyle.IsColor,
// 			&area.SponsorStyle.Price,
// 			&area.SponsorStyle.CreatedAt,
// 			&area.SponsorStyle.UpdatedAt,
// 			&area.User.ID,
// 			&area.User.Name,
// 			&area.User.BureauID,
// 			&area.User.RoleID,
// 			&area.User.CreatedAt,
// 			&area.User.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, errors.Wrapf(err, "cannot connect SQL")
// 		}

// 		activities = append(activities, area)
// 	}
// 	return activities, nil
// }
