package graph

//PathImpl an implementation of the PathAPI
type PathImpl struct {
	fromNode int
	pathTo   []int
}

//HasPathTo true if it does have a path from the origin node to n passed in
func (p PathImpl) HasPathTo(n int) bool {
	if p.pathTo[n] != -1 {
		return true
	}
	return false
}

//PathTo an int slice with n as the first index, and the subsequent indexes are the path to the origin of the Path being created
func (p PathImpl) PathTo(n int) []int {
	toreturn := make([]int, 0, len(p.pathTo))
	current := n
	for {
		toreturn = append(toreturn, current)
		if p.pathTo[current] == current {
			break
		}
		current = p.pathTo[current]
	}
	return toreturn
}

//BFSPathImpl an implementation of the PathAPI
type BFSPathImpl struct {
	fromNode int
	pathTo   []int
	DistTo   []int
}

//HasPathTo true if it does have a path from the origin node to n passed in
func (p BFSPathImpl) HasPathTo(n int) bool {
	if p.pathTo[n] != -1 {
		return true
	}
	return false
}

//PathTo an int slice with n as the first index, and the subsequent indexes are the path to the origin of the Path being created
func (p BFSPathImpl) PathTo(n int) []int {
	toreturn := make([]int, 0, len(p.pathTo))
	current := n
	for {
		toreturn = append(toreturn, current)
		if p.pathTo[current] == current {
			break
		}
		current = p.pathTo[current]
	}
	return toreturn
}

//DistanceTo returns the distance from the root of this path to the node n
func (p BFSPathImpl) DistanceTo(n int) int {
	return p.DistTo[n]
}
