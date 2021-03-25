package main

import (
	"fmt"
	"os"
	"strconv"
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
//Representa a Borda
type Borda struct {
	items []*GraphNode
	custos map[*GraphNode]int
	top   int
}

//Representa o Explorados
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

// Init: Inicializa a Borda
func Init(size int) *Borda {
	s := &Borda{
		items: make([]*GraphNode, size),
		custos: make(map[*GraphNode]int),
		top: -1,
	}
	return s
}

// IsInitialized: verifica se a Borda está inicializada
func (s *Borda) IsInitialized() bool {
	if cap(s.items) == 0 {
		return true
	}
	return false
}

// IsFull: verifica se a borda está cheia
func (s *Borda) IsFull() bool {
	if (cap(s.items) - 1) == s.top {
		return true
	}
	return false
}

// IsEmpty: Verifica se a Borda está vazia
func (s *Borda) IsEmpty() bool {
	if -1 == s.top {
		return true
	}
	return false
}
//retornaPai: Retorna nó que antecedeu determinado nó na busca
//OBS: Usamos o nome pai para facilitar o uso, mas não utilizamos pai na estrutura do grafo.
func retornaPai(pai map[*GraphNode][]*GraphNode, no *GraphNode) *GraphNode {
	for node,i := range pai {
		for _,j := range i{
			if j.id == no.id {
				return node
			}
		}	
	}
	return nil
}

// Enfileira: Enfileira um nó na Borda com seu custo
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

// Print: printa elementos da Borda
func (s *Borda) Print() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element, "Custo=",s.custos[element])
	}
}
//PrintExp: Printa os Explorados
func (s *Explorados) PrintExp() {
	for i, element := range s.items {
		fmt.Println("Number=", i, "Element=", element)
	}
}

// Desenfileira: Desenfileira um elemento da Borda de acordo com seu custo
func (s *Borda) Desenfileira() (int, *GraphNode){
	c := 0
	n := &GraphNode{
		id:    "",
		edges: make(map[string]int),
	}
	for i, element := range s.items {
		if element == nil {
			break
		}else if i == 0{
			c = s.custos[element]
			n = element
		}else if s.custos[element] <= c && i != 0{
			c = s.custos[element]
			n = element
		}	
	}
	
	for i, element := range s.items {
		if element.id == n.id {
			aux := s.items[s.top]
			s.items[s.top] = nil
			s.items[i]=aux
			delete(s.custos, element)
			s.top--
			
			return c, element
		}
	}
	return 0, nil
}

//searchBorda: Retorna se determinado nó está na Borda com custo maior
func (s *Borda) searchBorda(node *GraphNode, custo int) bool{
	for _,i := range s.items {
		if i == node && s.custos[i] > custo{
			return true
		}
	}
	return false
}

//searchExplorados: Verifica se um nó está no Explorados
func (s *Explorados) searchExplorados(node *GraphNode) bool{
	for _,i := range s.items {
		if i == node {
			return true
		}
	}
	return false
}

//searchNode: Procura um nó no Grafo e o retorna
func (g *Graph) searchNode (node string) *GraphNode{
	for _,i := range g.nodes {
		if i.id == node {
			return i
		}
	}

	return nil
}

//searchB: Verifica se um nó está na Borda
func (s *Borda) searchB(node *GraphNode) bool{
	for _,i := range s.items {
		if i == node {
			return true
		}
	}
	return false
}

//HeuristicasTable: Retorna um map com os valores de heuristicas
func HeuristicasTable() (map[string]int) {
	heuristicas := make(map[string]int)

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

	return heuristicas
}

//Realiza a Busca com Informação A*
func A_Star(g *Graph, inicio *GraphNode,  final string) int {
	heuristica:=make(map[string]int)
	heuristica = HeuristicasTable()
	custoNo:=0
	custoCaminho:=make(map[string]int)
	soma := make(map[string]int)
	pai := make(map[*GraphNode][]*GraphNode)
	e := &Explorados{
		items: []*GraphNode{},
	}
	
	b := Init(len(g.nodes))
	b.Enfileira(inicio, 0+heuristica[inicio.id])	
	for {
		if b.IsEmpty() {
			fmt.Println("Erro, borda está vazia.")
			return 0
		}
		custo, node := b.Desenfileira()
		soma[node.id] = custo
		if node.id == final {
			fmt.Println("Destino encontrado")
			fmt.Println("Custo - ", soma[node.id])
			return 0		
		}
		e.items = append(e.items, node)

		if len(e.items)-1 != 0{
			custoCaminho[node.id] = node.edges[retornaPai(pai, node).id]+custoCaminho[retornaPai(pai, node).id]
		}else{
			custoCaminho[node.id]=0
		}

		for _,filho := range g.Neighbors(node){
			pai[node] = append(pai[node], filho)
			if !(e.searchExplorados(filho)) && !(b.searchB(filho)){
				custoNo=custoCaminho[node.id]+filho.edges[node.id]+heuristica[filho.id]
				b.Enfileira(filho, custoNo)
			}else if b.searchBorda(filho, custoCaminho[node.id]+filho.edges[node.id]+heuristica[filho.id]){
				custoNo = custoCaminho[node.id]+filho.edges[node.id]+heuristica[filho.id]
				b.custos[filho]=custoNo
						
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

	//Procura pelo nó de partida no Grafo
	node:=graph.searchNode(Partida)
	if node != nil {
		//Busca A*
		A_Star(graph, node, Final)
	}	
}