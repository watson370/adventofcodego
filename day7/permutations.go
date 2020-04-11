package day7

func CalcPerms(vals []int) [][]int {

	//final list of permutations
	permcollector := make([][]int, 0)
	//list of vals left and their index in the perm that they can go in
	todo := make([]ValsLeftIndex, 0)
	//need to prime with each
	for i := range vals {
		valsleft := make([]int, len(vals)-1)
		valsleftwriteindex := 0
		for ii := range vals {
			if ii != i {
				valsleft[valsleftwriteindex] = vals[ii]
				valsleftwriteindex = valsleftwriteindex + 1
			}
		}
		todo = append(todo, ValsLeftIndex{
			MyVal:    vals[i],
			ValsLeft: valsleft,
			Index:    0,
		})
	}
	currentperm := make([]int, len(vals))
	for len(todo) > 0 {
		//grab a todo
		currenttodo := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		currentperm[currenttodo.Index] = currenttodo.MyVal
		//check if your at the end of current perm
		if len(currentperm)-1 == currenttodo.Index {
			//add this perm to the perm collector
			permcopy := make([]int, len(currentperm))
			copy(permcopy, currentperm)
			permcollector = append(permcollector, permcopy)
		} else {
			//add todo for each
			for nexttodoindex := range currenttodo.ValsLeft {
				valsleft := make([]int, len(currenttodo.ValsLeft)-1)
				valsleftwriteindex := 0
				for ii := range currenttodo.ValsLeft {
					if ii != nexttodoindex {
						valsleft[valsleftwriteindex] = currenttodo.ValsLeft[ii]
						valsleftwriteindex = valsleftwriteindex + 1
					}
				}
				todo = append(todo, ValsLeftIndex{
					MyVal:    currenttodo.ValsLeft[nexttodoindex],
					ValsLeft: valsleft,
					Index:    currenttodo.Index + 1,
				})
			}
		}

	}
	return permcollector

}

type ValsLeftIndex struct {
	MyVal    int
	ValsLeft []int
	Index    int
}
