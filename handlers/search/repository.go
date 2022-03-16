package search

import (
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/models"
	"os"
	"path/filepath"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getGroups(groupID string) (models.SearchGroups, error) {
	items := make(models.SearchGroups, 0)
	query := r.db.NewSelect().Model(&items).Order("search_group_order")

	if groupID != "" {
		query = query.Where("id = ?", groupID)
	}
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) search(searchGroup *models.SearchGroup, search string) error {
	querySelect := fmt.Sprintf("SELECT %s as value, %s as label", searchGroup.ValueColumn, searchGroup.LabelColumn)
	queryFrom := fmt.Sprintf("FROM %s", searchGroup.Table)
	queryWhere := r.helper.SQL.WhereLikeWithLowerTranslit(searchGroup.SearchColumn, search)
	query := fmt.Sprintf("%s %s %s", querySelect, queryFrom, queryWhere)
	rows, err := r.db.QueryContext(r.ctx, query)
	if err != nil {
		return err
	}
	err = r.db.ScanRows(r.ctx, rows, &searchGroup.SearchElements)
	return err
}

func (r *Repository) elasticSearch(model *models.SearchModel) error {
	var re map[string]interface{}
	if r.helper.Search.On {
		res, err := r.elasticsearch.Search(
			r.elasticsearch.Search.WithIndex("divisions"),
			r.elasticsearch.Search.WithPretty(),
		)
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&re)
		if err != nil {
			return err
		}
	} else {
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
	}
	model.ParseMap(re)
	return nil
}
