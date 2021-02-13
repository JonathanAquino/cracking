package ch8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// parens returns all valid combinations of pairs of parentheses.
func parens(n int) []string {
	return parensHelper(n, n, 0)
}

func parensHelper(remainingOpenParens, remainingClosedParens, currentlyOpenParens int) []string {
	if remainingOpenParens == 0 && remainingClosedParens == 0 {
		return []string{""}
	}
	results := []string{}
	if remainingOpenParens > 0 {
		for _, subresult := range parensHelper(remainingOpenParens-1, remainingClosedParens, currentlyOpenParens+1) {
			results = append(results, "("+subresult)
		}
	}
	if remainingClosedParens > 0 && currentlyOpenParens > 0 {
		for _, subresult := range parensHelper(remainingOpenParens, remainingClosedParens-1, currentlyOpenParens-1) {
			results = append(results, ")"+subresult)
		}
	}
	return results
}

func TestParens(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	assert.Equal(t, expected, parens(3))
}
