package chats

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Chat) error {
	return s.repository.Create(item)
}

func (s *Service) Update(item *models.Chat) error {
	return s.repository.Update(item)
}

func (s *Service) GetAll() (models.ChatsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(slug string) (*models.Chat, error) {
	item, err := s.repository.Get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
