package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

// ReadString prompts the user and returns their raw input as a string.
func ReadString(prompt string) string {
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}

// ReadInt prompts the user until they enter a valid integer within [min, max].
func ReadInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt + ": ")
		scanner.Scan()
		input := scanner.Text()

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		if value < min || value > max {
			fmt.Printf("Please enter a number between %d and %d.\n", min, max)
			continue
		}
		return value
	}
}
