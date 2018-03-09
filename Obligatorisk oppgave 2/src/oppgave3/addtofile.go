package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func main() {
	fmt.Println("Legg sammen to tall fra terminal og summer i en annen funksjon")
	//fmt.Println(addUp(10, 10))

	input1 := os.Args[1]
	input2 := os.Args[2]
	tall, _ := strconv.Atoi(input1)
	tall2, _ := strconv.Atoi(input2)
}

	func CreateFile(){

		file, err := os.Create("tall.txt") //
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close() // Make sure to close the file when you're done

		len, err := file.WriteAt().

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		fmt.Printf("\nLength: %d bytes", len)
		fmt.Printf("\nFile Name: %s", file.Name())
	}

