package main

import (
	"fmt"
	handler "main/Handler"
)

func main() {
	fmt.Println("Welcome To Master Chief Golang")

	err := handler.Handle()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Thanks for visiting")
}
