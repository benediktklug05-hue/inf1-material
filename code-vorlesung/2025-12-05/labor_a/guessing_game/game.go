package guessinggame

import "fmt"

func GuessingGame() {
	guess := ReadNumber()
	if NumberGood(guess) {
		fmt.Println("Richtig geraten :-)")
	}

	fmt.Println("Zu viele falsche Zahlen")
}
