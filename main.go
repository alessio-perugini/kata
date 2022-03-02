package main

//Game of Life

//	A cell evolves according to these rules:
//	A living cell, if it has less than 2 living cells next to it, dies in the next generation
//	A living cell, if it has 2 or 3 living cells next to it, remains alive in the next generation
//	A living cell, if it has more than 3 living neighbors, dies in the next generation
//	A dead cell, if it has 3 living neighbors, resurrects in the next generation

func main() {

}

type Universe struct {
	universe []uint8
}

func NewUniverse(universe []uint8) *Universe {
	return &Universe{universe: universe}
}

func (u *Universe) Evolve() {

}

func (u *Universe) LivingNeighbors(x, y uint8) uint8 {
	// 1 1 1 1 1 1 1 1
	// 1 X 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	// 1 1 1 1 1 1 1 1
	var top, bottom, left, right bool
	if x-1 > 0 {
		top = isAlive(u.universe[x], y)
	}

	if x+1 <= uint8(len(u.universe)-1) {
		bottom = isAlive(u.universe[x], y)
	}

	if y < 8 && y > 0 {
		left = isAlive(u.universe[x], y-1)
	}

	if y > 1 && y < 9 {
		right = isAlive(u.universe[x], y+1)
	}

	var total uint8
	if top {
		total++
	}
	if bottom {
		total++
	}
	if left {
		total++
	}
	if right {
		total++
	}

	return total
}

func isAlive(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func setBit(n uint8, pos uint8) uint8 {
	n |= (1 << pos)
	return n
}
