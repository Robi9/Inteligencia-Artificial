package main

import (
	"fmt"
	"strconv"
)

// Graph : representa um grafo
type Graph struct {
	nodes []*GraphNode
}

// GraphNode : representa um nó do grafo
type GraphNode struct {
	id    string
	edges map[string]int
}

// New : retorna uma nova instancia de um grafo
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

// AddNode : adiciona um novo nó no grafo
func (g *Graph) AddNode(nomeNo string) (id string) {
	id = nomeNo
	g.nodes = append(g.nodes, &GraphNode{
		id:    id,
		edges: make(map[string]int),
	})
	return
}

// AddEdge : adiciona uma aresta no grafo junto com seu peso
func (g *Graph) AddEdge(n1, n2 string, w int) {
	for _, node := range g.nodes{
		if node.id == n1{
			node.edges[n2] = w
		}
	}
}

// Neighbors : retorna os vizinhos do nó de entrada
func (g *Graph) Neighbors(id string) []string {
	neighbors := []string{}
	for _, node := range g.nodes {
		for edge := range node.edges {
			if node.id == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}

// Nodes : retorna a lista dos nomes de cada nó do grafo
func (g *Graph) Nodes() []string {
	nodes := make([]string, len(g.nodes))
	for j, i := range g.nodes {
		nodes[j] = i.id
	}
	return nodes
}

// Edges : retorna a lista de arestas do grafo com seus pesos
func (g *Graph) Edges() [][3]string {
	edges := make([][3]string, 0, len(g.nodes))
	for j,i := range g.nodes {
		for k, v := range g.nodes[j].edges {
			s := strconv.Itoa(v)
			edges = append(edges, [3]string{i.id, k, s})
		}
	}
	return edges
}

func main() {

//GRAFO DA ROMENIA

	//ADIÇÃO DOS NÓS
	graph := New()
	node0 := graph.AddNode("ARAD")
	node1 := graph.AddNode("ZERIND")
	node2 := graph.AddNode("TIMISOARA")
	node3 := graph.AddNode("SIBIU")
	node4 := graph.AddNode("ORADEA")
	node5 := graph.AddNode("LUGOJ")
	node6 := graph.AddNode("FAGARAS")
	node7 := graph.AddNode("RIMNICU VILCEA")
	node8 := graph.AddNode("MEHADIA")
	node9 := graph.AddNode("BUCHAREST")
	node10 := graph.AddNode("PITEST")
	node11 := graph.AddNode("CRAIOVA")
	node12 := graph.AddNode("DROBETA")
	node13 := graph.AddNode("GIURGIU")
	node14 := graph.AddNode("URZICENI")
	node15 := graph.AddNode("VASLUI")
	node16 := graph.AddNode("HIRSOVA")
	node17 := graph.AddNode("IASI")
	node18 := graph.AddNode("EFORIE")
	node19 := graph.AddNode("NEAMT")
	
	//ADIÇÃO DAS ARESTAS
	graph.AddEdge(node0, node1, 75)
	graph.AddEdge(node0, node2, 118)
	graph.AddEdge(node0, node3, 140)
	graph.AddEdge(node1, node4, 71)
	graph.AddEdge(node2, node5, 111)
	graph.AddEdge(node3, node4, 151)
	graph.AddEdge(node3, node6, 99)
	graph.AddEdge(node3, node7, 80)
	graph.AddEdge(node5, node8, 70)
	graph.AddEdge(node6, node9, 211)
	graph.AddEdge(node7, node10, 97)
	graph.AddEdge(node7, node11, 146)
	graph.AddEdge(node8, node12, 75)
	graph.AddEdge(node9, node10, 101)
	graph.AddEdge(node9, node13, 90)
	graph.AddEdge(node9, node14, 85)
	graph.AddEdge(node10, node11, 138)
	graph.AddEdge(node11, node12, 120)
	graph.AddEdge(node14, node15, 142)
	graph.AddEdge(node14, node16, 98)
	graph.AddEdge(node15, node17, 92)
	graph.AddEdge(node16, node18, 86)
	graph.AddEdge(node17, node19, 87)
	
	//imprime os nós do grafo
	fmt.Println(graph.Nodes()) 

	//imprime as arestas e seus respectivos pesos
	fmt.Println(graph.Edges())

	//imprime os vizinhos do nó 10
	fmt.Println(graph.Neighbors(node1))
	
}