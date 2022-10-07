package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := os.Args
	if len(input) < 2 {
		fmt.Println("Please enter some text")
		return
	}

	//read file
	lines := readfile("standard.txt")

	//convert input string to runes slice
	runes := []rune(input[1])

	// nested loop to print line by line depending on input.
	splittwo := "\\n"
	words := strings.Split(string(runes), splittwo)
	for _, word := range words {

		if word == "" { // the new line "\\n" was at the end of "words" slice, and Split create the "" word
			fmt.Println()
		} else { // usual case letter print
			// iterate of letter art high, it is like vertical step to print horizontal sequences of letter ascii art
			for h := 1; h < 9; h++ { // from one to ignore the empty new line from standart.txt
				for _, l := range []byte(word) { // get the letter from word in form of byte

					ind := (int(l)-32)*9 + h // potential index (the height from up to bottom) in "lines" for required letter line(because art letter is multilined)
					if ind < len(lines) {    // check the index is inside available ascii art symbols ... f.e. standart.txt
						fmt.Print(lines[ind]) // print the line from high "h" for the word letter "l"
					}

					// slow old version of check

					// for i, line := range lines {
					// 	if i == (int(l)-32)*9+h {
					// 		fmt.Print(line)
					// 		// fmt.Print(i)
					// 	}
					// }

				}

				fmt.Println()

			}
		}
	}
}

// function to read file
func readfile(name string) []string {
	//read full file
	file, err := os.Open(name)
	//error check
	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	} else {
		// fmt.Println("File opened successfully")
	}
	//close file
	defer file.Close()
	//reads full file content to data as bytes
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error!")
		os.Exit(2)
	}

	lines := strings.Split(string(data), "\n")

	return lines
}
