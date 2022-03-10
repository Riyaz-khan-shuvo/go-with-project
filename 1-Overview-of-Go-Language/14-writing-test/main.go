package main

import (
	"errors"
	"log"
)

func main() {

	result, err := Divide(100.00, 0)
	if err != nil {
		log.Println(err)
	}
	log.Println("Result of Division is : ", result)
}

func Divide(x, y float32) (float32, error) {

	var result float32

	if y == 0 {
		return result, errors.New("can't divide by 0")
	}

	result = x / y
	return result, nil

}
