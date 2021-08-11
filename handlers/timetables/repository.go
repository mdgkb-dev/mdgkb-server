package timetables

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	getAllWeekdays(*gin.Context) ([]models.Weekday, error)
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) getAllWeekdays(ctx *gin.Context) (items []models.Weekday, err error) {
	err = r.db.NewSelect().Model(&items).Order("number").Scan(ctx)
	return items, err
}
