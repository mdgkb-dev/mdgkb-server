package models

import (
	"fmt"
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
		q = q.Join(fmt.Sprintf("JOIN %s ON %s.document_type_id = document_types.id", i.ItemsFor, i.ItemsFor))
	}
}
