package search

import (
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/models"
	"os"
	"path/filepath"
)

func (s *Service) SearchMain(searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = s.repository.getGroups(searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	search := s.helper.Util.TranslitToRu(searchModel.Query)
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
	fmt.Println(searchModel.SearchGroupID)
	fmt.Println(searchModel.SearchGroupID)
	fmt.Println(searchModel.SearchGroupID)
	fmt.Println(searchModel.SearchGroupID)
	fmt.Println(searchModel.SearchGroupID)
	searchModel.SearchGroups, err = s.repository.getGroups(searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	search := s.helper.Util.TranslitToRu(searchModel.Query)
	err = s.repository.search(searchModel.SearchGroup, search)
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
	if !s.helper.Search.On {
		return dummy(model)
	}
	if model.Suggester {
		return s.repository.elasticSuggester(model)
	}
	return s.repository.elasticSearch(model)
}

func dummy(model *models.SearchModel) error {
	var re map[string]interface{}
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path = filepath.Join(path, "dummy")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&re)
	if err != nil {
		return err
	}
	model.ParseMap(re)
	return nil
}
