package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapRoute struct {
	bun.BaseModel `bun:"map_routes,alias:map_routes"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	StartNode   *MapNode      `bun:"rel:belongs-to" json:"startNode"`
	StartNodeID uuid.NullUUID `bun:"type:uuid"  json:"startNodeId"`

	EndNode   *MapNode      `bun:"rel:belongs-to" json:"endNode"`
	EndNodeID uuid.NullUUID `bun:"type:uuid"  json:"endNodeId"`
}

type MapRoutes []*MapRoute
