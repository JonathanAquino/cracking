package ch8

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/adam-lavrik/go-imath/ix"
)

// StackOfBoxes returns the tallest height that can be made from a bunch
// of boxes, given that a box can only rest on top of another box if the
// first one has a smaller width, depth, and height.
func StackOfBoxes(boxes []Stackable) int {
	// First we sort the boxes by volume. Then we know that for a box to
	// be on top of another box, the first must appear after the second
	// in this sorted list. That said, just because a box appears after
	// another box on this list doesn't mean that it is allowed to be on top.
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].CanGoBelow(boxes[j])
	})
	// Set the first Stackable to the Ground.
	boxes = append([]Stackable{Ground{}}, boxes...)
	maxHeight, _ := stackOfBoxesHelper(0, boxes)
	return maxHeight
}

// stackOfBoxesHelper considers the ith box in the list of boxes sorted by
// decreasing volume. It returns the tallest height that can be constructed
// from the remaining boxes. If there are any misfits that cannot be placed on
// this box or one of the ones on top of it, they are returned.
func stackOfBoxesHelper(i int, boxes []Stackable) (int, []int) {
	if i == len(boxes)-1 {
		return 0, []int{}
	}
	// We will look at each box in the sorted list. If the next box fits
	// on top, we will add its height to the total. If not, we will return
	// the misfits in an attempt to branch them from an earlier (bigger) box.
	maxHeight := 0
	misfits := []int{}
	queue := []int{i + 1}
	for len(queue) > 0 {
		next := queue[0]
		nextBox := boxes[next]
		queue = queue[1:]
		if !boxes[i].CanGoBelow(nextBox) {
			misfits = append(misfits, next)
			continue
		}
		height := nextBox.GetHeight()
		subheight, submisfits := stackOfBoxesHelper(next, boxes)
		maxHeight = ix.Max(maxHeight, height+subheight)
		for _, submisfit := range submisfits {
			queue = append(queue, submisfit)
		}
	}
	return maxHeight, misfits
}

// Stackable is something that can be stacked.
type Stackable interface {
	CanGoBelow(other Stackable) bool
	GetLength() int
	GetWidth() int
	GetHeight() int
}

// A Box is something with length, width, and height.
type Box struct {
	length, width, height int
}

func (b Box) GetLength() int {
	return b.length
}

func (b Box) GetWidth() int {
	return b.width
}

func (b Box) GetHeight() int {
	return b.height
}

func (b Box) CanGoBelow(other Stackable) bool {
	return b.length > other.GetLength() && b.height > other.GetHeight() && b.width > other.GetWidth()
}

// Ground represents the floor.
type Ground struct{}

func (g Ground) GetLength() int {
	return 0
}

func (g Ground) GetWidth() int {
	return 0
}

func (g Ground) GetHeight() int {
	return 0
}

func (g Ground) CanGoBelow(other Stackable) bool {
	return true
}

func TestStackOfBoxes(t *testing.T) {
	assert.Equal(t, 15, StackOfBoxes([]Stackable{
		Box{length: 1, width: 1, height: 1},
		Box{length: 3, width: 3, height: 3},
		Box{length: 5, width: 5, height: 5},
		Box{length: 4, width: 4, height: 4},
		Box{length: 2, width: 2, height: 2},
	}))
}

func TestStackOfBoxes2(t *testing.T) {
	// Here we consider two possible branches: the long branch and the wide branch
	assert.Equal(t, 147, StackOfBoxes([]Stackable{
		Box{length: 50, width: 50, height: 50},
		Box{length: 10, width: 49, height: 49},
		Box{length: 9, width: 48, height: 48},
		Box{length: 49, width: 10, height: 47},
		Box{length: 48, width: 9, height: 46},
	}))
}
