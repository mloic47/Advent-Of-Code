package main

import (
	"bufio"
	"fmt"
	"os"
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

func sortAndPrintListWithForLoop(list []int) []int {
	// iterate over the list
	for i := 0; i < len(list); i++ {
		// iterate over the list starting from the next element
		for j := i + 1; j < len(list); j++ {
			// if the current element is greater than the next element
			if list[i] > list[j] {
				// swap the elements
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	// print the list
	fmt.Println(list)
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

		// print the lists
		fmt.Println("List 1:", list1)
		fmt.Println("List 2:", list2)

		// sort the lists in ascending order
		fmt.Println("Sorted list 1:")
		if fmt.Sprint(sortAndPrintListWithSort(list1)) == fmt.Sprint(sortAndPrintListWithForLoop(list1)) {
			fmt.Println(sortAndPrintListWithSort(list1))
			fmt.Println(sortAndPrintListWithForLoop(list1))
			// compare the results
			if fmt.Sprint(sortAndPrintListWithSort(list1)) == fmt.Sprint(sortAndPrintListWithForLoop(list1)) {
				fmt.Println("Both results are the same")
			} else {
				fmt.Println("Both results are different")
			}
			// // sort and print list2
			fmt.Println("Sorted list 2:")
			sortAndPrintListWithSort(list2)
		}
	}
}
