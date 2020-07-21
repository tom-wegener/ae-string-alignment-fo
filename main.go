package main

import (
	"log"
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

// verticesCount = n+1

func main() {
	var c Child

	verticesCount, _, Aij, Bij, Cij := parseFile("input/n=10/CCNFP10g1a.txt")
	//inputToCustomersStruct(verticesCount, customerDemand)
	costsA := inputToGraph(verticesCount, Aij)

	costsB := inputToGraph(verticesCount, Bij)

	costsC := inputToGraph(verticesCount, Cij)

	c.initiateFlow(verticesCount)
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)
	c.mutate()
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)
	c.mutate()
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)
	c.mutate()
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)
	c.mutate()
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)
	c.mutate()
	c.costEstimatorOne(costsA, costsB, costsC, verticesCount)
	println(c.fitness)

}

func errFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
