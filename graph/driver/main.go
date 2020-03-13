package main

import (
	"log"

	"github.com/watson370/adventofcodego/graph"
)

func main() {
	mygraph := graph.Create(10)
	mygraph.AddEdge(0, 1)
	mygraph.AddEdge(1, 2)
	mygraph.AddEdge(1, 7)

	// adj := mygraph.Adjacent(1)
	// log.Println(adj)
	mygraph.AddEdge(2, 5)
	// log.Println(mygraph.Adjacent(1))
	p := graph.CreatePath(&mygraph, 0)
	log.Println(mygraph)
	log.Println(p)
	log.Printf("the path to 5 is %v \n", p.PathTo(5))

	p2 := graph.CreatePath2(&mygraph, 0)
	log.Println(p2)
	log.Printf("the path to 5 is %v \n", p2.PathTo(5))

	//DFS
	p3 := graph.CreatePathDFS(&mygraph, 0)
	log.Println(p3)
	log.Printf("the path to 5 is %v \n", p3.PathTo(5))

}
