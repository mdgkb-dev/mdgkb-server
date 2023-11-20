package search

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
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
	searchModel.Query = s.helper.Util.TranslitToRu(searchModel.Query)
	searchModel.SearchGroup, err = s.repository.getGroupByKey(searchModel.Key)
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

func (s *Service) Search(model *models.SearchModel) (err error) {
	model.TranslitQuery = s.helper.Util.TranslitToRu(model.Query)
	model.SearchGroup, err = s.repository.getGroupByKey(model.Key)
	if err != nil {
		return err
	}
	if model.Suggester {
		return s.repository.elasticSuggester(model)
	}
	err = s.repository.fullTextSearch(model)
	if err != nil {
		return err
	}
	model.BuildRoutes()
	return err
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
