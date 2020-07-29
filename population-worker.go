package main

import (
	"math/rand"
	"os"
	"sort"
)

// ByFitness implements sort-Interface for the Childs based on their fitness
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

func ranking(population []Child) (pool []Child) {
	sort.Sort(ByFitness(population))
	numOfRanks := len(population)
	for i, individuum := range population {
		rank := i / 2 / numOfRanks * (1 - (i / numOfRanks))

		for j := numOfRanks; j > rank; j-- {
			pool = append(pool, individuum)
		}
	}
	return
}

func selection(population []Child, costA, costB, costC [][]int64) (pool []Child) {
	for i := 0; i < cfg.PopulationSize; i++ {
		r1 := rand.Intn(len(population))
		r2 := rand.Intn(len(population))
		x := population[r1]
		c := population[r2]
		c.twoPointCrossover(&x)
		c.randomMutate()
		c.costCalculator(costA, costB, costC)
		if c.fitness < 0 {
			printChild(x, 200)
			for i, row := range x.flow {
				for j := range row {
					print(x.flow[i][j], ",")
				}
				println()
			}
			os.Exit(0)
		}
		pool = append(pool, c)
	}
	return
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
				x.flow[i][j] = c.flow[i][j]
			}
			n++
		}
	}

}

func (x *Child) randomMutate() {
	for i, row := range x.flow {
		for j := range row {
			if rand.Float64() < cfg.MutationRate {
				if x.storage[i]+x.flow[i][j] > 0 && x.flow[i][j] > 0 {
					x.storage[i] = x.storage[i] + x.flow[i][j]
					x.storage[j] = x.storage[j] - x.flow[i][j]
					x.flow[i][j] = rand.Int63n(x.storage[i])
					x.storage[i] = x.storage[i] - x.flow[i][j]
					x.storage[j] = x.storage[j] + x.flow[i][j]
				}

			}
		}
	}
}
