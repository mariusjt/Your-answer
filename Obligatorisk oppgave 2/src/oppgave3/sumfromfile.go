package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	func ReadFile() {
		data, err := ioutil.ReadFile("test.txt")
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
		fmt.Printf("\nLength: %d bytes", len(data))
		fmt.Printf("\nData: %s", data)
		fmt.Printf("\nError: %v", err)
	}

	func main() {
		fmt.Printf("########Create a file and Write the content #########\n")
		CreateFile()

		fmt.Printf("\n\n########Read file #########\n")
		ReadFile()
	}
}
