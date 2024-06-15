package main

import (
	"fmt"

	"github.com/zavexeon/dice/tokenizer"
)

func main() {
	result, _ := tokenizer.Tokenize("(5. x 2.7) / 2 ^ .5")

	fmt.Println(result)
}
