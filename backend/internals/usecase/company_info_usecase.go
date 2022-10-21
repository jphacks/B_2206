package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type companyInfoUseCase struct {
	rep rep.CompanyInfoRepository
}

type CompanyInfoUseCase interface {
	GetCompanyInfos(context.Context) ([]domain.CompanyInfo,error)
	GetCompanyInfoByID(context.Context,string) (domain.CompanyInfo,error)
	CreateCompanyInfo(context.Context,string) error
	UpdateCompanyInfo(context.Context,string,string) error
	DestroyCompanyInfo(context.Context, string) error
}

func NewCompanyInfoUseCase(rep rep.CompanyInfoRepository) CompanyInfoUseCase {
	return &companyInfoUseCase{rep}
}

func (b *companyInfoUseCase) GetCompanyInfos(c context.Context) ([]domain.CompanyInfo,error) {
	companyInfo := domain.CompanyInfo{}
	var companyInfos []domain.CompanyInfo

	//クエリ実行
	rows,err := b.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&companyInfo.ID,
			&companyInfo.Name,
			&companyInfo.CreatedAt,
			&companyInfo.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "can not connect SQL")
		}
		companyInfos = append(companyInfos, companyInfo)
	}
	return companyInfos,nil
}

func (b *companyInfoUseCase) GetCompanyInfoByID(c context.Context, id string) (domain.CompanyInfo, error) {
	var companyInfo domain.CompanyInfo
	row, err := b.rep.Find(c, id)
	err = row.Scan(
		&companyInfo.ID,
		&companyInfo.Name,
		&companyInfo.CreatedAt,
		&companyInfo.UpdatedAt,
	)
	if err != nil {
		return companyInfo,err
	}
	return companyInfo,nil
}  

func (b *companyInfoUseCase) CreateCompanyInfo(c context.Context, name string) error {
	err := b.rep.Create(c,name)
	return err
}

func (b *companyInfoUseCase) UpdateCompanyInfo(c context.Context, id string, name string) error {
	err := b.rep.Update(c ,id, name)
	return err
}

func (b *companyInfoUseCase) DestroyCompanyInfo(c context.Context, id string) error{
	err := b.rep.Destroy(c, id)
	return err
}