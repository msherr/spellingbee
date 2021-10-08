/**

A solver of NYTimes Spelling Bees

Micah Sherr <msherr@cs.georgetown.edu>
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/fatih/color"
)

const DEFAULT_DICTFILE string = "/usr/share/dict/words"
const MIN_LENGTH int = 4

func isGood(puzzleLetters, centerLetter, candidateWord string) bool {
	if len(candidateWord) < MIN_LENGTH {
		return false
	}
	foundCenter := false
	for _, letter := range strings.Split(candidateWord, "") {
		if letter == centerLetter {
			foundCenter = true
		}
		if strings.Index(puzzleLetters, letter) == -1 {
			return false
		}
	}
	return foundCenter
}

func isSolution(puzzleLetters, centerLetter, candidateWord string) bool {
	if !isGood(puzzleLetters, centerLetter, candidateWord) {
		return false
	}
	for _, letter := range strings.Split(puzzleLetters, "") {
		if strings.Index(candidateWord, letter) == -1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Spelling bee cheater!  Shame on you.")

	parser := argparse.NewParser("spellingbee-cheater", "solves NYTimes spelling bees")
	dictfile := parser.String("d", "dict", &argparse.Options{
		Default: DEFAULT_DICTFILE,
		Help:    "path to dictionary file",
	})
	puzzle := parser.String("p", "puzzle", &argparse.Options{
		Required: true,
		Help:     "puzzle letters (put center letter first)",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		panic("invalid usage")
	}
	centerLetter := strings.Split(*puzzle, "")[0]

	dat, err := os.ReadFile(*dictfile)
	if err != nil {
		panic("cannot open dictionary file")
	}
	words := make([]string, 0)
	word_candidates := strings.Split(string(dat), "\n")
	for _, candidate := range word_candidates {
		if candidate == strings.ToLower(candidate) {
			words = append(words, candidate)
		}
	}

	winners := make([]string, 0)
	for _, word := range words {
		word = strings.Trim(word, "\n\r")
		if isGood(*puzzle, centerLetter, word) {
			winners = append(winners, word)
		}
	}
	sort.Slice(winners, func(i, j int) bool {
		return len(winners[i]) < len(winners[j])
	})

	blueBold := color.New(color.FgBlue, color.Bold)
	for _, winner := range winners {
		if isSolution(*puzzle, centerLetter, winner) {
			blueBold.Println(winner)
		} else {
			fmt.Println(winner)
		}
	}

}
