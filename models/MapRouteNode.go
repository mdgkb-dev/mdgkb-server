package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapRouteNode struct {
	bun.BaseModel `bun:"map_route_nodes,alias:map_route_nodes"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	MapRoute   *MapRoute     `bun:"rel:belongs-to" json:"mapRoute"`
	MapRouteID uuid.NullUUID `bun:"type:uuid"  json:"mapRouteId"`

	MapNode   *MapNode      `bun:"rel:belongs-to" json:"mapNode"`
	MapNodeID uuid.NullUUID `bun:"type:uuid"  json:"mapNodeId"`
}

type MapRouteNodes []*MapRouteNode