package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type classificationUseCase struct {
	rep rep.ClassificationRepository
}

type ClassificationUseCase interface {
	GetClassifications(context.Context) ([]domain.Classification, error)
	GetClassificationByID(context.Context, string) (domain.Classification, error)
	CreateClassification(context.Context, string, string, string) error
	UpdateClassification(context.Context, string, string, string, string) error
	DestroyClassification(context.Context, string) error
	//classificationsに紐づくyearとsourceの取得
	GetClassificationWithYearAndSource(context.Context, string) (domain.ClassificationYearSource, error)
}

func NewClassificationUseCase(rep rep.ClassificationRepository) ClassificationUseCase {
	return &classificationUseCase{rep}
}

//classificationsの取得(Gets)
func (b *classificationUseCase) GetClassifications(c context.Context) ([]domain.Classification, error) {

	classification := domain.Classification{}
	var classifications []domain.Classification

	// クエリー実行
	rows, err := b.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&classification.ID,
			&classification.Price,
			&classification.YearID,
			&classification.SourceID,
			&classification.CreatedAt,
			&classification.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}

		classifications = append(classifications, classification)
	}
	return classifications, nil
}

//classificationの取得(Get)
func (b *classificationUseCase) GetClassificationByID(c context.Context, id string) (domain.Classification, error) {
	var classification domain.Classification

	row, err := b.rep.Find(c, id)
	err = row.Scan(
		&classification.ID,
		&classification.Price,
		&classification.YearID,
		&classification.SourceID,
		&classification.CreatedAt,
		&classification.UpdatedAt,
	)

	if err != nil {
		return classification, err
	}

	return classification, nil
}

//classificationの作成(Create)
func (b *classificationUseCase) CreateClassification(c context.Context, price string, yearID string, sourceID string) error {
	err := b.rep.Create(c, price, yearID, sourceID)
	return err
}

//classificationの更新(Update)
func (b *classificationUseCase) UpdateClassification(c context.Context, id string, price string, yearID string, sourceID string) error {
	err := b.rep.Update(c, id, price, yearID, sourceID)
	return err
}

//classificationの削除(Delete)
func (b *classificationUseCase) DestroyClassification(c context.Context, id string) error {
	err := b.rep.Destroy(c, id)
	return err
}

//classificationに紐づくyearとsourceの取得(Get)
func (b *classificationUseCase) GetClassificationWithYearAndSource(c context.Context, id string) (domain.ClassificationYearSource, error) {
	var classificationyearsource domain.ClassificationYearSource

	row, err := b.rep.FindYearAndSource(c, id)
	err = row.Scan(
		&classificationyearsource.Classification.ID,
		&classificationyearsource.Classification.Price,	
		&classificationyearsource.Classification.YearID,	
		&classificationyearsource.Classification.SourceID,	
		&classificationyearsource.Classification.CreatedAt,
		&classificationyearsource.Classification.UpdatedAt,
		&classificationyearsource.Year.ID,
		&classificationyearsource.Year.Year,
		&classificationyearsource.Year.CreatedAt,
		&classificationyearsource.Year.UpdatedAt,
		&classificationyearsource.Source.ID,
		&classificationyearsource.Source.Name,
		&classificationyearsource.Source.CreatedAt,
		&classificationyearsource.Source.UpdatedAt,
	)
	if err != nil {
		return classificationyearsource, err
	}
	return classificationyearsource , nil
}