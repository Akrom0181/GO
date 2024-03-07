package main

import (
    "fmt"
    "math/rand"
)

func sendRandomNumbers(numCh chan int) {
    number := rand.Intn(10) 
    fmt.Printf("Sending %d random numbers\n", number)
    for i := 0; i < number; i++ {
        randomNumber := rand.Intn(100) 
        numCh <- randomNumber
    }
    close(numCh)
}

func main() {
    numCh := make(chan int)
    go sendRandomNumbers(numCh)
    for num := range numCh {
        fmt.Println("Received random number:", num)
    }
}
