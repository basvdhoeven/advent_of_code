package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := readInput()
	safeReports := 0
	for _, report := range reports {
		if safeReport(report) {
			safeReports += 1
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeReports)

	safeReportsWithProblemDampener := 0
	for _, report := range reports {
		if safeReportWithProblemDampener(report) {
			safeReportsWithProblemDampener += 1
		}
	}

	fmt.Printf("Number of safe reports with problem dampener: %d\n", safeReportsWithProblemDampener)
}

func safeReportWithProblemDampener(report []int) bool {
	for i := range report {
		var reportWithMissingLevel []int
		for j := range report {
			if i != j {
				reportWithMissingLevel = append(reportWithMissingLevel, report[j])
			}
		}

		if safeReport(reportWithMissingLevel) {
			return true
		}
	}

	return false
}

// safe report:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func safeReport(report []int) bool {
	increasing := report[1] > report[0]
	for i := range report {
		if i == 0 {
			continue
		}
		difference := report[i] - report[i-1]
		if increasing {
			if difference < 1 || difference > 3 {
				return false
			}
		} else {
			if difference > -1 || difference < -3 {
				return false
			}
		}
	}

	return true
}

func readInput() [][]int {
	file, err := os.Open("2024/day_02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		intLevels := make([]int, len(levels))
		for i, s := range levels {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			intLevels[i] = num
		}
		reports = append(reports, intLevels)
	}

	return reports
}
