package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
	"github.com/pkg/errors"
)

type tagUseCase struct {
	rep rep.TagRepository
}

type TagUseCase interface {
	GetTags(context.Context) ([]domain.Tag, error)
	GetTagByID(context.Context, string) (domain.Tag, error)
	CreateTag(context.Context, string) error
	UpdateTag(context.Context, string, string) error
	DestroyTag(context.Context, string) error
}

func NewTagUseCase(rep rep.TagRepository) TagUseCase {
	return &tagUseCase{rep}
}

func (s *tagUseCase) GetTags(c context.Context) ([]domain.Tag, error) {

	tag := domain.Tag{}
	var tags []domain.Tag

	// クエリー実行
	rows, err := s.rep.All(c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)

		if err != nil {
			return nil, errors.Wrapf(err, "cannot connect SQL")
		}

		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *tagUseCase) GetTagByID(c context.Context, id string) (domain.Tag, error) {
	var tag domain.Tag

	row, err := s.rep.Find(c, id)
	err = row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)

	if err != nil {
		return tag, err
	}

	return tag, nil
}

func (s *tagUseCase) CreateTag(c context.Context, name string) error {
	err := s.rep.Create(c, name)
	return err
}

func (s *tagUseCase) UpdateTag(c context.Context, id string, name string) error {
	err := s.rep.Update(c, id, name)
	return err
}

func (s *tagUseCase) DestroyTag(c context.Context, id string) error {
	err := s.rep.Destroy(c, id)
	return err
}
