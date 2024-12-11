package corruptedMemory

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "strconv"
	"strings"
)

var regex = `mul\(\d{1,9},\d{1,9}\)`

func CopyFileContentToAList(filePath string)[]string {
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
		// loop through the line and collect numbers to be multiplied
		for _, value := range line {
			list = append(list, string(value))
		}			
		// add the numbers to the lists
		// list = append(list, line)
	}
	fmt.Println(list)
	return list 
}

func SumOfMultiplicationExcludingMulsThatProceedDonts(list []string) int {
	matches := regexp.MustCompile(`mul\(\d{1,9},\d{1,9}\)|don't\(\)|do\(\)`).FindAllString(strings.Join(list, ""), -1)
	// reDont := regexp.MustCompile(`don't\(\)`)
	fmt.Println("list of muls without dont",matches)
	// in the list remove don't() and the next indexs till the next do()
	var removed []string
	removeEnabling := true
	for _, value := range matches {
		// if a value is after a don't() remove it
		switch value {
		case "don't()":
			removeEnabling = false
		case "do()":
			removeEnabling = true
		default:
			if removeEnabling {
				removed = append(removed, value)
			}
		}
	}
	fmt.Println("list of muls without dont",removed)
	return SumOfMultiplication(removed)
}

	

func SumOfMultiplication(list []string) int {
	matches := regexp.MustCompile(regex).FindAllString(strings.Join(list, ""), -1)
	fmt.Println("list of muls",matches)
	// var ExtractedNumbersInt []int
	var ExtractNumbers []string
	var listOfMuls [][]int
	fmt.Println("list of muls",matches)
	for _, value := range matches {
		// if a value is after a don't() remove it
		

		ExtractNumbers = regexp.MustCompile(`\d{1,9}`).FindAllString(value, -1)
		
		var ExtractedNumbersInt []int
		for _, value := range ExtractNumbers {
			number, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return 0
			}
			ExtractedNumbersInt = append(ExtractedNumbersInt, number)
		}
		listOfMuls = append(listOfMuls, ExtractedNumbersInt)
	}
	fmt.Println("list of muls",listOfMuls)
	var sum int
	for _, value := range listOfMuls {
		sum += value[0] * value[1]
	}
	fmt.Println("sum of muls",sum)

	return sum
}