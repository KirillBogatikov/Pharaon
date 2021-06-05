package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/user"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-card/data/sql"
	"pharaon-card/model"
)

type TagRepository struct {
	repo *psql.SqlRepository
}

func NewTagRepository() (*TagRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &TagRepository{repo}, nil
}

func (t *TagRepository) GetByName(name string) (*model.Tag, error) {
	rows, err := t.repo.DB.Queryx(sql.TagGetByName, &DBTag{Name: name})
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	tag := &DBTag{}
	err = rows.StructScan(tag)
	if err != nil {
		return nil, err
	}

	return &model.Tag{
		Id:   &tag.TagId,
		Name: tag.Name,
	}, nil
}

func (t *TagRepository) List() ([]model.Tag, error) {
	rows, err := t.repo.DB.Queryx(sql.TagList)
	if err != nil {
		return nil, err
	}

	result := make([]model.Tag, 0)
	for rows.Next() {
		tag := &model.Tag{}
		err = rows.StructScan(tag)
		if err != nil {
			return nil, err
		}

		result = append(result, *tag)
	}

	return result, nil
}

func (t *TagRepository) Search(query string) ([]model.Tag, error) {
	rows, err := t.repo.DB.Queryx(sql.TagSearchByName, &DBTag{Name: query})
	if err != nil {
		return nil, err
	}

	result := make([]model.Tag, 0)
	err = rows.StructScan(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (t *TagRepository) Insert(tag *model.Tag) error {
	_, err := t.repo.DB.Exec(sql.TagInsert, &DBTag{
		TagId: *tag.Id,
		Name:  tag.Name,
	})
	return err
}

func (t *TagRepository) Delete(id uuid.UUID) (bool, error) {
	result, err := t.repo.DB.Exec(sql.TagDelete, &DBTag{TagId: id})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
