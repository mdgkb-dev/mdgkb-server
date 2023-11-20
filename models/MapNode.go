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
}

type MapNodes []*MapNode

type Graph struct {
	nodes map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]map[int]int)}
}

func (g *Graph) AddNode(name int, neighbors map[int]int) {
	g.nodes[name] = neighbors
}

// func (item *MapRouteNode) Dijkstra(graph *Graph, start, end int) ([]int, int) {
// 	distances := make(map[int]int)
// 	visited := make(map[int]bool)
// 	previous := make(map[int]int)
// 	var currentNode int

// 	// Initialize distances
// 	for node := range graph.nodes {
// 		distances[node] = math.MaxInt32
// 	}
// 	distances[start] = 0

// 	for len(visited) < len(graph.nodes) {
// 		currentNode = minDistance(distances, visited)
// 		visited[currentNode] = true

// 		for neighbor, weight := range graph.nodes[currentNode] {
// 			if distances[currentNode]+weight < distances[neighbor] {
// 				distances[neighbor] = distances[currentNode] + weight
// 				previous[neighbor] = currentNode
// 			}
// 		}
// 	}

// 	// Reconstruct the path
// 	path := []int{end}
// 	for previousNode := previous[end]; previousNode != 0; previousNode = previous[previousNode] {
// 		path = append([]int{previousNode}, path...)
// 	}

// 	return path, distances[end]
// }

// func (item *MapRouteNode) minDistance(distances map[int]int, visited map[int]bool) int {
// 	min := math.MaxInt32
// 	var minNode int

// 	for node, distance := range distances {
// 		if distance < min && !visited[node] {
// 			min = distance
// 			minNode = node
// 		}
// 	}

// 	return minNode
// }

// func (item *MapRouteNode) MakeGraph(graph *Graph) {

// 	db := InitDB()

// 	var (
// 		id       int
// 		neighbor int
// 	)

// 	count := 0
// 	rows, err := db.Query("select count(*) from nodes", 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&count)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	rows.Close()

// 	for i := 1; i <= count; i++ {
// 		rows, err := db.Query("select n.id,e.next_node from nodes n inner join edges e on n.id = e.previous_node where n.id="+strconv.Itoa(i)+" order by n.id", 1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()

// 		var t map[int]int = make(map[int]int)
// 		for rows.Next() {
// 			err := rows.Scan(&id, &neighbor)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			t[neighbor] = 1
// 		}

// 		graph.AddNode(i, t)
// 	}
// }
