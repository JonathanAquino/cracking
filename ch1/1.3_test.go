package ch1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func urlify(b *[]byte, length int) {
	numSpaces := 0
	for i := 0; i < length; i++ {
		c := (*b)[i]
		if c == ' ' {
			numSpaces++
		}
	}
	finalLength := length + (numSpaces * 2)
	j := length
	for i := finalLength - 1; i >= 0; i-- {
		j--
		if (*b)[j] == ' ' {
			(*b)[i] = '0'
			i--
			(*b)[i] = '2'
			i--
			(*b)[i] = '%'
		} else {
			(*b)[i] = (*b)[j]
		}
	}
	println(numSpaces)
}

func Test1Dot3(t *testing.T) {
	b := []byte("Mr John Smith    ")
	urlify(&b, 13)
	assert.Equal(t, []byte("Mr%20John%20Smith"), b)
}
