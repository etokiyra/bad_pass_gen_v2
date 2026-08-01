package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main(){
	data, err := os.ReadFile("rockyou.txt")
	if err != nil {
		panic(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	fmt.Println("Loaded", len(lines), "lines")

	index := rand.Intn(len(lines))
	password := lines[index]
	fmt.Println("Your super duper srong password is:", password)
}