package main

import (
	"log"

	"github.com/watson370/adventofcodego/graph"
)

func main() {
	mygraph := graph.Create(10)
	mygraph.AddEdge(1, 2)
	adj := mygraph.Adjacent(1)
	log.Println(adj)
	mygraph.AddEdge(1, 5)
	log.Println(mygraph.Adjacent(1))

}
