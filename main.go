package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("scanning error:", scanner.Err())
			continue
		}
		user_input := scanner.Text()
		cleaned_input := cleanInput(user_input)
		first_word := ""
		if len(cleaned_input) > 0 {
			first_word = cleaned_input[0]
		}
		fmt.Println("Your command was:", first_word)
	}
}
