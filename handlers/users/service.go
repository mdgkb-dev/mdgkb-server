package users

import (
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.User) error {
	//err := fileInfos.CreateService(s.repository.getDB()).Create(item.FileInfo)
	//if err != nil {
	//	return err
	//}
	err := human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	//item.SetIdForChildren()
	return nil
}

//
//func (s *Service) Update(item *models.User) error {
//	err := fileInfos.CreateService(s.repository.getDB()).Upsert(item.FileInfo)
//	if err != nil {
//		return err
//	}
//	err = human.CreateService(s.repository.getDB()).Update(item.Human)
//	if err != nil {
//		return err
//	}
//	err = timetables.CreateService(s.repository.getDB()).Upsert(item.Timetable)
//	if err != nil {
//		return err
//	}
//	item.SetForeignKeys()
//	err = s.repository.update(item)
//	if err != nil {
//		return err
//	}
//	item.SetIdForChildren()
//	UserRegaliaService := UserRegalia.CreateService(s.repository.getDB())
//	err = UserRegaliaService.UpsertMany(item.UserRegalias)
//	if err != nil {
//		return err
//	}
//	err = UserRegaliaService.DeleteMany(item.UserRegaliasForDelete)
//	if err != nil {
//		return err
//	}
//	educationsService := educations.CreateService(s.repository.getDB())
//	err = educationsService.UpsertMany(item.Educations)
//	if err != nil {
//		return err
//	}
//	err = educationsService.DeleteMany(item.EducationsForDelete)
//	if err != nil {
//		return err
//	}
//	return nil
//}

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
