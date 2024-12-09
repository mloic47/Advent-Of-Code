package corruptedMemory

import (
	"bufio"
	"fmt"
	"os"
)

// goal: multiply some numbers in form mul(x,y)
// ingnore invalid characters

func copyFileContentToAList(filePath string)[]string {
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

