package search

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
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

func (r *Repository) getGroupByKey(key string) (*models.SearchGroup, error) {
	item := models.SearchGroup{}
	query := r.db().NewSelect().Model(&item).
		Relation("SearchGroupMetaColumns").Where("key = ?", key)
	err := query.Scan(r.ctx)
	return &item, err
}

func (r *Repository) search(searchModel *models.SearchModel) error {
	querySelect := fmt.Sprintf("SELECT %s.%s as value, substring(%s for 40) as label", searchModel.SearchGroup.Table, searchModel.SearchGroup.ValueColumn, searchModel.SearchGroup.LabelColumn)
	queryFrom := fmt.Sprintf("FROM %s", searchModel.SearchGroup.Table)
	join := ""

	condition := fmt.Sprintf("where replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ' , '') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+searchModel.Query+"%'")
	conditionTranslitToRu := fmt.Sprintf("or replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ', '') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToRu(searchModel.Query)+"%'")
	conditionTranslitToEng := fmt.Sprintf("or replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ', '') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToEng(searchModel.Query)+"%'")

	queryOrder := fmt.Sprintf("ORDER BY %s", searchModel.SearchGroup.LabelColumn)
	query := fmt.Sprintf("%s %s %s %s %s %s %s", querySelect, queryFrom, join, condition, conditionTranslitToRu, conditionTranslitToEng, queryOrder)
	rows, err := r.db().QueryContext(r.ctx, query)
	if err != nil {
		return err
	}

	err = r.db().ScanRows(r.ctx, rows, &searchModel.SearchGroup.SearchElements)
	return err
}

type Patname struct {
	bun.BaseModel `bun:"lexemes,alias:lexemes"`
	Pos           int
	Lexeme        string
}

func (r *Repository) fullTextSearch(model *models.SearchModel) (err error) {
	query := strings.ToLower(model.Query)
	src1 := r.db().NewSelect().
		ColumnExpr("lexeme").
		ColumnExpr("positions[1] as pos").
		TableExpr(fmt.Sprintf("unnest(to_tsvector('russian','%s'))", query))

	subq := r.db().NewSelect().
		ColumnExpr("f.pos").
		ColumnExpr("'('||string_agg(l.lexeme,'|')||')' as tq").
		TableExpr("fio as f").Join("join lexemes as l on l.lexeme % f.lexeme").Group("f.pos")

	src2 := r.db().NewSelect().
		ColumnExpr("to_tsquery('simple', string_agg(q.tq,'|')) as q").
		TableExpr("(?) AS q", subq)

	q := r.db().NewSelect().With("fio", src1).With("query", src2).
		ColumnExpr("ts_rank_cd(search_column, ?) as rank1, search_column  <=> (select q from query) as rank2", query).
		Column("label", "description", "value", "key").
		Table("search_elements").
		Where("search_column @@ (select q from query)")
	if model.SearchGroup.ID.Valid {
		q.Where("key = '?'", bun.Safe(model.SearchGroup.Key))
	}
	q.OrderExpr("rank1 desc, rank2 asc")

	model.Pagination.CreatePagination(q)
	model.Count, err = q.ScanAndCount(r.ctx, &model.SearchElements)

	return err
}

func (r *Repository) elasticSuggester(model *models.SearchModel) (err error) {
	query := strings.ToLower(model.Query)
	src1 := r.db().NewSelect().
		ColumnExpr("lexeme").
		ColumnExpr("positions[1] as pos").
		TableExpr(fmt.Sprintf("unnest(to_tsvector('russian','%s'))", query))

	subq := r.db().NewSelect().
		ColumnExpr("f.pos").
		ColumnExpr("'('||string_agg(l.lexeme,'|')||')' as tq").
		TableExpr("fio as f").Join("join lexemes as l on l.lexeme % f.lexeme").Group("f.pos")

	src2 := r.db().NewSelect().
		ColumnExpr("to_tsquery('simple', string_agg(q.tq,'|')) as q").
		TableExpr("(?) AS q", subq)

	q := r.db().NewSelect().With("fio", src1).With("query", src2).
		ColumnExpr("ts_rank_cd(search_column, ?) as rank1, search_column  <=> (select q from query) as rank2", query).
		Column("label", "description", "value", "key").
		Table("search_elements").
		Where("search_column @@ (select q from query)")
	if model.SearchGroup.ID.Valid {
		q.Where("key = '?'", bun.Safe(model.SearchGroup.Key))
	}
	q.OrderExpr("rank1 desc, rank2 asc")

	r.queryFilter.HandleQuery(q)
	q.OrderExpr("rank1 desc, rank2 asc")
	q.Limit(10)
	err = q.Scan(r.ctx, &model.SearchElements)
	return err
}
