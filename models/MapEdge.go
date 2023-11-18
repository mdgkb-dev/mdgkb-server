package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapEdge struct {
	bun.BaseModel `bun:"map_edges,alias:map_edges"`

	PreviousNode   *MapNode      `bun:"rel:belongs-to" json:"previousNode"`
	PreviousNodeID uuid.NullUUID `bun:"type:uuid"  json:"previousNodeId"`

	NextNode   *MapNode      `bun:"rel:belongs-to" json:"nextNode"`
	NextNodeID uuid.NullUUID `bun:"type:uuid"  json:"nextNodeId"`
}

type MapEdges []*MapEdge
