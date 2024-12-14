package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	idListA, idListB := readInput()
	sort.Ints(idListA)
	sort.Ints(idListB)

	totalDist := 0
	for i := range idListA {
		if idListA[i] > idListB[i] {
			totalDist += idListA[i] - idListB[i]
		} else {
			totalDist += idListB[i] - idListA[i]
		}
	}

	fmt.Printf("Total distance between lists is: %d\n", totalDist)

	countIdListB := make(map[int]int)
	for _, idB := range idListB {
		val, ok := countIdListB[idB]
		if !ok {
			countIdListB[idB] = idB
		} else {
			countIdListB[idB] = val + idB
		}
	}

	similarityScore := 0
	for _, idA := range idListA {
		similarityScore += countIdListB[idA]
	}

	fmt.Printf("Similarity score between lists is: %d\n", similarityScore)

}

func readInput() ([]int, []int) {
	file, err := os.Open("2024/day_01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var idListA, idListB []int
	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), "   ")
		idA, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatal(err)
		}
		idListA = append(idListA, idA)

		idB, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatal(err)
		}
		idListB = append(idListB, idB)
	}

	return idListA, idListB
}
