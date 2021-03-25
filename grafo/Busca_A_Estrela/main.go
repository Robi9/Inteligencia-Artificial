package main

import (
	"fmt"
	"strconv"
	"os"
)

// Graph : representa um grafo
type Graph struct {
	nodes []*GraphNode
}

// GraphNode : representa um nó do grafo
type GraphNode struct {
	id    string
	edges map[string]int //*GraphNode
}

// New : retorna uma nova instancia de um grafo
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

type Borda struct {
	items []*GraphNode
	custos map[*GraphNode]int
	top   int
}


type Explorados struct{
	items []*GraphNode
}

// AddNode : adiciona um novo nó no grafo
func (g *Graph) AddNode(nomeNo string) (node *GraphNode) {
	
	node = &GraphNode{
		id:    nomeNo,
		edges: make(map[string]int),
	}
	g.nodes = append(g.nodes, node)
	return
}

// AddEdge : adiciona uma aresta no grafo junto com seu peso
func (g *Graph) AddEdge(n1, n2 *GraphNode, w int) {
	for _, node := range g.nodes{
		if node.id == n1.id{
			node.edges[n2.id] = w
		}
	}

	for _, node := range g.nodes{
		if node.id == n2.id{
			node.edges[n1.id] = w
		}
	}
}

// Neighbors : retorna os vizinhos do nó de entrada
func (g *Graph) Neighbors(id *GraphNode) []*GraphNode {
	neighbors := []string{}
	nodes := []*GraphNode{}

	for _, node := range g.nodes {
		if node.id == id.id {
			for edge := range node.edges {
				//if edge == id {
					neighbors = append(neighbors, edge)
				//}
			}
			break
		}
	}

	for _, i := range neighbors {
		for _, j := range g.nodes {
			if i == j.id {
				nodes = append(nodes, j)
			}
		}
	}
	return nodes
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

// Init - Borda initialization
func Init(size int) *Borda {
	s := &Borda{
		items: make([]*GraphNode, size),
		custos: make(map[*GraphNode]int),
		top: -1,
	}
	return s
}

// IsInitialized - checks Borda initialized or not
func (s *Borda) IsInitialized() bool {
	if cap(s.items) == 0 {
		return true
	}
	return false
}

// IsFull - checks if Borda is full
func (s *Borda) IsFull() bool {
	if (cap(s.items) - 1) == s.top {
		return true
	}
	return false
}

// IsEmpty - checks if Borda is empty
func (s *Borda) IsEmpty() bool {
	if -1 == s.top {
		return true
	}
	return false
}

func retornaPai(pai map[string][]*GraphNode, no string) string {
	for nome,i := range pai {
		for _,j := range i{
			if j.id == no {
				return nome
			}
		}	
	}
	return ""
}

// Push - pushes element into Borda
func (s *Borda) Enfileira(element *GraphNode, custo int) {
	s.top++
	if s.top == -1 {
		s.items[0] = element
		s.custos[element] = custo
	} else {
		s.items[s.top] = element
		s.custos[element] = custo
	}
}

// Print - prints element from Borda
func (s *Borda) Print() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

func (s *Explorados) PrintExp() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

// Pop - pop element from Borda
func (s *Borda) Desenfileira() (int, *GraphNode){
	c := 0
	n := &GraphNode{
		id:    "",
		edges: make(map[string]int),
	}
	//var n *GraphNode
	for _, element := range s.items {
		if s.custos[element] == 0 {
			
		}
		if s.custos[element] <= c {
			c = s.custos[element]
			n = element
		}
	}

	
	for _, element := range s.items {
		fmt.Println(element.id)
		if element.id == n.id {
			s.items[s.top] = nil
			delete(s.custos, element)
			s.top--
			//s.Print()
			return c, element
		}
	}
	return 0, nil
}

// Peek - gives top element
func (s *Borda) Peek() int {
	return s.top
}

func (s *Borda) searchBorda(node *GraphNode, custo int) bool{
	for _,i := range s.items {
		if i == node && s.custos[node] < custo{
			return true
		}
	}
	return false
}

func (s *Explorados) searchExplorados(node *GraphNode) bool{
	for _,i := range s.items {
		if i == node {
			return true
		}
	}
	return false
}

func (g *Graph) searchNode (node string) *GraphNode{
	for _,i := range g.nodes {
		if i.id == node {
			return i
		}
	}

	return nil
}
func (s *Borda) Substitui(node *GraphNode, custo int) {
	for j,i := range s.items {
		if i == node {
			s.items[j]=node
			s.custos[node]=custo
		}
	}
}


func HeusticasTable() (heuristicas map[string]int) {
	//var heuristicas map[string]int

	heuristicas["ARAD"]= 366
	heuristicas["ZERIND"]= 374
	heuristicas["TIMISOARA"]= 329
	heuristicas["SIBIU"]= 253
	heuristicas["ORADEA"]= 380
	heuristicas["LUGOJ"]= 244
	heuristicas["FAGARAS"]= 176
	heuristicas["RIMNICU VILCEA"]= 193
	heuristicas["MEHADIA"]= 241
	heuristicas["BUCHAREST"]= 0
	heuristicas["PITEST"]= 100
	heuristicas["CRAIOVA"]= 160
	heuristicas["DROBETA"]= 242
	heuristicas["GIURGIU"]= 77
	heuristicas["URZICENI"]= 80
	heuristicas["VASLUI"]= 199
	heuristicas["HIRSOVA"]= 151
	heuristicas["IASI"]= 226
	heuristicas["EFORIE"]= 161
	heuristicas["NEAMT"]= 234

	return
}

func A_Star(g *Graph, inicio *GraphNode,  final string) int {
	soma := make(map[string]int)
	pai := make(map[string][]*GraphNode)
	e := &Explorados{
		items: []*GraphNode{},
	}
	b := Init(len(g.nodes))
	b.Enfileira(inicio, 0)
	
	for {
		if b.IsEmpty() {
			fmt.Println("Erro, borda está vazia.")
			return 0
		}
		//b.Print()
		custo, node := b.Desenfileira()
		fmt.Println(node.id)
		soma[node.id] = custo
		if node.id == final {
			fmt.Println("Destino encontrado")
			fmt.Println("Custo - ", soma[node.id]) 
			return 0		
		}

		e.items = append(e.items, node)

		/*if len(e.items)-1 != 0 {
			soma[node.id] = custo //node.edges[retornaPai(pai, node.id)]+soma[retornaPai(pai, node.id)]
		}else{
			soma[node.id]=custo
		}*/

		for _,filho := range g.Neighbors(node){
			pai[node.id] = append(pai[node.id], filho)
			if !(b.searchBorda(filho, soma[retornaPai(pai, filho.id)]+filho.edges[retornaPai(pai, node.id)])) && !(e.searchExplorados(filho)){
				b.Enfileira(filho, soma[retornaPai(pai, filho.id)]+filho.edges[retornaPai(pai, node.id)])
			}else if b.searchBorda(filho, soma[retornaPai(pai, filho.id)]+filho.edges[retornaPai(pai, node.id)]){
				b.Substitui(filho, soma[retornaPai(pai, filho.id)]+filho.edges[retornaPai(pai, node.id)])
			}
		}
	}
}

func main() {

	Partida := os.Args[1] 
	Final := os.Args[2]
	
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

	node:=graph.searchNode(Partida)
	if node != nil {
		//Busca em Profundidade
		A_Star(graph, node, Final)
	}
	
	
}

/*
*/