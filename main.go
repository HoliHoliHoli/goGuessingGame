package main

// a CLI guessing game

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number!")

	// generate a random number
	// ":=" can only be used inside a function or else use "var"
	// ":=" declares a variable and assigns it a value
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) // generates random integer between 0 and n

	var guess int
	var i int = 1

	for {
		fmt.Println("Please input your guess between 0 and 20: ")
		fmt.Scan(&guess)
		if guess > (secretNumber + 3) {
			fmt.Println("Way too Big!")
		} else if guess > secretNumber {
			fmt.Println("Too Big!")
		} else if guess < (secretNumber - 3) {
			fmt.Println("Way too Small!")
		} else if guess < secretNumber {
			fmt.Println("Too Small!")
		} else {
			fmt.Printf("You Win! That took %d try(ies)", i)
			break
		}
		i++
	}

	/*
		fmt.Scan(&guess)

		fmt.Printf("You guessed %d\n", guess)

		if guess == secretNumber {
			fmt.Println("Correct! You win")
		} else {
			fmt.Printf("Wrong! You lose!, number is %d", secretNumber)
		}
	*/
}
