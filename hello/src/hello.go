package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Who're you?")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Well howdy there %s! 🤠", name)
}
