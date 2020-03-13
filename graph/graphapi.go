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
	Edges() int
}
