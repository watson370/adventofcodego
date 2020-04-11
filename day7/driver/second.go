// package main

// import "log"

// func main() {
// 	// make a list of vals for the first position needed
// 	// put the first one on the todo list
// 	// make a list fo vals for the second position given the current one on the todo list, put them on the needed list
// 	// take off one from needed and put on todo
// 	// givent the secon postiion one on the todo list, make a list of needed vals in the third and put on needed list
// 	// when you take an item off of the needed list and put it on the todo list, remove it from the possible values in the list in the needed list
// 	// if you hit an item in todo that is in the last position, add it to permcollector and remove it from todo
// 	// if you have an item in todo that has no more vals left, remove it from todo

// 	stillNeeded := make([]Todo, 0)
// 	todo := make([]Todo, 0)
// 	vals := []int{0, 1, 2, 3, 4}
// 	currentperm := make([]int, len(vals))
// 	permcollector := make([][]int, 0)
// 	for i := range vals {
// 		valsLeft := make([]int, len(vals)-1)
// 		vlwrite := 0
// 		for ii := range vals {
// 			if i != ii {
// 				valsLeft[vlwrite] = vals[ii]
// 				vlwrite = vlwrite + 1
// 			}
// 		}
// 		log.Printf("current val %d vals left %v\n", vals[i], valsLeft)
// 		stillNeeded = append(stillNeeded, Todo{
// 			Index:     0,
// 			MyVal:     vals[i],
// 			ValsLeft:  valsLeft,
// 			PutOnWith: []int{0, 0, 0, 0, 0},
// 		})
// 	}
// 	for len(stillNeeded) > 0 {
// 		//grab a todo
// 		todo = append(todo, stillNeeded[len(stillNeeded)-1])
// 		stillNeeded = stillNeeded[:len(stillNeeded)-1]
// 		log.Printf("grabbing next todo off of needed %v \n", todo[len(todo)-1])

// 		for len(todo) > 0 {
// 			//grab a todo
// 			currenttodo := todo[len(todo)-1]
// 			todo = todo[:len(todo)-1]
// 			currentperm[currenttodo.Index] = currenttodo.MyVal
// 			log.Printf("working %v currentperm is %v \n", currenttodo, currentperm)
// 			//check if your at the end of current perm
// 			if len(currentperm)-1 == currenttodo.Index {
// 				//add this perm to the perm collector
// 				permcopy := make([]int, len(currentperm))
// 				copy(permcopy, currentperm)
// 				permcollector = append(permcollector, permcopy)
// 				log.Printf("HHHHHHHHHHHHHH %v\n", permcopy)
// 			} else {
// 				//add todo for each
// 				for nexttodoindex := range currenttodo.ValsLeft {
// 					valsleft := make([]int, len(currenttodo.ValsLeft)-1)
// 					valsleftwriteindex := 0
// 					for ii := range currenttodo.ValsLeft {
// 						if ii != nexttodoindex {
// 							// if valsleftwriteindex == len(valsleft){
// 							// 	log.Printf("")
// 							// }
// 							valsleft[valsleftwriteindex] = currenttodo.ValsLeft[ii]
// 							valsleftwriteindex = valsleftwriteindex + 1
// 						}
// 					}
// 					log.Printf("current val %d vals left %v, previous vals %v\n", currenttodo.ValsLeft[nexttodoindex], valsleft, currentperm)
// 					//append all but the last to needed, and the last to todo
// 					mcopy := make([]int, len(currentperm)-1)
// 					copy(mcopy, currentperm)
// 					if nexttodoindex == len(currenttodo.ValsLeft)-1 {
// 						todo = append(todo, Todo{
// 							MyVal:     currenttodo.ValsLeft[nexttodoindex],
// 							ValsLeft:  valsleft,
// 							Index:     currenttodo.Index + 1,
// 							PutOnWith: mcopy,
// 						})
// 					} else {
// 						stillNeeded = append(stillNeeded, Todo{
// 							MyVal:     currenttodo.ValsLeft[nexttodoindex],
// 							ValsLeft:  valsleft,
// 							Index:     currenttodo.Index + 1,
// 							PutOnWith: mcopy,
// 						})
// 					}

// 				}
// 			}

// 		}
// 	}
// 	log.Println(permcollector)
// }

// // type Needed struct{
// // 	Index int
// // 	MyVal int
// // 	ValsLeft []int
// // }
// type Todo struct {
// 	Index     int
// 	MyVal     int
// 	ValsLeft  []int
// 	PutOnWith []int
// }
