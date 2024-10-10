package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	word, err := getRandomWord("words.txt")
	if err != nil {
		fmt.Println("Error reading word file:", err)
		return
	}
	attempts := 6
	currentWordState := initializeCurrentWordState(word)
	scanner := bufio.NewScanner(os.Stdin)
	guessedLetters := make(map[string]bool)

	fmt.Println("Welcome to Hangman!")
	for attempts > 0 {
		displayCurrentState(currentWordState, attempts)
		userInput := getUserInput(scanner)
		if !isValidInput(userInput) {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}

		if guessedLetters[userInput] {
			fmt.Println("You've already guessed that letter.")
			continue
		}

		guessedLetters[userInput] = true
		correctGuess := updateGuessed(word, currentWordState, userInput)

		if !correctGuess {
			attempts--
		}

		displayHangman(6 - attempts)

		if isWordGuessed(currentWordState, word) {
			fmt.Println("Congratulations! You've guessed the word:", word)
			return
		}

		if attempts == 0 {
			fmt.Println("Game over! The word was:", word)
			return
		}
	}
}

func getRandomWord(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	words := strings.Split(string(data), "\n")
	return words[rand.Intn(len(words))], nil
}

func isWordGuessed(guessed []string, word string) bool {
	return strings.Join(guessed, "") == word
}

func displayHangman(incorrectGuesses int) {
	if incorrectGuesses >= 0 && incorrectGuesses < len(hangmanStates) {
		fmt.Println(hangmanStates[incorrectGuesses])
	}
}

func updateGuessed(word string, guessed []string, letter string) bool {
	correctGuess := false
	for i, char := range word {
		if string(char) == letter {
			guessed[i] = letter
			correctGuess = true
		}
	}
	return correctGuess
}

func isValidInput(input string) bool {
	return utf8.RuneCountInString(input) == 1
}

func getUserInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func initializeCurrentWordState(word string) []string {
	currentWordState := make([]string, len(word))
	for i := range currentWordState {
		currentWordState[i] = "_"
	}
	return currentWordState
}

func displayCurrentState(currentWordState []string, attempts int) {
	fmt.Println("Current word state:", strings.Join(currentWordState, " "))
	fmt.Println("Attempts left:", attempts)
}
