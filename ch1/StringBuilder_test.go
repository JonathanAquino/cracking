package ch1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringBuilder struct {
	words []string
}

func (sb *stringBuilder) append(s string) {
	sb.words = append(sb.words, s)
}

func (sb *stringBuilder) toString() string {
	return strings.Join(sb.words, "")
}

func TestStringBuilder(t *testing.T) {
	sb := stringBuilder{}
	sb.append("foo")
	sb.append("bar")
	sb.append("baz")
	assert.Equal(t, "foobarbaz", sb.toString())
}
