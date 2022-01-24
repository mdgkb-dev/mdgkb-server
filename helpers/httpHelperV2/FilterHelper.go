package httpHelper

import (
	"github.com/gin-gonic/gin"
)

type QueryFilter struct {
	ID           *string
	FilterModels FilterModels
	SortModels   SortModels
	Pagination   *Pagination
}

func (i *HTTPHelper) CreateQueryFilter(c *gin.Context) (*QueryFilter, error) {
	filterModels, err := CreateFilterModels(c)
	if err != nil {
		return nil, err
	}
	sortModels, err := CreateSortModels(c)
	if err != nil {
		return nil, err
	}
	pagination, err := CreatePagination(c)
	if err != nil {
		return nil, err
	}
	id := c.Param("id")
	return &QueryFilter{ID: &id, FilterModels: filterModels, SortModels: sortModels, Pagination: pagination}, nil
}

func CreateSortModels(c *gin.Context) (SortModels, error) {
	sortModels := make(SortModels, 0)
	if c.Query("sortModel") == "" {
		return nil, nil
	}
	for _, arg := range c.QueryArray("sortModel") {
		sortModel, err := ParseJSONToSortModel(arg)
		if err != nil {
			return nil, err
		}
		sortModels = append(sortModels, &sortModel)
	}

	return sortModels, nil
}

func CreateFilterModels(c *gin.Context) (FilterModels, error) {
	filterModels := make(FilterModels, 0)
	if c.Query("filterModel") == "" {
		return nil, nil
	}
	for _, arg := range c.QueryArray("filterModel") {
		filterModel, err := ParseJSONToFilterModel(arg)
		if err != nil {
			return nil, err
		}
		filterModels = append(filterModels, &filterModel)
	}

	return filterModels, nil
}
