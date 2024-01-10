package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapNode struct {
	bun.BaseModel `bun:"map_nodes,alias:map_nodes"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	IsEntry       bool          `json:"isEntry"`
	Neighbors     MapNodes      `bun:"-"`
	// NeighborsUUID []uuid.NullUUID `json:"neighborNodesMapRoutes"`
	NeighborsNames []string `bun:"-" json:"neighborsNames"`
}

type MapNodes []*MapNode

func (items MapNodes) InitNeighbors() {
	for i := range items {
		items[i].InitNeighbors()
	}
}

func (i *MapNode) InitNeighbors() {
	// for _, v := range i.NeighborsUUID {
	// 	i.Neighbors = append(i.Neighbors, &MapNode{ID: v})
	// }
	for _, v := range i.NeighborsNames {
		i.Neighbors = append(i.Neighbors, &MapNode{Name: v})
	}
}
