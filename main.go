package main

import (
	"log"
	"os"
	"sort"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Config is the config-struct
type Config struct {
	Input                string
	Generations          int
	HillclimbGenerations int
	Initiate             string
	Mutate               string
	Dot                  bool
	PopulationSize       int
	MutationDruck        float64
	Crossover            string
	Algorithmus          string
	TurnierGegner        int
}

var cfg Config

func main() {

	readConfig(&cfg)

	// Extract the date out of the File and check if there could be errors
	verticesCount, customerDemand, Aij, Bij, Cij := parseFile(cfg.Input)
	costsA := inputToGraph(verticesCount, Aij)
	costsB := inputToGraph(verticesCount, Bij)
	costsC := inputToGraph(verticesCount, Cij)
	network, err := createNetwork(costsA, costsB, costsC)
	errFunc(err)

	// Generate a Graph based on the costs of A
	makeGraph(costsA)

	// Calculate the demand which is also the capacity of the source
	// The demand could be seen as negative storage capacity
	// This fact will be used later on
	// Also the last node is the source
	var sourceCapacity int64
	var demand []int64
	for i := range customerDemand {
		sourceCapacity = sourceCapacity + customerDemand[i]
		demand = append(demand, -1*customerDemand[i])
	}
	demand = append(demand, sourceCapacity)

	var csvList [][]int64
	if cfg.Algorithmus == "hillclimber" {
		csvList = hillclimb(verticesCount, demand, network, costsA, costsB, costsC)
	} else if cfg.Algorithmus == "evolutionär" {
		csvList = geneticAlgorithm(verticesCount, demand, network, costsA, costsB, costsC)
	}
	var csvString string
	for _, row := range csvList {
		csvString = csvString + strconv.FormatInt(row[0], 10) + "," + strconv.FormatInt(row[1], 10) + ",\n"
	}
	f, err := os.Create("report/tmp.csv")
	errFunc(err)
	defer f.Close()
	_, err = f.WriteString(csvString)
	errFunc(err)

}

func printChild(x Child, n int) {
	print(n, ",")
	for i := range x.storage {
		print(x.storage[i], ",")
	}
	println(x.fitness, ",")
}

func geneticAlgorithm(verticesCount int, demand []int64, network [][]bool, costsA, costsB, costsC [][]int64) (csvList [][]int64) {
	var population []Child
	population = populate(network, verticesCount, demand, costsA, costsB, costsC)
	population = selectionTurnier(population)
	for i := 0; i < cfg.Generations; i++ {
		population = selectionTurnier(population)
		population = evolution(population, costsA, costsB, costsC, network)

		// Output best child of this generation
		var pool []Child
		copy(pool, population)
		sort.Sort(ByFitness(pool))
		csvList = append(csvList, []int64{int64(i + 1), population[0].fitness})
		printChild(population[0], i+1)
	}
	return
}

func hillclimb(verticesCount int, demand []int64, network [][]bool, costsA, costsB, costsC [][]int64) (csvList [][]int64) {

	c := new(Child) // Always Child
	x := new(Child) // Always parent
	x.demand = make([]int64, verticesCount)
	copy(x.demand, demand)

	// Initiate the flow but make it dependend from the config
	if cfg.Initiate == "zero" {
		x.initiateFlowZero(verticesCount)
	} else if cfg.Initiate == "one" {
		x.initiateFlowOne(verticesCount, network)
	} else if cfg.Initiate == "two" {
		x.initiateFlowTwo(verticesCount, network)
	}

	for _, storage := range x.storage {
		print(storage, ",")
	}
	x.costCalculator(costsA, costsB, costsC)
	println(x.fitness, ",")

	for i := 0; i < cfg.HillclimbGenerations; i++ {
		x.findNeighbourOne(c)
		c.costCalculator(costsA, costsB, costsC)
		csvList = append(csvList, []int64{int64(i + 1), c.fitness})
		for k := range x.storage {
			print(x.storage[k], ", ")
		}
		println(x.fitness, ",")
		if c.fitness < x.fitness {
			c.toParent(x)
		}
	}
	return csvList
}

func readConfig(cfg *Config) {
	f, err := os.Open("cfg.yml")
	errFunc(err)
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	errFunc(err)
}

func makeGraph(network [][]int64) {
	graph := `digraph graphname
	{
`
	for i, row := range network {
		for j := range row {
			if network[i][j] != 0 {
				graph = graph + "    " + strconv.Itoa(i) + " -> " + strconv.Itoa(j) + "[ label=" + strconv.FormatInt(network[i][j], 10) + "];\n"
			}
		}
	}
	graph = graph + "}"

	// Dump dot-graph to file
	f, err := os.Create("input.dot")
	errFunc(err)
	_, err = f.WriteString(graph)
	errFunc(err)
}

func errFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
