package main

import (
	"fmt"

	"goLearn.com/adventOfCode/thirdday"
)

func main() {
	fmt.Println("-----------DAY 03--------------------------")
	corruptedList := corruptedMemory.CopyFileContentToAList("thirdday/input.text")
	// fmt.Println("txt file to list", corruptedList)
	// listOfChars := corruptedMemory.SumOfMultiplication(corruptedList)
	// fmt.Println("list of chars",listOfChars)
	fmt.Println("Sum of multiplication excluding muls that proceed don'ts", corruptedMemory.SumOfMultiplicationExcludingMulsThatProceedDonts(corruptedList))
}