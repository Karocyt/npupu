package heuristics

func mScore(nb int, x int, y int, size int) float32 {
	tmp := nb_pos[nb]
	x1 := tmp[0]
	y1 := tmp[1]
	return float32(abs((x1 - x)) + abs((y1 - y)))
}

func manhattan(grid []int, size int) float32 {
	var score float32
	_, nb_pos = makeGoal(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += mScore(grid[get1d(x, y, size)], x, y, size)
		}
	}
	return score
}
