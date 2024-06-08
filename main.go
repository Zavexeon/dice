package main

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strconv"
	"strings"
)

func main() {

}

func rollDice(sides int, count int) int {
	total := 0
	for i := 0; i < count; i++ {
		total += rand.IntN(sides) + 1
	}
	return total
}

// * standardizes dice expressions by making them all fixed length of 3, ex: d6 -> 1d6
func normalizeDiceExpr(expr string) (string, error) {
	if match, _ := regexp.MatchString(`([1-9]\d*)d([1-9]]\d*)`, expr); match {
		return "", fmt.Errorf("invalid dice expression: %s", expr)
	}

	splitExpr := strings.Split(expr, "d")

	if splitExpr[0] == "" {
		splitExpr[0] = "1"
	}

	return strings.Join(splitExpr, "d"), nil
}

func evalDiceExpr(expr string) (int, error) {
	normalExpr, err := normalizeDiceExpr(expr)

	if err != nil {
		return 0, err
	}

	splitExpr := strings.Split(normalExpr, "d")
	sides, _ := strconv.Atoi(splitExpr[1])
	times := 1

	if splitExpr[0] != "" {
		sides, _ = strconv.Atoi(splitExpr[0])
	}

	return rollDice(sides, times), nil
}
