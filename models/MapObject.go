package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapObject struct {
	bun.BaseModel `bun:"map_objects,alias:map_objects"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	MapNode       *MapNode      `bun:"rel:belongs-to" json:"mapNode"`
	MapNodeID     uuid.NullUUID `bun:"type:uuid"  json:"mapNodeId"`
	NodeName      string        `json:"nodeName"`
	Name          string        `json:"name"`
}

type MapObjects []*MapObject
