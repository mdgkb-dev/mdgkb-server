package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MapRoute struct {
	bun.BaseModel `bun:"map_routes,alias:map_routes"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	StartNode     *MapNode      `bun:"rel:belongs-to" json:"startNode"`
	StartNodeID   uuid.NullUUID `bun:"type:uuid"  json:"startNodeId"`
	StartNodeName string        `json:"startNodeName"`

	EndNode       *MapNode      `bun:"rel:belongs-to" json:"endNode"`
	EndNodeID     uuid.NullUUID `bun:"type:uuid"  json:"endNodeId"`
	EndNodeName   string        `json:"endNodeName"`
	MapRouteNodes MapRouteNodes `bun:"rel:has-many" json:"mapRouteNodes"`
}

type MapRoutes []*MapRoute

func (items MapRoutes) GetMapRouteNodes() MapRouteNodes {
	itemsForGet := make(MapRouteNodes, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].MapRouteNodes...)
	}
	return itemsForGet
}

func (items MapRoutes) SetIdForChildren() {
	for i := range items {
		for j := range items[i].MapRouteNodes {
			items[i].MapRouteNodes[j].MapRouteID = items[i].ID
		}
	}
}

func (i *MapRoute) Calculate(nodes MapNodes, g *Graph) {
	routeNodes, _ := BFSWithPath(g, i.StartNode, i.EndNode)

	for index, v := range routeNodes {
		mrn := MapRouteNode{MapRouteID: i.ID, MapNodeID: v.ID, MapNodeName: v.Name, Order: uint(index)}
		i.MapRouteNodes = append(i.MapRouteNodes, &mrn)
	}
}

func (items MapRoutes) Calculate(nodes MapNodes) MapRoutes {
	routes := make(MapRoutes, 0)
	for _, node := range nodes {
		if !node.IsEntry {
			continue
		}

		for _, pair := range nodes {
			if !pair.IsEntry {
				continue
			}

			if node != pair {
				routes = append(routes, &MapRoute{StartNode: node, EndNode: pair, StartNodeName: node.Name, EndNodeName: pair.Name})
				// items = append(items, &MapRoute{StartNode: node, EndNode: pair})
			}
		}
	}

	g := NewGraph()
	g.Init(nodes)
	fmt.Println("routes", len(items), len(g.nodes))

	for i := range routes {
		routes[i].Calculate(nodes, g)
		// if len(items[i].MapRouteNodes) > 0 {
		// 	fmt.Println("i:", items[i].MapRouteNodes)
		// }
	}
	return routes
	// fmt.Println("return")
}

type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

func BFSWithPath(graph *Graph, start, end *MapNode) (MapNodes, int) {
	visited := make(map[string]bool)
	previous := make(map[string]string)
	var currentNode string

	// fmt.Println("start", &start.Name)
	// fmt.Println("end", &end.Name)
	queue := []string{start.Name}
	visited[start.Name] = true

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]

		if currentNode == end.Name {
			// Построение пути от конечного узла к начальному
			path := MapNodes{end}
			previousNode := end.Name

			for {
				previousNode = previous[previousNode]
				path = append(MapNodes{&MapNode{Name: previousNode}}, path...)
				if previousNode == start.Name {
					break
				}
			}
			// for i, node := range path {
			// fmt.Printf("path[%d]", i)
			// fmt.Println("=", node.Name)
			// }
			// fmt.Println("")
			return path, len(path) - 1 // Длина пути минус 1, чтобы получить количество ребер
		}

		neighbors, found := graph.nodes[currentNode]
		if !found {
			continue
		}

		for _, neighbor := range neighbors {
			// fmt.Println("currentNode", currentNode)
			// fmt.Println("neighbor", neighbor)
			// fmt.Println("visited[neighbor]", visited[neighbor])
			if !visited[neighbor] {

				queue = append(queue, neighbor)
				visited[neighbor] = true
				previous[neighbor] = currentNode

			}
		}
	}

	// Если путь не найден
	// fmt.Println("путь не найден")
	return nil, -1
}

func (i *Graph) Init(nodes MapNodes) {
	i.nodes = make(map[string][]string)
	for _, v := range nodes {
		// i.nodes[v.ID] = v.NeighborsUUID
		i.nodes[v.Name] = v.NeighborsNames
	}
}
