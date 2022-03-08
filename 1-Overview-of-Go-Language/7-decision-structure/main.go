package main

import (
	"fmt"
	"log"
)

func main() {
	var myMarks int
	fmt.Print("Enter Your Marks : ")
	fmt.Scanf("%d", &myMarks)
	fmt.Println("Your Enter Marks is : ", myMarks)
	if myMarks > 100 || myMarks < 0 {
		log.Println("You Enter Wrong Marks !!!")
	}
	if myMarks > 79 && myMarks <= 100 {
		log.Println("A+")
	}

}
