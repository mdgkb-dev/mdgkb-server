package pagesComments

import (
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PageComments) error {
	if len(items) == 0 {
		return nil
	}
	commentsService := comments.CreateService(s.repository.getDB())
	err := commentsService.UpsertMany(items.GetComments())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.PageComments) error {
	if len(items) == 0 {
		return nil
	}
	commentsService := comments.CreateService(s.repository.getDB())
	err := commentsService.UpsertMany(items.GetComments())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
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
