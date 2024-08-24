package main

import (
	"fmt"
	"math/rand"
	"time"
)

func start(config *Config) {
	var input string

	for {
		fmt.Printf("\nPlease select the difficulty level:\n"+
			"1.Easy (%v chances)\n"+
			"2.Medium (%v chances)\n"+
			"3.Hard (%v chances)\n"+"4.Back\n", config.Easy, config.Medium, config.Hard)
		fmt.Printf("Enter Your Choice: ")
		fmt.Scanln(&input)
		if input == "1" || input == "2" || input == "3" {
			switch input {
			case "1":
				play("Easy", config, config.Easy)
			case "2":
				play("Medium", config, config.Medium)
			case "3":
				play("Hard", config, config.Hard)
			}
		} else if input == "4" {
			break
		} else {
			fmt.Printf("\nError: Invalid Input.\n")
		}
	}
}

func play(input string, config *Config, chances int) {
	number := rand.Intn(100) + 1
	var guess int
	counter := 0

	fmt.Printf("\nGreat! You have selected the %v difficulty level.\n"+
		"Let's start the game!\n", input)

	var elapsedTime time.Duration
	startTime := time.Now()

	for guess != number && chances != 0 {
		fmt.Printf("\nEnter your guess (%v Chances, ELapsedTime %v): ", chances, time.Since(startTime).Seconds())
		fmt.Scanln(&guess)
		if guess == 0 {
			fmt.Printf("\nError: Invalid Input.\n")
			// Clear the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		chances--
		counter++
		if guess > number {
			fmt.Printf("Incorrect! The number is less than %v.\n", guess)
		} else if guess < number {
			fmt.Printf("Incorrect! The number is greater than %v.\n", guess)
		} else if guess == number {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", counter)
			elapsedTime = time.Duration(time.Since(startTime).Seconds())
			break
		}
		if chances == 0 {
			fmt.Printf("\nYOU LOST! You have consumed all your %v attempts.\n", counter)
			elapsedTime = time.Duration(time.Since(startTime).Seconds())
		}
	}

	switch input {
	case "Easy":
		if counter < config.HScore[0].Attempts || config.HScore[0].Attempts == 0 {
			config.HScore[0].Time = elapsedTime
			config.HScore[0].Attempts = counter
			config.HScore[0].Date = time.Now().Format(time.DateOnly)
			config.saveChanges()
		} else if counter == config.HScore[0].Attempts {
			if elapsedTime < config.HScore[0].Time {
				config.HScore[0].Time = elapsedTime
				config.HScore[0].Attempts = counter
				config.HScore[0].Date = time.Now().Format(time.DateOnly)
				config.saveChanges()
			}
		}
	case "Medium":
		if counter < config.HScore[1].Attempts || config.HScore[1].Attempts == 0 {
			config.HScore[1].Time = elapsedTime
			config.HScore[1].Attempts = counter
			config.HScore[1].Date = time.Now().Format(time.DateOnly)
			config.saveChanges()
		} else if counter == config.HScore[1].Attempts {
			if elapsedTime < config.HScore[1].Time {
				config.HScore[1].Time = elapsedTime
				config.HScore[1].Attempts = counter
				config.HScore[1].Date = time.Now().Format(time.DateOnly)
				config.saveChanges()
			}
		}
	case "Hard":
		if counter < config.HScore[2].Attempts || config.HScore[2].Attempts == 0 {
			config.HScore[2].Time = elapsedTime
			config.HScore[2].Attempts = counter
			config.HScore[2].Date = time.Now().Format(time.DateOnly)
			config.saveChanges()
		} else if counter == config.HScore[2].Attempts {
			if elapsedTime < config.HScore[2].Time {
				config.HScore[2].Time = elapsedTime
				config.HScore[2].Attempts = counter
				config.HScore[2].Date = time.Now().Format(time.DateOnly)
				config.saveChanges()
			}
		}
	}

}
