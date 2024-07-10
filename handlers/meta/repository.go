package meta

import (
	"context"
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getCount(table *string) (res *int, err error) {
	num := 0
	err = r.db().QueryRow("SELECT COUNT (id) FROM ?", *table).Scan(&num)
	return &num, err
}

func (r *Repository) getOptions(optionModel *models.OptionModel) (models.Options, error) {
	options := make(models.Options, 0)
	query := "SELECT ?::varchar as value, ? as label FROM ? ORDER BY ?"
	queryContext, err := r.db().Query(query, bun.Ident(optionModel.Value), bun.Ident(optionModel.Label), bun.Ident(optionModel.TableName), bun.Ident(optionModel.SortColumn))
	if err != nil {
		return nil, err
	}
	err = r.db().ScanRows(context.TODO(), queryContext, &options)
	return options, err
}

func (r *Repository) getApplicationsCounts() (models.ApplicationsCounts, error) {
	items := make(models.ApplicationsCounts, 0)
	err := r.db().NewSelect().
		Model(&items).Scan(context.TODO())
	return items, err
}

func (r *Repository) GetGroups(c context.Context, groupID string) (models.SearchGroups, error) {
	items := make(models.SearchGroups, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("SearchGroupMetaColumns").
		Order("search_group_order").Where("route is not null")

	if groupID != "" {
		query = query.Where("id = ?", groupID)
	}
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Search(c context.Context, searchGroup *models.SearchGroup, searchModel *models.SearchModel) error {
	querySelect := fmt.Sprintf("SELECT %s.%s as value, substring(%s for 40) as label", searchGroup.Table, searchGroup.ValueColumn, searchGroup.LabelColumn)
	queryFrom := fmt.Sprintf("FROM %s", searchGroup.Table)
	join := ""

	condition := fmt.Sprintf("where replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ' , '') ILIKE %s", searchGroup.SearchColumn, "'%"+searchModel.Query+"%'")
	conditionTranslitToRu := fmt.Sprintf("or replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ', '') ILIKE %s", searchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToRu(searchModel.Query)+"%'")
	conditionTranslitToEng := fmt.Sprintf("or replace(regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g'), ' ', '') ILIKE %s", searchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToEng(searchModel.Query)+"%'")

	queryOrder := fmt.Sprintf("ORDER BY %s", searchGroup.LabelColumn)
	query := fmt.Sprintf("%s %s %s %s %s %s %s", querySelect, queryFrom, join, condition, conditionTranslitToRu, conditionTranslitToEng, queryOrder)

	rows, err := r.helper.DB.IDB(c).QueryContext(c, query)
	if err != nil {
		return err
	}

	err = r.helper.DB.DB.ScanRows(c, rows, &searchGroup.SearchElements)
	return err
}
