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
func move(from, to, other string, count int, stacks map[string]*Stack) {
	println(fmt.Sprintf("move(%s, %s, %s, %d)", from, to, other, count))
	if count == 0 {
		return
	}
	if count == 1 {
		if !stacks[to].isEmpty() && stacks[to].peek() < stacks[from].peek() {
			move(from, other, to, 1, stacks)
			move(to, from, other, 1, stacks)
			move(other, to, from, 1, stacks)
			move(from, to, other, 1, stacks)
			return
		}
		item := stacks[from].pop()
		stacks[to].push(item)
		printStacks(stacks)
		return
	}
	move(from, other, to, 1, stacks)
	toDisksToMove := stacks[to].size()
	if !stacks[to].isEmpty() && stacks[to].peek() > stacks[from].peek() {
		toDisksToMove = 0
	}
	move(to, from, other, toDisksToMove, stacks)
	move(other, to, from, 1, stacks)
	move(from, to, other, toDisksToMove, stacks)
	move(from, to, other, stacks[from].size(), stacks)
}

// Outputs the contents of the stacks named a, b, and c.
func printStacks(stacks map[string]*Stack) {
	printStack(stacks["a"])
	printStack(stacks["b"])
	printStack(stacks["c"])
	println()
}

// Outputs the contents of the stack
func printStack(stack *Stack) {
	itemStrings := []string{}
	for _, item := range stack.items {
		itemStrings = append(itemStrings, strconv.Itoa(item))
	}
	println(">" + strings.Join(itemStrings, ","))
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

func Test8Dot6Move(t *testing.T) {
	stacks := map[string]*Stack{
		"a": &Stack{items: []int{8, 7, 6, 5, 4, 3, 2, 1}},
		"b": &Stack{items: []int{}},
		"c": &Stack{items: []int{}},
	}
	printStacks(stacks)
	move("a", "c", "b", stacks["a"].size(), stacks)
	expected := map[string]*Stack{
		"a": &Stack{items: []int{}},
		"b": &Stack{items: []int{}},
		"c": &Stack{items: []int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	assert.Equal(t, *expected["a"], *stacks["a"])
	assert.Equal(t, *expected["b"], *stacks["b"])
	assert.Equal(t, *expected["c"], *stacks["c"])
}
