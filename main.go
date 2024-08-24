package main

import "fmt"

func main() {
	config := &Config{}
	config.initConfigFile()

	fmt.Printf("Welcome to the Number Guessing Game!\n" +
		"I'm thinking of a number between 1 and 100.\n")

	var input string

	for {
		fmt.Printf("\nChoose one of the following:\n" +
			"1.Start Game\n" +
			"2.Exit\n")
		fmt.Printf("Enter Your Choice: ")
		fmt.Scanln(&input)

		if input == "1" {
			start(config)
		} else if input == "2" {
			break
		} else {
			fmt.Printf("\nError: Invalid Input.\n")
		}
	}
}
