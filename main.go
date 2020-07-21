package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	X     int
	Y     int
	costA int
	costB int
	costC int
}

type customer struct {
	demand int
}

func main() {
	parseFiletoStructs("input/n=10/CCNFP10g1a.txt")
}

func parseFiletoStructs(inPath string) {
	file, err := os.Open(inPath)
	errFunc(err)
	defer file.Close()

	step := 0
	scanner := bufio.NewScanner(file)
	var customerDemand []int
	var Aij []int
	var Bij []int
	var Cij []int
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Cost") {
			step = step + 1
			fmt.Println(step)
		} else if step == 0 {
			// Extract Customer Demands
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			customerDemand = append(customerDemand, intLine)

			// Extract Variable Costs Component A
		} else if step == 1 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			Aij = append(Aij, intLine)

			// Extract Variable Cost Component B
		} else if step == 2 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			Bij = append(Bij, intLine)

			// Extract fixed Cost Component C
		} else if step == 3 {
			trimmedLine := strings.TrimSpace(scanner.Text())
			intLine, err := strconv.Atoi(trimmedLine)
			errFunc(err)
			Cij = append(Cij, intLine)
		}
	}
	errFunc(scanner.Err())
}

func errFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
