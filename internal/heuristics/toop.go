package heuristics

func t_score(grid []int, nb int, x int, y int, size int) float32 {
	x1, y1 := get_2d(nb, size)
	nb1 := grid[get_1d(x1 , y1, size)]
	if nb == nb1 {
		return 1
	} else {
		return 0
	}
	return 0
}

// Tiles out-of place
func toop(grid []int, size int) float32 {
	var score float32
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += t_score(grid, grid[get_1d(x , y, size)], x, y, size)
		}
	}
	return score
}
