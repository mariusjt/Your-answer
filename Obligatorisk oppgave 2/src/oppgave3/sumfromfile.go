package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	ReadFile()

}

func ReadFile() {
	data, err := ioutil.ReadFile("nummer.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
}

func addup(tall int, tall2 int) {
	sum := tall + tall2

}

