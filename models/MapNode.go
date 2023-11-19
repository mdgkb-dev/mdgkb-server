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
	Neighbors     MapNodes
	NeighborsUUID []uuid.NullUUID `json:"neighbors"`
}

type MapNodes []*MapNode

func (items MapNodes) ToMap() map[*MapNode]int {
	m := make(map[*MapNode]int)
	for _, v := range items {
		m[v] = 1
	}
	return m
}

func (items MapNodes) InitNeighbors() {
	for i := range items {
		items[i].InitNeighbors()
	}
}

func (i *MapNode) InitNeighbors() {
	for _, v := range i.NeighborsUUID {
		i.Neighbors = append(i.Neighbors, &MapNode{ID: v})
	}
}
