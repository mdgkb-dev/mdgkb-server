package comments

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateMany(items models.Comments) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
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

func (s *Service) GetAll() (models.CommentsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) GetAllMain() (models.Comments, error) {
	return s.repository.getAllMain()
}

func (s *Service) UpdateOne(item *models.Comment) error {
	err := s.repository.updateOne(item)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertOne(item *models.Comment) error {
	err := s.repository.upsertOne(item)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	newComment, err := s.repository.get(item.ID)
	if err != nil {
		return err
	}
	s.helper.Broker.SendEvent("comment-create", newComment)
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	fmt.Println("Comment-crest")
	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
