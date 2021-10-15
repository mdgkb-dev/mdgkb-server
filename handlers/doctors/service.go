package doctors

import (
	"mdgkb/mdgkb-server/handlers/doctorRegalia"
	"mdgkb/mdgkb-server/handlers/educations"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Doctor) error {
	err := fileInfos.CreateService(s.repository.getDB()).Create(item.FileInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.getDB()).Create(item.Human)
	if err != nil {
		return err
	}

	item.SetForeignKeys()
	err = s.Create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = doctorRegalia.CreateService(s.repository.getDB()).CreateMany(item.DoctorRegalias)
	if err != nil {
		return err
	}
	err = educations.CreateService(s.repository.getDB()).CreateMany(item.Educations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Doctor) error {
	err := fileInfos.CreateService(s.repository.getDB()).Upsert(item.FileInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.getDB()).Update(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.Update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = doctorRegalia.CreateService(s.repository.getDB()).UpsertMany(item.DoctorRegalias)
	if err != nil {
		return err
	}
	err = educations.CreateService(s.repository.getDB()).UpsertMany(item.Educations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Doctors, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Doctor, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByDivisionID(divisionID string) (models.Doctors, error) {
	return s.repository.getByDivisionID(divisionID)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateComment(item *models.DoctorComment) error {
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DoctorComment) error {
	return s.repository.updateComment(item)
}

func (s *Service) RemoveComment(id string) error {
	return s.repository.removeComment(id)
}
