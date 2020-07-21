package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func inputToGraph(verticesCount int, Aij []int64) (edges [][]int64) {
	var edgesX []int64
	var edgesAll [][]int64
	for i := 0; i < verticesCount; i++ {
		for j := 0; j < (verticesCount - 1); j++ {
			edgesX = append(edgesX, Aij[i*(verticesCount-1)+j])
		}
		edgesAll = append(edgesAll, edgesX)
	}
	edges = edgesAll
	return
}

func inputToCustomersStruct(verticesCount int, customerDemand []int) (customers []customerStruct) {
	for i := 1; i < verticesCount; i++ {
		var aCustomer customerStruct
		aCustomer.ID = i
		aCustomer.Demand = customerDemand[i-1]
		customers = append(customers, aCustomer)
	}
	return
}

func parseFile(inPath string) (verticesCount int, customerDemand, Aij, Bij, Cij []int64) {

	file, err := os.Open(inPath)
	errFunc(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	step := 0

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Cost") {
			step = step + 1
		} else if step == 0 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			verticesCount = intLine

			step = step + 1
		} else if step == 1 {
			// Extract Customer Demands
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			int64Line := int64(intLine)
			customerDemand = append(customerDemand, int64Line)

			// Extract Variable Costs Component A
		} else if step == 2 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			int64Line := int64(intLine)
			Aij = append(Aij, int64Line)

			// Extract Variable Cost Component B
		} else if step == 3 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			int64Line := int64(intLine)
			Bij = append(Bij, int64Line)

			// Extract fixed Cost Component C
		} else if step == 4 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			int64Line := int64(intLine)
			Cij = append(Cij, int64Line)
		}
	}
	errFunc(scanner.Err())
	return
}
