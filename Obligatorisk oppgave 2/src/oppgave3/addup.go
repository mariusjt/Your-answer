package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Legg sammen to tall fra terminal og summer i en annen funksjon")
	//fmt.Println(addUp(10, 10))
	c := make(chan int)
	input1 := os.Args[1]
	input2 := os.Args[2]
	tall,_ := strconv.Atoi(input1)
	tall2,_ := strconv.Atoi(input2)

	go addUp(tall, tall2, c)
	sum := <-c
	fmt.Println("Tall mottatt:")
	fmt.Println(input1, input2)
	fmt.Println("Sum:")
	fmt.Println(sum)
}

func addUp(tall int, tall2 int, c chan int){
	sum := tall + tall2
	c <- sum
}
