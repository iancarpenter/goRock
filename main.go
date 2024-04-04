package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

// function to get user input
func getUserInput(reader io.Reader) string {
	fmt.Print("Enter rock, paper, or scissors: ")
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}

// function to get computer input
func getComputerInput() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

// function to compare user input and computer input
func compareInput(userInput string, computerInput string) string {

	if userInput == computerInput {
		return "draw"
	}
	if userInput == "rock" && computerInput == "scissors" {
		return "win"
	}
	if userInput == "paper" && computerInput == "rock" {
		return "win"
	}
	if userInput == "scissors" && computerInput == "paper" {
		return "win"
	}
	return "lose"
}

func main() {
	// initialize score
	score := 0

	// loop to play the game
	for {
		// get user input
		userInput := getUserInput(os.Stdin)

		// get computer input
		computerInput := getComputerInput()

		// compare user input and computer input
		result := compareInput(userInput, computerInput)

		// print the result
		fmt.Printf("You choose %s, computer choose %s, you %s\n", userInput, computerInput, result)

		// update score
		if result == "win" {
			score++
		}

		// ask user to play again or not
		fmt.Print("Play again? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.Trim(playAgain, "\n")

		// if no, exit the game
		if playAgain == "no" {
			break
		}
	}

	// print the total score
	fmt.Printf("Total score: %d\n", score)
}
