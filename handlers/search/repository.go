package search

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

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

func (r *Repository) fullTextSearch(model *models.SearchModel) error {
	query := "phraseto_tsquery('russian', '\"?\"')"
	rank := fmt.Sprintf("ts_rank_cd(search_column, %s) as rank", query)
	q := r.db().NewSelect().Column("label", "description", "value", "key").
		ColumnExpr(rank, bun.Safe(model.Query)).
		Table("search_elements").
		Where("search_column @@ "+query, bun.Safe(model.Query))
	if model.SearchGroup.ID.Valid {
		q.Where("key = '?'", bun.Safe(model.SearchGroup.Key))
	}
	q.Order("rank desc")
	err := q.Scan(r.ctx, &model.SearchElements)
	return err
}

func (r *Repository) elasticSuggester(model *models.SearchModel) (err error) {
	q := r.db().NewSelect().Column("label").Table("search_items").
		Where(`search_column @@ to_tsquery('russian', '"?"')`, bun.Safe(model.Query))
	//if model.SearchGroup != nil {
	//	q.Where("search_group", bun.Safe(model.SearchGroup.Key))
	//}
	err = q.Scan(r.ctx, &model.SearchElements)
	return err
}
