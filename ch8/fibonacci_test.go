package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic implementation
func fib1(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fib1(i-1) + fib1(i-2)
}

// Top-down memoization
func fib2(i int) int {
	memo := make([]int, i+1)
	return fib2Helper(i, memo)
}

func fib2Helper(i int, memo []int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	result := fib2Helper(i-1, memo) + fib2Helper(i-2, memo)
	memo[i] = result
	return result
}

// Bottom-up memoization
func fib3(i int) int {
	memo := make([]int, i+1)
	memo[0] = 0
	memo[1] = 1
	for j := 2; j <= i; j++ {
		memo[j] = memo[j-1] + memo[j-2]
	}
	return memo[i]
}

// Bottom-up memoization with optimization
func fib4(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	a := 0
	b := 1
	var c int
	for j := 2; j <= i; j++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func TestFib1(t *testing.T) {
	assert.Equal(t, 5, fib1(5))
	assert.Equal(t, 55, fib1(10))
	assert.Equal(t, 610, fib1(15))
}

func TestFib2(t *testing.T) {
	assert.Equal(t, 5, fib2(5))
	assert.Equal(t, 55, fib2(10))
	assert.Equal(t, 610, fib2(15))
}

func TestFib3(t *testing.T) {
	assert.Equal(t, 5, fib3(5))
	assert.Equal(t, 55, fib3(10))
	assert.Equal(t, 610, fib3(15))
}

func TestFib4(t *testing.T) {
	assert.Equal(t, 5, fib4(5))
	assert.Equal(t, 55, fib4(10))
	assert.Equal(t, 610, fib4(15))
}
