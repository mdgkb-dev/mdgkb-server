package search

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) SearchMain(searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = s.repository.getGroups(searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	for i := range searchModel.SearchGroups {
		err = s.repository.search(searchModel)
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
	err = s.repository.search(searchModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SearchGroups() (models.SearchGroups, error) {
	return s.repository.getGroups("")
}

func (s *Service) Search(model *models.SearchModel) error {
	model.TranslitQuery = s.helper.Util.TranslitToRu(model.Query)
	if model.Suggester {
		return s.repository.elasticSuggester(model)
	}
	err := s.repository.fullTextSearch(model)
	if err != nil {
		return err
	}
	model.BuildRoutes()
	return err
}
