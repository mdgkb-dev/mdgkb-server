package specializations

import "mdgkb/mdgkb-server/models"

func (s *Service) GetAll() (models.Specializations, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Specialization, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.Specialization) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Specialization) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.Specializations) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(id []string) error {
	if len(id) == 0 {
		return nil
	}
	return s.repository.deleteMany(id)
}
