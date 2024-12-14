package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	memory, err := os.ReadFile("2024/day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	memoryString := string(memory)
	mulContents := getMulContents(memoryString)

	fmt.Printf("Sum of multiplications is: %d\n", getSumOfMultiplications(mulContents))

	memoryDoOnly := getDoMemory(memoryString)
	mulContentsDo := getMulContents(memoryDoOnly)

	fmt.Printf("Sum of multiplications do instructions is: %d\n", getSumOfMultiplications(mulContentsDo))
}

func getDoMemory(memory string) string {
	do := true
	var doMemory string
	for {
		if do {
			dontIndex := strings.Index(memory, "don't()")
			if dontIndex == -1 {
				doMemory += memory
				break
			}

			doMemory += memory[:dontIndex]
			memory = memory[dontIndex+7:]
			do = false
			continue
		} else {
			doIndex := strings.Index(memory, "do()")
			if doIndex == -1 {
				break
			}

			do = true
			memory = memory[doIndex+4:]
		}
	}

	return doMemory
}

func getSumOfMultiplications(contents []string) int {
	sumOfMultiplications := 0
	for _, content := range contents {
		commaSeperatedContent := strings.Split(content, ",")
		if len(commaSeperatedContent) != 2 {
			continue
		}
		numberA, err := strconv.Atoi(commaSeperatedContent[0])
		if err != nil {
			continue
		}
		numberB, err := strconv.Atoi(commaSeperatedContent[1])
		if err != nil {
			continue
		}

		sumOfMultiplications += numberA * numberB
	}

	return sumOfMultiplications
}

func getMulContents(memory string) []string {
	var mulContents []string
	for {
		mulOpeningBracketIndex := strings.Index(memory, "mul(")
		if mulOpeningBracketIndex == -1 {
			break
		}
		contentStart := mulOpeningBracketIndex + 4

		closingBracketIndex := strings.Index(memory[mulOpeningBracketIndex+4:], ")")
		if closingBracketIndex == -1 {
			break
		}
		contentLength := closingBracketIndex

		mulContents = append(mulContents, memory[contentStart:contentStart+contentLength])
		memory = memory[contentStart:]
	}

	return mulContents
}
