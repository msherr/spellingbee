/**

A solver of NYTimes Spelling Bees

Micah Sherr <msherr@cs.georgetown.edu>
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/fatih/color"
)

const DEFAULT_DICTFILE string = "/usr/share/dict/words"
const MIN_LENGTH int = 4

func isGood(letterMap map[string]bool, centerLetter, candidateWord string) (isValid, usesAllLetters bool) {
	foundLetters := make(map[string]bool)
	if len(candidateWord) < MIN_LENGTH {
		return false, false
	}
	foundCenter := false
	for _, letter := range strings.Split(candidateWord, "") {
		foundCenter = foundCenter || (letter == centerLetter)
		if _, ok := letterMap[letter]; !ok { // contains letter not in puzzle
			return false, false
		}
		foundLetters[letter] = true
	}
	return foundCenter, len(foundLetters) == len(letterMap)
}

func makeMap(letters string) map[string]bool {
	letterMap := make(map[string]bool)
	for _, letter := range strings.Split(letters, "") {
		letterMap[letter] = true
	}
	return letterMap
}

func main() {
	parser := argparse.NewParser("spellingbee", "solves NYTimes spelling bees")
	dictfile := parser.String("d", "dict", &argparse.Options{
		Default: DEFAULT_DICTFILE,
		Help:    "path to dictionary file",
	})
	puzzle := parser.String("p", "puzzle", &argparse.Options{
		Required: true,
		Help:     "puzzle letters (put center letter first)",
	})
	quietmode := parser.Flag("q", "quiet", &argparse.Options{
		Help: "quiet mode; only output words",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		panic("invalid usage")
	}

	if !*quietmode {
		fmt.Print("Spelling bee cheater!  Shame on you.\n\n\n")
	}

	*puzzle = strings.ToLower(*puzzle)
	centerLetter := strings.Split(*puzzle, "")[0]
	letterMap := makeMap(*puzzle)

	// read dictionary; consider only lower-case words
	dat, err := ioutil.ReadFile(*dictfile)
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

	// figure out the list of words, keeping the solutions separate
	winners := make([]string, 0)
	solutions := make([]string, 0)
	for _, word := range words {
		isValid, usesAllLetters := isGood(letterMap, centerLetter, word)
		if isValid && !usesAllLetters {
			winners = append(winners, word)
		} else if isValid {
			solutions = append(solutions, word)
		}
	}

	// sort by length
	sort.Slice(winners, func(i, j int) bool {
		return len(winners[i]) < len(winners[j])
	})
	for _, winner := range winners {
		fmt.Println(winner)
	}

	blueBold := color.New(color.FgBlue, color.Bold)
	for _, solution := range solutions {
		blueBold.Println(solution)
	}

	if !*quietmode {
		fmt.Printf("\n\nFound %d words.\n", len(winners)+len(solutions))
	}

}
