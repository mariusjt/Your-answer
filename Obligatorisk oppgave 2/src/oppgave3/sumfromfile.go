package main

import (
	"fmt"
	"os"
	"io"
	"log"
)

func main() {
	ReadFile()

}
func ReadFile() {

	file, err := os.Open("nummer.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)

		file.WriteString(string(sum))

			defer file.Close() // lukker filen.

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
		}
	}
}
