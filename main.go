package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type edgeStruct struct {
	X    int
	Y    int
	Cost int
}

type customerStruct struct {
	ID     int
	Demand int
}

// Config is the config-struct
type Config struct {
	input       string
	generations int
	initiate    string
	estimator   string
	mutate      string
}

// verticesCount = n+1
var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	var cfg Config
	readConfig(&cfg)

	fmt.Printf("%+v", cfg)

	var c Child // Always Child
	var x Child // Always parent
	maxGenerations := cfg.generations

	// Extract the date out of the File
	verticesCount, customerDemand, Aij, Bij, Cij := parseFile(cfg.input)
	costsA := inputToGraph(verticesCount, Aij)
	costsB := inputToGraph(verticesCount, Bij)
	costsC := inputToGraph(verticesCount, Cij)
	for i, row := range costsA {
		for j := range row {
			print(costsA[i][j], ",")
		}
		println()
	}
	network, err := createNetwork(costsA, costsB, costsC)
	errFunc(err)
	if cfg.initiate == "dumb" {
		x.initiateFlowDumb(verticesCount)
	} else if cfg.initiate == "smarter" {
		x.initiateFlowSmarter(verticesCount, customerDemand, network)
	}

	x.costCalculator(costsA, costsB, costsC, customerDemand)
	println(x.fitness)

	for i := 0; i < maxGenerations; i++ {
		x.findNeighbour(&c)
		c.costCalculator(costsA, costsB, costsC, customerDemand)
		if c.fitness < x.fitness {
			c.toParent(&x)
			println(c.fitness)
		}
	}

}

func readConfig(cfg *Config) {
	f, err := os.Open("cfg.yml")
	errFunc(err)
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	err = yaml.Unmarshal([]byte(data), cfg)
	fmt.Printf("--- t:\n%v\n\n", cfg)
	errFunc(err)

	t := T{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}

func errFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
