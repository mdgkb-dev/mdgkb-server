package models

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type DocumentsParams struct {
	ItemsFor string `json:"itemsFor"`
}

func CreateDocumentsParamsFromContext(c *gin.Context) DocumentsParams {
	return DocumentsParams{ItemsFor: c.Query("items-for")}
}

func (i *DocumentsParams) CreateJoin(q *bun.SelectQuery) {
	if i.ItemsFor != "" {
		q.Join("JOIN ? ON ?.document_type_id = document_types.id", bun.Ident(i.ItemsFor), bun.Ident(i.ItemsFor))
	}
}
