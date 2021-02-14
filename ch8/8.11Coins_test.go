package ch8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Coins returns all combinations of the given coins to produce the target amount.
func Coins(amount int, denominations []int) [][]int {
	if len(denominations) == 1 {
		return [][]int{[]int{amount}}
	}
	if len(denominations) == 0 {
		return [][]int{[]int{}}
	}
	if amount == 0 {
		combination := make([]int, len(denominations))
		return [][]int{combination}
	}
	maxNumber := amount / denominations[0]
	combinations := [][]int{}
	for i := 0; i <= maxNumber; i++ {
		for _, subcombination := range Coins(amount-i*denominations[0], denominations[1:]) {
			combination := append([]int{i}, subcombination...)
			combinations = append(combinations, combination)
		}
	}
	return combinations
}

func TestCoins(t *testing.T) {
	combinations := Coins(30, []int{25, 10, 5, 1})
	formattedCombinations := []string{}
	for _, combination := range combinations {
		formattedCombinations = append(formattedCombinations,
			fmt.Sprintf("%dq%dd%dn%dp", combination[0], combination[1], combination[2], combination[3]))
	}
	expected := []string{
		"0q0d0n30p",
		"0q0d1n25p",
		"0q0d2n20p",
		"0q0d3n15p",
		"0q0d4n10p",
		"0q0d5n5p",
		"0q0d6n0p",
		"0q1d0n20p",
		"0q1d1n15p",
		"0q1d2n10p",
		"0q1d3n5p",
		"0q1d4n0p",
		"0q2d0n10p",
		"0q2d1n5p",
		"0q2d2n0p",
		"0q3d0n0p",
		"1q0d0n5p",
		"1q0d1n0p",
	}
	assert.Equal(t, expected, formattedCombinations)
}
