package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"io/ioutil"
)

func main() {
	lineCounter(getFile())
	fileReader(getFile())
}

func getFile() string {
	input := os.Args
	if len(input) <1 {
		fmt.Println("Error: No filename entered")
	}
	fileName := input[1]
	filePath := "../files/" + fileName

	return filePath

}

func lineCounter(filePath string){
	file, err := os.Open(filePath)
	if err != nil {
		 errHandler()
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	fmt.Println("Number of lines in file:", lineCount)
}
var m map[rune] int

func fileReader(filePath string){
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not open file")
		log.Fatal(err)
	}
	buf := bufio.NewReader(file)
	for true {
		x, _, err := buf.ReadRune()
		if err != nil {
			break
		}
		m[x]++
	}



}

func errHandler(){
	log.Fatal()
}