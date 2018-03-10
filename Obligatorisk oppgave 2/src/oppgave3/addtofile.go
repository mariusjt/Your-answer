package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func main() {
	CreateFile()
}

	func CreateFile(){

		fmt.Println("Legg sammen to tall fra terminal og summer i en annen funksjon")
		//fmt.Println(addUp(10, 10))

		input1 := os.Args[1]
		input2 := os.Args[2]
		tall, _ := strconv.Atoi(input1)
		tall2, _ := strconv.Atoi(input2)

		file, err := os.Create("nummer.txt") //
		if err != nil {
			log.Fatalf("feil skrive fil: %er", err)
		}

		file.WriteString(string(tall))
		file.WriteString(string(tall2))

		defer file.Close() // lukker filen.

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}

	}


