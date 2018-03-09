package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"io/ioutil"
	"strings"
)

func main() {
	lineCounter(getFile())
	fileReader(getFile())
	Counter()
}

func getFile() string {
	input := os.Args
	if len(input) <1 {
		fmt.Println("Error: No filename entered")
	}

	filename = input[1]
	filePath := "../files/" + filename

	return filePath
}

var filename string

func lineCounter(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		 errHandler()
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	fmt.Printf("Information about '%s':\n", filename)
	fmt.Println("Number of lines in file:", lineCount)
}
var Map = make(map[rune]int)

func fileReader(filePath string){
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could not open file")
		errHandler()
	}
	fileString := string(file)
	fileSplit := []string(strings.Split(fileString, ""))

	for i := 0; i < len(fileSplit); i++ {
		addToMap(fileSplit[i])
	}
}

var regLetters [] string //registrerte bokstaver

func addToMap(letter string) {
	registered := false
	for i := 0; i < len(regLetters); i++ {
		if letter == regLetters[i] {
			registered = true
		}
	}

	stringVer := []rune(letter)
	intVer := int(stringVer[0])

	if registered == true {
		Map[rune(intVer)]++
	} else {
		Map[rune(intVer)] = 1
		regLetters = append(regLetters, letter)
	}
}

func Counter(){
	fmt.Println("Most common runes:")

	for i := 1; i <= 5; i++ {
		number := i
		highest := 0
		mostUsed := ""

		for i := 0; i < len(regLetters); i++ {
			runeVer := []rune(regLetters[i])
			counts := Map[runeVer[0]]

			if counts > highest {
				highest = counts
				mostUsed = regLetters[i]
			}
		}

		topRune := []rune(mostUsed)
		fmt.Printf("%d, Rune: '%s', Counts: %d\n", number, mostUsed, highest)
		delete(Map, topRune[0])
	}
}


func errHandler(){
	log.Fatal()
}