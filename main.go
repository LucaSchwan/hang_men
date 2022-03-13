package main

import (
	"bufio"
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

func gameloop(word string, foundLetter []byte, absentLetter []byte, lives int) ([]byte, []byte, int) {
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

	return []byte{}, []byte{}, 0
}

func main() {
	genWord()

}
