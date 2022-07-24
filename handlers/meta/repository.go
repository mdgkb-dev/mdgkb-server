package meta

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getCount(table *string) (res *int, err error) {
	num := 0
	query := fmt.Sprintf("SELECT COUNT (id) FROM %s", *table)
	err = r.db().QueryRow(query).Scan(&num)
	return &num, err
}

func (r *Repository) getOptions(optionModel *models.OptionModel) (models.Options, error) {
	options := make(models.Options, 0)
	query := fmt.Sprintf("SELECT %s::varchar as value, %s as label FROM %s ORDER BY %s", optionModel.Value, optionModel.Label, optionModel.TableName, optionModel.SortColumn)
	queryContext, err := r.db().QueryContext(r.ctx, query)
	if err != nil {
		return nil, err
	}
	err = r.db().ScanRows(r.ctx, queryContext, &options)
	return options, err
}

func (r *Repository) getApplicationsCounts() (models.ApplicationsCounts, error) {
	items := make(models.ApplicationsCounts, 0)
	err := r.db().NewSelect().
		Model(&items).Scan(r.ctx)
	return items, err
}
