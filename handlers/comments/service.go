package comments

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.Comments) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.Comments) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}

func (s *Service) GetAll(params *commentsParams) (models.Comments, error) {
	items, err := s.repository.getAll(params)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateOne(item *models.Comment) error {
	return s.repository.updateOne(item)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
