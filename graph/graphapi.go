package graph

//Graph  an interface with the methods all graphs should impl
//uses ints for the vertices, clients will use a symbol table to map their node value to the int one
type Graph interface {
	//AddEdge adds an edge between v and w
	AddEdge(v int, w int)
	//Adjacent returns a slice of ints representing nodes adjacent to v
	Adjacent(v int) []int
	//Vertices returns the total number of vertices in this graph
	Vertices() int
	//Nodes returns the total number of nodes in this graph
	Nodes() int
}

//Create returns a new graph with
func Create(size int) Impl {
	nimpl := Impl{
		V:   size,
		E:   0,
		via: make([][]int, size),
	}
	return nimpl
}

//Impl an impl of the Graph interface
type Impl struct {
	V int //number of vertexs
	E int // number of edges
	//use a vertex indexed array
	via [][]int //a slice where each index containse a slice with the index of adjacent vertices

}

//AddEdge adds an edge betwee node v and node w
func (g *Impl) AddEdge(v int, w int) {
	if v < 0 || w < 0 || v >= g.V || w >= g.V {
		panic("I didn't like that")
	}
	g.E = g.E + 1
	g.via[v] = append(g.via[v], w)
	g.via[w] = append(g.via[w], v)
}

//Adjacent returns a slice of nodes adjacent to v
func (g *Impl) Adjacent(v int) []int {
	return g.via[v]
}

//Vertices returns the total number of vertices in this graph
func (g *Impl) Vertices() int {
	return len(g.via)
}

//Nodes returns the total number of nodes in this graph
func (g *Impl) Nodes() int {
	return 2
}
