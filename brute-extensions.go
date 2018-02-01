package main

import (
	"fmt"
	"net/http"
	"os"
)

func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		AddLetter(c, "", alphabet, length) // We start by feeding it an empty string
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func AddLetter(c chan string, combo string, alphabet string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}
}

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		fmt.Println("Usaage: ./brute-extensions [URL]")
		fmt.Println("Example: ./brute-extensions http://www.example.com/testfile")
		os.Exit(1)
	}

	for combination := range GenerateCombinations("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 3) {
		res, err := http.Head(argsWithProg[1] + "." + combination)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Print(combination)
		// If we hit a 403, just erase the current combination we're trying
		if res.StatusCode == 403 {
			fmt.Print("\r")
		} else {
			fmt.Print("\r\n")
		}
	}
	fmt.Println("Done!")
}
