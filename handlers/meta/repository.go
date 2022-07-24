package meta

import (
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
	queryContext, err := r.db().Query(query, optionModel.Value, optionModel.Label, optionModel.TableName, optionModel.SortColumn)
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
