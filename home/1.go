package main

import (
	"fmt"
	"math/rand"

)


func generateRandomNumber(ch chan<- int) {
	randomNumber := rand.Intn(100)
	ch <- randomNumber
}

func main() {
	randomNumberChannel := make(chan int)
	go generateRandomNumber(randomNumberChannel)

	randomNumber := <-randomNumberChannel

	fmt.Println("Random number:", randomNumber)
}
