package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bufio"
)



func main() {
		content, err := ioutil.ReadFile("text.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File contents: %s", content)

	file, err := os.Open("text.txt") // For read access.
	buf := bufio.NewReader(file)

	var m map[rune]int

	for true {
		x, _, err := buf.ReadRune()
		if err != nil {
			break
		}
		m[x]++;
	}

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	if err != nil {
		log.Fatal(err)
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
			bufio.ScanRunes(file)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}
