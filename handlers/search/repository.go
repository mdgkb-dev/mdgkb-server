package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getGroups(groupID string) (models.SearchGroups, error) {
	items := make(models.SearchGroups, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("SearchGroupMetaColumns").
		Order("search_group_order")

	if groupID != "" {
		query = query.Where("id = ?", groupID)
	}
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) search(searchModel *models.SearchModel) error {
	querySelect := fmt.Sprintf("SELECT id, %s as value, %s as label", searchModel.SearchGroup.ValueColumn, searchModel.SearchGroup.LabelColumn)
	queryFrom := fmt.Sprintf("FROM %s", searchModel.SearchGroup.Table)
	search := searchModel.Query
	if searchModel.MustBeTranslate {
		search = r.helper.Util.TranslitToRu(searchModel.Query)
	}
	queryWhere := r.helper.SQL.WhereLikeWithLowerTranslit(searchModel.SearchGroup.SearchColumn, search)
	query := fmt.Sprintf("%s %s %s", querySelect, queryFrom, queryWhere)
	rows, err := r.db().QueryContext(r.ctx, query)
	if err != nil {
		return err
	}
	err = r.db().ScanRows(r.ctx, rows, &searchModel.SearchGroup.SearchElements)
	return err
}

func (r *Repository) elasticSearch(model *models.SearchModel) error {
	var data map[string]interface{}
	query, indexes := model.BuildQuery()
	res, err := r.elasticsearch.Search(
		r.elasticsearch.Search.WithIndex(indexes...),
		r.elasticsearch.Search.WithBody(esutil.NewJSONReader(&query)),
		r.elasticsearch.Search.WithPretty(),
	)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}
	model.ParseMap(data)
	return nil
}

func (r *Repository) elasticSuggester(model *models.SearchModel) error {
	var re map[string]interface{}
	//indexes := []string{}
	//if model.SearchGroup != nil {
	//	indexes = append(indexes, model.SearchGroup.Table)
	//}
	should := make([]interface{}, 0)
	should = append(should, map[string]interface{}{
		"prefix": map[string]interface{}{
			"name": map[string]interface{}{
				"value": model.Query,
			},
		},
	})
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": should[0],
	}
	_ = json.NewEncoder(&buf).Encode(query)
	res, err := r.elasticsearch.Search(
		//r.elasticsearch.Search.WithIndex(indexes...),
		r.elasticsearch.Search.WithBody(&buf),
		//r.elasticsearch.Get.
		r.elasticsearch.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&re)
	if err != nil {
		return err
	}
	model.ParseMap(re)
	return nil
}
