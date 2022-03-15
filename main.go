package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func genWord() string {
	readFile, err := os.Open("word_list.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var words []string

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	readFile.Close()

	rand.Seed(time.Now().Unix())

	word := words[rand.Intn(len(words))]

	return word
}

func gameloop(word string, foundLetters *list.List, guessedLetters *list.List, lives int) (int, bool, bool) {
	hangmen := []string{`  +---+
  |   |
      |
      |
      |
      |
=========
`,
		`  +---+
  |   |
  O   |
      |
      |
      |
=========
`,
		`  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`,
		`  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
`,
		`  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`,
		`  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`,
		`  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`}
	fmt.Print(hangmen[len(hangmen)-1-lives])
	fmt.Print("\n")

	fmt.Printf("Lives: %d", lives)
	fmt.Print("\n")

	fmt.Print("Word: ")
	for i := 0; i < len(word); i++ {
		char := word[i]
		if listContains(foundLetters, char) {
			fmt.Printf("%c ", char)
		} else {
			fmt.Print("_ ")
		}
	}

	fmt.Print("\n")

	fmt.Print("Found Letters: ")
	for foundLetter := foundLetters.Front(); foundLetter != nil; foundLetter = foundLetter.Next() {
		fmt.Printf("%c ", foundLetter.Value)
	}

	fmt.Print("\n")

	fmt.Print("Guessed Letters: ")
	for guessedLetter := guessedLetters.Front(); guessedLetter != nil; guessedLetter = guessedLetter.Next() {
		fmt.Printf("%c ", guessedLetter.Value)
	}

	fmt.Print("\n")

	var letter byte
	for letter < 97 || letter > 123 {
		var input string
		fmt.Print("Enter Letter: ")
		fmt.Scanf("%s", &input)
		letter = input[0]
		if len(input) > 1 {
			fmt.Println("You've inputed more than one letter")
			letter = 0
		}
		if listContains(foundLetters, letter) {
			fmt.Println("Letter was already found")
			letter = 0
		}
		if listContains(guessedLetters, letter) {
			fmt.Println("Letter was already guessed")
			letter = 0
		}
	}

	correct := false

	for i := 0; i < len(word); i++ {
		if letter == word[i] {
			correct = true
		}
	}

	if correct {
		foundLetters.PushFront(letter)
	} else {
		guessedLetters.PushFront(letter)
		lives -= 1
	}

	fmt.Print("\n")

	done := false
	if lives == 0 {
		fmt.Print(hangmen[len(hangmen)-1-lives])
		fmt.Print("\n")
		fmt.Println("You've Lost!")
		done = true
	}

	won := false
	if len(word) == foundLetters.Len() {
		won = true
		done = true
	}

	return lives, done, won
}

func listContains(checkList *list.List, element interface{}) bool {
	contains := false
	for listElement := checkList.Front(); listElement != nil; listElement = listElement.Next() {
		if element == listElement.Value {
			contains = true
		}
	}
	return contains
}

func main() {
	word := genWord()

	foundLetters := list.New()
	guessedLetters := list.New()
	lives := 6
	won := false
	for done := false; !done; {
		lives, done, won = gameloop(word, foundLetters, guessedLetters, lives)
	}

	if won {
		fmt.Print("\n")
		fmt.Println("You've Won!")
	}
}
