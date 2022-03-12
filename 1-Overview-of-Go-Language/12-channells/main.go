package main

import (
	"log"

	"github.com/Riyaz-khan-shuvo/go-with-project/1-Overview-of-Go-Language/12-channells/helpers"
)

const numPoll int = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumbers(numPoll)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)

}
