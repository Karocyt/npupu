package solver

import (
	"errors"
	"fmt"
)

/*Solve Function:
** Solves solver.openedStates given puzzle with A* algorithm.
** 	- 1st argument: solver.openedStates grid in the format [N*N]int
** 	- 2nd argument: size N of the aforementioned grid
** 	- 3rd argument: score function of type 'func([]int) int' used as heuristic
** return value: error e (unsolvable)
 */

func bestScore (l []*gridState) (cur *gridState) {
	for _, item := range l {
		if cur == nil || item.score < cur.score {
			cur = item
		}
	}
	return cur
}

func (solver *Solver) findState (state *gridState) int {
	for i, find := range  solver.openedStates {
		if find == state {
			return i
		}
	}
	return -1
}

func (solver *Solver) closeState (state *gridState) {
	idx := solver.findState(state)
	if idx == -1 {
		fmt.Println("NANI?! idx -1 ? ")
	}
	solver.openedStates[idx] = solver.openedStates[len(solver.openedStates)-1] 
	solver.openedStates = solver.openedStates[:len(solver.openedStates)-1]
}

func (solver *Solver) Solve() (e error) {
	cur := solver.openedStates[0]
	for cur != nil && cur.score != 0 {
		nextStates := cur.generateNextStates()
		solver.explored[cur.mapKey()] = true

		for _, child := range  nextStates {
			if solver.explored[child.mapKey()] == false {
				child.score = solver.fn(child.grid, size, child.depth)
				solver.openedStates = append(solver.openedStates, &child)
			}
		}
		solver.closeState(cur)
		cur = bestScore(solver.openedStates)
	}
	if len(solver.openedStates) == 0 {
		return errors.New("Error: pupu not solvable(empty open states)")
	}
	solver.Solution = make([]*gridState, 1)
	solver.Solution[0] = cur
	fmt.Println(cur)
	return
}