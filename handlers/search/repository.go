package search

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

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
