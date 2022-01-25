package httpHelper

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type QueryFilter struct {
	ID           *string
	FilterModels FilterModels
	SortModels   SortModels
	Pagination   *Pagination
}

func (i *QueryFilter) CreateFilter(query *bun.SelectQuery) {
	createFilter(query, i.FilterModels)
}

func (i *HTTPHelper) CreateQueryFilter(c *gin.Context) (*QueryFilter, error) {
	filterModels, err := createFilterModels(c)
	if err != nil {
		return nil, err
	}
	sortModels, err := createSortModels(c)
	if err != nil {
		return nil, err
	}
	pagination, err := createPagination(c)
	if err != nil {
		return nil, err
	}
	id := c.Param("id")
	return &QueryFilter{ID: &id, FilterModels: filterModels, SortModels: sortModels, Pagination: pagination}, nil
}

func createSortModels(c *gin.Context) (SortModels, error) {
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

func createFilterModels(c *gin.Context) (FilterModels, error) {
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

func createPagination(c *gin.Context) (*Pagination, error) {
	return parseJSONToPagination(c.Query("pagination"))
}
