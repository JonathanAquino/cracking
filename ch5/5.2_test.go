package ch5

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// floatToBinaryString converts the given float to a binary representation as
// a string. For example, 0.71875 becomes "0.10111".
func floatToBinaryString(x float64) (string, error) {
	n := 0
	sum := 0.
	result := "0."
	for true {
		if sum == x {
			return result, nil
		}
		n++
		if n > 32 {
			return "", fmt.Errorf("%g cannot be represented in 32 characters or less", x)
		}
		if sum+math.Pow(2., float64(-n)) <= x {
			sum += math.Pow(2., float64(-n))
			result += "1"
		} else {
			result += "0"
		}
	}
	return result, nil
}

func TestFloatToBinaryString(t *testing.T) {
	result, _ := floatToBinaryString(0.71875)
	assert.Equal(t, "0.10111", result)
}

func TestFloatToBinaryString2(t *testing.T) {
	_, err := floatToBinaryString(0.72)
	assert.NotNil(t, err)
}
