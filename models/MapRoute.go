package models

import (
	"math"

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

	MapRouteNodes MapRouteNodes `bun:"rel:has-many" json:"mapRouteNodes"`
}

type MapRoutes []*MapRoute

func (i *MapRoute) Calculate(nodes MapNodes, g *Graph) {
	routeNodes, _ := Dijkstra(g, i.StartNode, i.EndNode)

	for _, v := range routeNodes {
		mrn := MapRouteNode{MapRouteID: i.ID, MapNodeID: v.ID}
		i.MapRouteNodes = append(i.MapRouteNodes, &mrn)
	}
}

func (items MapRoutes) Calculate(nodes MapNodes) {
	for _, node := range nodes {
		if !node.IsEntry {
			continue
		}
		for _, pair := range nodes {
			if !pair.IsEntry {
				continue
			}
			items = append(items, &MapRoute{StartNode: node, EndNode: pair})
		}
	}
	g := Graph{}
	g.Init(nodes)

	for i := range items {
		items[i].Calculate(nodes, &g)
	}
}

type Graph struct {
	nodes map[*MapNode]map[*MapNode]int
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[*MapNode]map[*MapNode]int)}
}

func Dijkstra(graph *Graph, start, end *MapNode) (MapNodes, int) {
	distances := make(map[*MapNode]int)
	visited := make(map[*MapNode]bool)
	previous := make(map[*MapNode]*MapNode)
	var currentNode *MapNode

	// Initialize distances
	for node := range graph.nodes {
		distances[node] = int(math.MaxInt32)
	}
	distances[start] = 0

	for len(visited) < len(graph.nodes) {
		currentNode = minDistance(distances, visited)
		visited[currentNode] = true

		for neighbor, weight := range graph.nodes[currentNode] {
			if distances[currentNode]+weight < distances[neighbor] {
				distances[neighbor] = distances[currentNode] + weight
				previous[neighbor] = currentNode
			}
		}
	}

	// Reconstruct the path
	path := MapNodes{end}
	for previousNode := previous[end]; previousNode != nil; previousNode = previous[previousNode] {
		path = append(MapNodes{previousNode}, path...)
	}

	return path, distances[end]
}

func minDistance(distances map[*MapNode]int, visited map[*MapNode]bool) *MapNode {
	min := int(math.MaxInt32)
	var minNode *MapNode

	for node, distance := range distances {
		if distance < min && !visited[node] {
			min = distance
			minNode = node
		}
	}

	return minNode
}

func (i *Graph) Init(nodes MapNodes) {

	for _, v := range nodes {
		i.nodes[v] = v.Neighbors.ToMap()
	}

}
