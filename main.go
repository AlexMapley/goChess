package main

import (
	"fmt"
)

func main(){

	board := CreateBoard()
	fmt.Println("Chess")

	console := Console{}
	
	console.DisplayConsole()
}