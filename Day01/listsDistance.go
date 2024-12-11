package main
// package day01

import (
	"bufio"
	"fmt"
	"os"
	// "path/filepath"
	"sort"
	"strconv"
	"strings"
)

func sortAndPrintListWithSort(list []int) []int {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	// print the sorted list
	for _, v := range list {
		fmt.Println(v)

	}
	// return the sorted list
	return list
}

func countOccurences(slice []int, value int)int{
	count := 0
	for _, v := range slice {
		if v == value {
			count ++
		}
	}
	return count
}

func listDifference (list1 []int, list2 []int) []int {
	var differences []int
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			difference := list1[i] - list2[i]
			differences = append(differences, difference)
		} else {
			difference := list2[i] - list1[i]
			differences = append(differences, difference)
		}
	}
	return differences
}

func sumSliceElement (slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func similarityScore(list1 []int, list2 []int) []int {
	var similarityScoreList []int
	for i := 0; i < len(list1); i++ {
		count := countOccurences(list2, list1[i])
		similarityScore := list1[i] * count
		similarityScoreList = append(similarityScoreList, similarityScore)
	}
	return similarityScoreList
}

// function that puts content of in file in a list
func CopyFileContentToList(filePath string)[]string {
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

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// create a list to store the content of file
	var list []string
	// read the file line by line
	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		// add the numbers to the lists
		list = append(list, line)
	}
	return list 
}




func main() {
	// open the file
	listDistanceInputFile, err := os.Open("listsDistanceInput.txt")
	// check if there is an error opening the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		defer listDistanceInputFile.Close()
		return
	}
	fmt.Println(listDistanceInputFile)
	fmt.Println("File opened successfully")

	// create a scanner to read the file
	scanner := bufio.NewScanner(listDistanceInputFile)
	// create 2 list to store the numbers
	var list1 []int
	var list2 []int
	// read the file line by line
	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		// the 2 numbers are separated by a space
		numbers := strings.Split(line, "   ")
		// convert the numbers to integers
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		// add the numbers to the lists
		list1 = append(list1, num1)
		list2 = append(list2, num2)

	
		// sort the lists
		sortedList1 := sortAndPrintListWithSort(list1)
		sortedList2 := sortAndPrintListWithSort(list2)

		// calculate the differences
		differences := listDifference(sortedList1, sortedList2)

		// print the differences
		fmt.Println("Differences:", differences)

		// sum up differences
		sum := sumSliceElement(differences)
		fmt.Println("Sum of differences:", sum)
		
		// calculating total similarity score
		similarityScoreList := similarityScore(sortedList1, sortedList2)
		// sum up the similarity score
		totalSimilarityScore := sumSliceElement(similarityScoreList)
		fmt.Println("Total similarity score:", totalSimilarityScore)
	}
}
