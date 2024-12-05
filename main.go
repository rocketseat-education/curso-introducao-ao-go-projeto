package main

import "fmt"

type Question struct {
	Text string
	Options []string
	Answer  int
}

type GameState struct {
	Name string
	Points string
	Questions []Question
}

func main() {
	fmt.Println("Ola")
}
