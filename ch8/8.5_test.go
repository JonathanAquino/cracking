package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// multiply multiplies the two numbers without using *, but only
// +, -, and bit shifting.
func multiply(a, b int) int {
	var lo, hi int
	if a < b {
		lo = a
		hi = b
	} else {
		lo = b
		hi = a
	}
	sum := 0
	for i := 0; i < lo; i++ {
		sum += hi
	}
	return sum
}

// multiplyV2 is like multiply but with fewer operations. This runs in
// O(log N) time unlike multiply which runs in O(N) time.
func multiplyV2(a, b int) int {
	var lo, hi int
	if a < b {
		lo = a
		hi = b
	} else {
		lo = b
		hi = a
	}
	if lo == 0 {
		return lo
	}
	if lo == 1 {
		return hi
	}
	// Shifting left multiplies by two. We need to find out how far
	// we can shift hi left by first finding out how far we can shift lo right.
	// Once we know that, we can shift hi left by that amount. And then we
	// can call multiplyV2 again on the remainder.
	shift := 1
	for {
		// Keep integer-dividing lo by 2 until we can divide no more.
		if lo>>shift <= 1 {
			break
		}
		shift++
	}
	// multiplier is the amount we multiply hi by when we shift hi left by shift.
	multiplier := 1 << shift
	// After we shift hi by shift, our remaining work is to multiply hi
	// by (lo - multiplier).
	return (hi << shift) + multiplyV2(hi, lo-multiplier)
}

func Test8Dot5(t *testing.T) {
	assert.Equal(t, 8*5, multiply(8, 5))
	assert.Equal(t, 1000*1000, multiply(1000, 1000))
	assert.Equal(t, 13*7, multiply(13, 7))
}

func Test8Dot5V2(t *testing.T) {
	assert.Equal(t, 8*5, multiplyV2(8, 5))
	assert.Equal(t, 1000*1000, multiplyV2(1000, 1000))
	assert.Equal(t, 13*7, multiplyV2(13, 7))
}
