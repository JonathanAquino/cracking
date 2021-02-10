package ch8

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The classic Towers of Hanoi problem.

// Moves count disks from the stack named from to the stack named to.
func move(from, to, other string, count int, stacks map[string]*Stack, indent int) {
	println(spaces(indent) + fmt.Sprintf("move(%s, %s, %s, %d)", from, to, other, count))
	if count == 0 {
		return
	}
	if count == 1 {
		if stacks[to].isEmpty() || stacks[to].peek() > stacks[from].peek() {
			item := stacks[from].pop()
			stacks[to].push(item)
			printStacks(stacks, indent)
			return
		}
	}
	fromDisk := stacks[from].peek()
	numSmallerToDisks := numSmaller(stacks[to], fromDisk)
	move(from, other, to, 1, stacks, indent+1)
	move(to, from, other, numSmallerToDisks, stacks, indent+1)
	move(other, to, from, 1, stacks, indent+1)
	move(from, to, other, numSmallerToDisks, stacks, indent+1)
	move(from, to, other, count-1, stacks, indent+1)
}

// Returns the number of disks that are smaller than n.
func numSmaller(stack *Stack, n int) int {
	num := 0
	for i := stack.size() - 1; i >= 0; i-- {
		if stack.items[i] >= n {
			break
		}
		num++
	}
	return num
}

// Outputs the contents of the stacks named a, b, and c.
func printStacks(stacks map[string]*Stack, indent int) {
	printStack(stacks["a"], indent)
	printStack(stacks["b"], indent)
	printStack(stacks["c"], indent)
	println()
}

// Outputs the contents of the stack
func printStack(stack *Stack, indent int) {
	itemStrings := []string{}
	for _, item := range stack.items {
		itemStrings = append(itemStrings, strconv.Itoa(item))
	}
	println(spaces(indent) + ">" + strings.Join(itemStrings, ","))
}

// spaces returns several spaces depending on the given indentation level.
func spaces(indent int) string {
	result := ""
	for i := 0; i < indent; i++ {
		result += " "
	}
	return result
}

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	if !s.isEmpty() && s.peek() < item {
		panic(fmt.Sprintf("attempting to put %d on top of %d", item, s.peek()))
	}
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *Stack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	return s.size() == 0
}

func (s *Stack) size() int {
	return len(s.items)
}

func Test8Dot6Stack(t *testing.T) {
	s := Stack{}
	s.push(5)
	s.push(4)
	assert.Equal(t, 4, s.peek())
	assert.Equal(t, 4, s.pop())
	assert.Equal(t, 5, s.peek())
	assert.Equal(t, 1, s.size())
	assert.False(t, s.isEmpty())
	assert.Equal(t, 5, s.pop())
	assert.True(t, s.isEmpty())
}

func Test8Dot6NumSmaller(t *testing.T) {
	s := Stack{}
	s.push(5)
	s.push(4)
	assert.Equal(t, 2, numSmaller(&s, 6))
	assert.Equal(t, 1, numSmaller(&s, 5))
	assert.Equal(t, 0, numSmaller(&s, 4))
}

func Test8Dot6Move(t *testing.T) {
	stacks := map[string]*Stack{
		"a": &Stack{items: []int{8, 7, 6, 5, 4, 3, 2, 1}},
		"b": &Stack{items: []int{}},
		"c": &Stack{items: []int{}},
	}
	printStacks(stacks, 0)
	move("a", "c", "b", stacks["a"].size(), stacks, 0)
	expected := map[string]*Stack{
		"a": &Stack{items: []int{}},
		"b": &Stack{items: []int{}},
		"c": &Stack{items: []int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	assert.Equal(t, *expected["a"], *stacks["a"])
	assert.Equal(t, *expected["b"], *stacks["b"])
	assert.Equal(t, *expected["c"], *stacks["c"])
}
