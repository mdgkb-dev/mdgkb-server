package search

import (
	"bytes"
	"encoding/json"
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
	search := searchModel.Query
	if searchModel.MustBeTranslate {
		search = r.helper.Util.TranslitToRu(searchModel.Query)
	}
	err := r.db().
		NewSelect().
		Model(&searchModel.SearchGroup.SearchElements).
		ColumnExpr("id, ? as value, ? as label", bun.Ident(searchModel.SearchGroup.ValueColumn), bun.Ident(searchModel.SearchGroup.LabelColumn)).
		ModelTableExpr(searchModel.SearchGroup.Table).
		Where("lower(regexp_replace(?, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", bun.Ident(searchModel.SearchGroup.SearchColumn), "%"+search+"%").
		Order(searchModel.SearchGroup.SearchColumn).
		Scan(r.ctx)
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
