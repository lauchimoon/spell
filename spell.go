package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	WORDS_FILE = "words.txt"
)

func readWords() ([]string, error) {
	var words []string
	f, err := os.Open(WORDS_FILE)
	if err != nil {
		return words, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() { // Read line by line
		words = append(words, strings.Trim(scanner.Text(), "\n"))
	}

	return words, nil
}

func spellcheck(wordList []string, text string) ([]string, int) {
	var mistakes []string
	wordMap := make(map[string]byte)
	exclude := []string{
		",", "",
		".", "",
		"!", "",
		"?", "",
		"¿", "",
		"¡", "",
		"\n", "",
		"\t", "",
	}

	// Remove all unwanted symbols (commas, periods...)
	replacer := strings.NewReplacer(exclude...)
	txtTrimmed := replacer.Replace(text)
	splitText := strings.Split(txtTrimmed, " ")

	// Transfer every element on the wordList slice
	// to the wordMap map
	for _, word := range wordList {
		wordMap[strings.ToLower(word)] = '0'
	}

	// Check every word in the wordlist.
	// If one of the words is not in the text,
	// add it to the list
	for _, word := range splitText {
		wLower := strings.ToLower(word)
		if _, ok := wordMap[wLower]; !ok {
			mistakes = append(mistakes, word)
		}
	}

	nmistakes := len(mistakes)

	return mistakes, nmistakes
}

func readFile(path string) (string, error) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func die(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func usage() {
	fmt.Println("spell: Rudimentary spell checker")
	fmt.Println("usage: spell <text> OR spell -f <file_to_read_from>")
}

func main() {
	words, err := readWords()
	if err != nil {
		die(err)
	}

	args := os.Args
	var mistakes []string
	var nmistakes int

	if len(args) <= 1 {
		usage()
		die("spell: missing text to spellcheck")
		os.Exit(1)
	} else if len(args) == 2 {
		mistakes, nmistakes = spellcheck(words, args[1])
	} else if len(args) == 3 {
		if args[1] != "-f" {
			die("spell: need -f parameter.")
		}

		text, err := readFile(args[2])
		if err != nil {
			die(err)
		}
		mistakes, nmistakes = spellcheck(words, text)
	}

	if nmistakes > 0 {
		fmt.Printf("%d mistake(s):\n", nmistakes)

		for _, mistake := range mistakes {
			fmt.Printf("  %s\n", mistake)
		}
	} else {
		fmt.Println("No mistakes")
	}
}
