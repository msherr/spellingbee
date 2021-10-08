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
		foundCenter = foundCenter || (letter == centerLetter)
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
	*puzzle = strings.ToLower(*puzzle)
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

	blueBold := color.New(color.FgBlue, color.Bold)

	// figure out the list of words, keeping the solutions separate
	winners := make([]string, 0)
	solutions := make([]string, 0)
	for _, word := range words {
		word = strings.Trim(word, "\n\r")
		if isGood(*puzzle, centerLetter, word) {
			if isSolution(*puzzle, centerLetter, word) {
				solutions = append(solutions, word)
			} else {
				winners = append(winners, word)
			}
		}
	}
	// sort by length
	sort.Slice(winners, func(i, j int) bool {
		return len(winners[i]) < len(winners[j])
	})
	for _, winner := range winners {
		fmt.Println(winner)
	}
	for _, solution := range solutions {
		blueBold.Println(solution)
	}

}
