package users

import (
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.CreateMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.UpsertMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.UpsertMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertEmail(item *models.User) error {
	err := s.repository.upsertEmail(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Users, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.User, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByEmail(email string) (*models.User, error) {
	item, err := s.repository.getByEmail(email)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(email string) (bool, error) {
	item, err := s.repository.emailExists(email)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) AddToUser(values map[string]interface{}, table string) error {
	err := s.repository.addToUser(values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveFromUser(values map[string]interface{}, table string) error {
	err := s.repository.removeFromUser(values, table)
	if err != nil {
		return err
	}
	return nil
}
