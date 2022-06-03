package main

import "fmt"

func main() {
	var number int
	var guessNumber int

	guessNumber = 51

	// fmt.Printf("guessNumber = %v\n", guessNumber)
	// fmt.Printf("number = %v\n", number)

	for number != guessNumber {
		fmt.Print("Enter number: ")
		fmt.Scan(&number)
		
		if number >= 1 && number <= 150 {
			if number < guessNumber{
				fmt.Println("Number is small")
			}

			if number > guessNumber{
				fmt.Println("Number is big")
			}

			if number == guessNumber{
				fmt.Println("Success")
			}
		} else {
			fmt.Println("Number incorrect")
		}
	}
}
