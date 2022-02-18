package search

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) SearchMain(searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = s.repository.getGroups(searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	search := s.helper.TranslitToRu(searchModel.Query)
	for i := range searchModel.SearchGroups {
		err = s.repository.search(searchModel.SearchGroups[i], search)
		if err != nil {
			return err
		}
		searchModel.SearchGroups[i].BuildRoutes()
	}
	return nil
}

func (s *Service) SearchObjects(searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = s.repository.getGroups(searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	search := s.helper.TranslitToRu(searchModel.Query)
	err = s.repository.search(searchModel.SearchGroup, search)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SearchGroups() (models.SearchGroups, error) {
	return s.repository.getGroups("")
}
