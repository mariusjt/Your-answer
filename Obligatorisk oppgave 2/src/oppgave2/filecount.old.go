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

	file, err := os.Open("../files/text.txt") // For read access.
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
}
