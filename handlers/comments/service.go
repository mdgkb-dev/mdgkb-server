package comments

import (
	"context"
	"fmt"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(c context.Context, items models.Comments) error {
	if len(items) == 0 {
		return nil
	}
	err := R.CreateMany(c, items)
	if err != nil {
		return err
	}
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Create(c context.Context, item *models.Comment) error {
	return R.Create(c, item)
}

func (s *Service) DeleteMany(c context.Context, idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}

func (s *Service) GetAll(c context.Context) (models.CommentsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) GetAllMain(c context.Context) (models.Comments, error) {
	return R.GetAllMain(c)
}

func (s *Service) UpdateOne(c context.Context, item *models.Comment) error {
	err := R.UpdateOne(c, item)
	if err != nil {
		return err
	}
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertOne(c context.Context, item *models.Comment) error {
	err := R.UpsertOne(c, item)
	if err != nil {
		return err
	}
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	newComment, err := R.Get(c, item.ID)
	if err != nil {
		return err
	}
	s.helper.Broker.SendEvent("comment-create", newComment)
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	fmt.Println("Comment-crest")
	return nil
}
