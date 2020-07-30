package main

import (
	"math/rand"
	"sort"
)

// ByFitness implements sort-Interface for []Child based on their fitness
type ByFitness []Child

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].fitness < a[j].fitness }

func populate(network [][]bool, verticesCount int, demand []int64, costsA, costsB, costsC [][]int64) (population []Child) {

	population = make([]Child, cfg.PopulationSize)
	for i := range population {
		var x Child
		x.demand = demand
		x.initiateFlowTwo(verticesCount, network)
		x.costCalculator(costsA, costsB, costsC)
		population[i] = x
	}
	return
}

func selectionRanking(population []Child) (pool []Child) {
	pool = make([]Child, 0)
	sort.Sort(ByFitness(population))
	numOfRanks := float64(len(population))
	var i float64
	for i = 0; i < numOfRanks; i++ {
		rank := (2 / numOfRanks) * (1 - (i / (numOfRanks - 1))) * numOfRanks

		for j := float64(0); j < rank; j++ {
			pool = append(pool, population[int(i)])
		}
	}
	return
}

func selectionTurnier(population []Child) (pool []Child) {
	pool = population

	for i := 0; i < len(population); i++ {
		r1 := rand.Intn(len(population))
		for j := 0; j < cfg.TurnierGegner; j++ {
			r2 := rand.Intn(len(population))
			if population[r1].fitness < population[r2].fitness {
				pool[r1] = population[r1]
			}
		}
	}
	return
}
func evolution(population []Child, costA, costB, costC [][]int64, network [][]bool) (pool []Child) {
	pool = make([]Child, 0)
	for i := 0; i < cfg.PopulationSize; i = i + 2 {
		var x Child
		var c Child
		r1 := rand.Intn(len(population))
		r2 := rand.Intn(len(population))
		x = population[r1]
		c = population[r2]
		c.onePointCrossover(&x)

		c.betterMutate(network)
		x.betterMutate(network)

		c.costCalculator(costA, costB, costC)
		x.costCalculator(costA, costB, costC)
		pool = append(pool, c, x)
	}
	return
}

func (x *Child) onePointCrossover(c *Child) {

	childLen := len(x.flow) * len(x.flow[0])
	crossoverPoint := rand.Intn(childLen)

	n := 0
	for i, row := range x.flow {
		for j := range row {
			if crossoverPoint < n {
				tmp := c.flow[i][j]
				c.flow[i][j] = x.flow[i][j]
				x.flow[i][j] = tmp
			}
			n++
		}
	}

}

func (x *Child) twoPointCrossover(c *Child) {

	childLen := len(x.flow) * len(x.flow[0])
	crossoverPointOne := rand.Intn(childLen)
	crossoverPointTwo := rand.Intn(childLen)
	if crossoverPointOne > crossoverPointTwo {
		tmp := crossoverPointOne
		crossoverPointOne = crossoverPointTwo
		crossoverPointTwo = tmp
	}

	n := 0
	for i, row := range x.flow {
		for j := range row {
			if crossoverPointOne < n && n < crossoverPointTwo {

				tmp := c.flow[i][j] - x.flow[i][j]
				x.flow[i][j] = c.flow[i][j]
				x.storage[i] = x.storage[i] + tmp
				x.storage[i] = x.storage[j] - tmp
			}
			n++
		}
	}

}

func (x *Child) randomMutate(network [][]bool) {
	for i, row := range x.flow {
		for j := range row {
			if rand.Float64() < cfg.MutationDruck {
				storeI := x.storage[i] + x.flow[i][j]
				if storeI > 0 && network[i][j] {
					x.storage[i] = x.storage[i] + x.flow[i][j]
					x.storage[j] = x.storage[j] - x.flow[i][j]
					x.flow[i][j] = rand.Int63n(storeI)
					x.storage[i] = x.storage[i] - x.flow[i][j]
					x.storage[j] = x.storage[j] + x.flow[i][j]
				}

			}
		}
	}
}

func (x *Child) betterMutate(network [][]bool) {
	localStorage := x.demand

	// Get List of all edges
	var edges [][]int
	for i, row := range network {
		for j, cell := range row {
			if cell {
				edge := []int{i, j}
				edges = append(edges, edge)
			}
		}
	}

	//Loop through the list and mutate if it has to
	for k := range edges {
		if rand.Float64() < cfg.MutationDruck {

			// Loop though al edges and mutate all following edges according to the storage
			for i := len(x.flow) - 1; i >= edges[k][0]; i-- {
				for j := edges[k][1]; j < len(x.flow[i]); j++ {
					tmp := localStorage[i] + x.flow[i][j]
					if tmp > 0 {
						randomInt := rand.Int63n(tmp)
						localStorage[i] = localStorage[i] - randomInt
						localStorage[j] = localStorage[j] + randomInt
						x.flow[i][j] = randomInt
					}
				}
			}
			x.storage = localStorage
		}

	}

	return

}
