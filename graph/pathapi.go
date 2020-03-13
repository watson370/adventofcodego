package graph

import "log"

//Path an interface with the methods for a Path
type Path interface {
	HasPathTo(n int) bool
	PathTo(n int) []int
}

//CreatePath takes in a graph and a node to start from, and builds an efficient way to query if there is a path to another node
func CreatePath(g Graph, node int) PathImpl {
	marked := make([]bool, g.Vertices())
	pathTo := make([]int, g.Vertices())
	for i := 0; i < len(pathTo); i++ {
		pathTo[i] = -1
	}
	todo := make([]int, 0, g.Vertices())
	adjacent := make([][]int, 0, g.Vertices()) //is it better to add capacity and append?
	//populate adjacent
	for i := 0; i < g.Vertices(); i++ {
		adjacent = append(adjacent, g.Adjacent(i))
	}
	log.Println(adjacent)
	//DFS in a while loop
	todo = append(todo, node)
	pathTo[node] = node
	for len(todo) > 0 {
		current := todo[len(todo)-1] //using slice as a stack just keep taking the last one
		marked[current] = true

		if len(adjacent[current]) > 0 {
			next := adjacent[current][len(adjacent[current])-1]
			adjacent[current] = adjacent[current][:len(adjacent[current])-1] //pop, ideally pull one at random instead of from end every time
			for marked[next] && len(adjacent[current]) > 0 {
				adjacent[current] = adjacent[current][:len(adjacent[current])-1] //pop, ideally pull one at random instead of from end every time
				next = adjacent[current][len(adjacent[current])-1]
			}
			if len(adjacent[current]) == 0 && marked[next] {
				log.Printf("finished processing %d, it has no more adjacents", todo[len(todo)-1])
				todo = todo[:len(todo)-1] //pop
			} else {
				todo = append(todo, next)
				log.Printf("processing next %d and removing from adjacent slice, adj is now %v todo is %v\n", next, adjacent[current], todo)
				//first encounter
				log.Printf("setting path to %d as %d \n", next, current)
				pathTo[next] = current

			}
		} else {
			log.Printf("finished processing %d, it has no more adjacents", todo[len(todo)-1])
			todo = todo[:len(todo)-1] //pop
		}

	}
	return PathImpl{
		fromNode: node,
		pathTo:   pathTo,
	}
	// return nil
}

//CreatePath2 takes in a graph and a node to start from, and builds an efficient way to query if there is a path to another node
func CreatePath2(g Graph, node int) PathImpl {
	marked := make([]bool, g.Vertices())
	pathTo := make([]int, g.Vertices())
	for i := 0; i < len(pathTo); i++ {
		pathTo[i] = -1
	}
	todo := make([]fromto, 0, g.Edges()*2)
	adjacent := make([][]int, 0, g.Vertices()) //is it better to add capacity and append?
	//populate adjacent
	for i := 0; i < g.Vertices(); i++ {
		adjacent = append(adjacent, g.Adjacent(i))
	}
	log.Println(adjacent)
	//DFS in a while loop
	todo = append(todo, fromto{
		from: node,
		to:   node,
	})

	for len(todo) > 0 {
		current := todo[len(todo)-1] //using slice as a stack just keep taking the last one
		if marked[current.to] {
			todo = todo[:len(todo)-1]
			continue
		}
		marked[current.to] = true
		pathTo[current.to] = current.from
		for _, adjto := range adjacent[current.to] {
			todo = append(todo, fromto{
				from: current.to,
				to:   adjto}) //need a way to append the adjacent node value along with the from value

		}
	}
	return PathImpl{
		fromNode: node,
		pathTo:   pathTo,
	}
	// return nil
}

type fromto struct {
	from int
	to   int
}

//CreatePathDFS takes in a graph and a node to start from, and builds an efficient way to query if there is a path to another node
func CreatePathDFS(g Graph, node int) BFSPathImpl {
	marked := make([]bool, g.Vertices())
	pathTo := make([]int, g.Vertices())
	distTo := make([]int, g.Vertices())
	for i := 0; i < len(pathTo); i++ {
		pathTo[i] = -1
	}
	//change todo into a queue
	todo := RingBuffer{
		buffer: make([]int, g.Edges()),
		out:    0,
		in:     0,
	}
	//BFS in a while loop
	todo.Add(node)
	marked[node] = true
	pathTo[node] = node
	//BFS
	//take one off the queue
	//check if marked, if so ignore
	//if not, mark it, set path to, set dist to, add it to the queue

	for !todo.Empty() {
		current, _ := todo.Pop()
		for _, a := range g.Adjacent(current) {
			if !marked[a] {
				marked[a] = true
				distTo[a] = distTo[current] + 1
				pathTo[a] = current
				todo.Add(a)
			}
		}
	}
	return BFSPathImpl{
		fromNode: node,
		pathTo:   pathTo,
		DistTo:   distTo,
	}
	// return nil
}
