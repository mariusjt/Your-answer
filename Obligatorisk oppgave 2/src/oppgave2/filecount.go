package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"strings"
	"os"
	"go/scanner"
)



func main() {
		content, err := ioutil.ReadFile("text.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File contents: %s", content)

	file, err := os.Open("text.txt") // For read access.
	if err != nil {
		log.Fatal(err)
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}
