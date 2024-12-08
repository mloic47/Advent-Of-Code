package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func countSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafeReport(report) {
			count++
		}
	}
	return count
}

func toleratedSafeReport(report []int) bool {
	// report is safe if when removing any 1 number from an unsafe report it becomes safe
	if len(report) < 2 {
		return false
	}
	for i := 0; i < len(report); i++ {
		// remove the number at index i
		var newReport []int
		for j := 0; j < len(report); j++ {
			if j != i {
				newReport = append(newReport, report[j])
			}
		}
		if isSafeReport(newReport) {
			return true
		}
	}
	return false
}

func isSafeReport(report []int) bool {
	// report is safe if both are increasing or decreasing
	// and any 2 consecutive numbers differ by at least 1 and at most 3
	if len(report) < 2 {
		return false
	}
	increasing := true
	decreasing := true
	for i := 0; i < len(report)-1; i++ {
		difference := report[i+1] - report[i]

		if difference < -3 || difference > 3 || difference == 0 {
			return false
		}

		if difference > 0 {
			decreasing = false
		}else if difference < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}


// read input file and parse the reports
func readReports(filePath string) [][]int {
	// open the file
	file, err := os.Open(filePath)
	// check if there is an error opening the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		defer file.Close()
		return nil
	}
	fmt.Println(file)
	fmt.Println("File opened successfully")

	var reports [][]int
	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports
}


func main() {
	// read the reports from the file
	reports := readReports("input.txt")
	if reports == nil {
		fmt.Println("Error reading reports")
		return
	}
	fmt.Println(reports)

	// count the number of safe reports
	count := countSafeReports(reports)
	fmt.Println("Number of safe reports:", count)

	// count the number of tolerated safe reports
	toleratedCount := 0
	for _, report := range reports {
		if toleratedSafeReport(report) {
			toleratedCount++
		}
	}
	fmt.Println("Number of tolerated safe reports:", toleratedCount)
	
}

// TODO: make the count func take as argument the function that checks if a report is safe