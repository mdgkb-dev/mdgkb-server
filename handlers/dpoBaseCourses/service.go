package dpoBaseCourses

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.DpoBaseCourses, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.DpoBaseCourse, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.DpoBaseCourse) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DpoBaseCourse) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
