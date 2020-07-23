package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// createNetwork: Check if the all matrices create the same network and create a neutral ground.
func createNetwork(a, b, c [][]int64) (network [][]bool, err error) {
	for i, row := range a {
		var networkX []bool
		for j := range row {
			if a[i][j] == 0 && b[i][j] == 50000000 && c[i][j] == 0 {
				networkX = append(networkX, false)
			} else if a[i][j] >= 0 && b[i][j] <= 50000000 && c[i][j] >= 0 {
				networkX = append(networkX, true)
			} else {
				err = errors.New("There are edges in one of your cost-matrices where there shouldn't be edges according to the other matrices")
			}
		}
		network = append(network, networkX)
	}
	return
}

func inputToGraph(verticesCount int, Aij []int64) (edges [][]int64) {

	var edgesAll [][]int64
	for i := 0; i < verticesCount; i++ {
		var edgesX []int64
		for j := 0; j < (verticesCount - 1); j++ {
			edgesX = append(edgesX, Aij[i*(verticesCount-1)+j])
		}
		edgesAll = append(edgesAll, edgesX)
	}
	edges = edgesAll
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
