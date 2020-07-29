package main

import (
	"math/rand"
	"os"
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
	gegner := 3

	for i := 0; i < len(pool); i++ {
		for j := 0; j < gegner; j++ {
			rI := rand.Intn(len(pool))
			if pool[i].fitness < pool[rI].fitness {
				pool[rI] = pool[i]
			}
		}
	}
	return
}
func evolution(population []Child, costA, costB, costC [][]int64, network [][]bool) (pool []Child) {
	pool = make([]Child, 0)
	for i := 0; i < cfg.PopulationSize; i++ {
		var x Child
		var c Child
		r1 := rand.Intn(len(population))
		r2 := rand.Intn(len(population))
		x = population[r1]
		c = population[r2]
		c.onePointCrossover(&x)

		c.betterMutate(network)

		c.costCalculator(costA, costB, costC)
		if c.fitness < 0 {
			println(c.fitness)
			for i, row := range c.flow {
				print(c.storage[i], ", ")
				for j := range row {
					print(c.flow[i][j], ", ")
				}
				println()
			}
			os.Exit(0)

		}
		pool = append(pool, c)
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
			if rand.Float64() < cfg.MutationRate {
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

	var edges [][]int
	for i, row := range network {
		for j, cell := range row {
			if cell {
				edge := []int{i, j}
				edges = append(edges, edge)
			}
		}
	}
	for i := range edges {
		if rand.Float64() < cfg.MutationRate {

			for i := edges[i][0]; i >= 0; i-- {
				for j := edges[i][1]; j >= 0; j-- {
					if localStorage[i] > 0 {
						randomInt := rand.Int63n(localStorage[i])
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
