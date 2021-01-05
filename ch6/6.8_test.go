package ch5

import (
	"fmt"
	"testing"

	"github.com/adam-lavrik/go-imath/ix"
)

// eggDrop does an egg drop simulation. We are trying to find at which floor of
// a 100 storey building the eggs crack. We are given two eggs to use. This
// function tries to minimize the number of egg drops required by dropping egg 1
// from the initialFloor, then initialFloor-1 more floors, then initialFloor-2
// more floors, etc. This function returns the maximum number of drops required in
// the worst case. By trying this function with different initialFloors, you
// can figure out which initialFloor to use to minimize the worst case.
func eggDrop(initialFloor int) int {
	prevFloor := 0
	floor := initialFloor
	egg1Drops := 0
	maxEggDrops := 0
	for floor < 100 {
		egg1Drops++
		egg2Drops := 0
		for f := prevFloor + 1; f <= floor-1; f++ {
			egg2Drops++
		}
		eggDrops := egg1Drops + egg2Drops
		maxEggDrops = ix.Max(maxEggDrops, eggDrops)
		// Reduce the number of floors by 1 to try to keep maxEggDrops the same.
		// But the delta has to be at least 2.
		delta := ix.Max(2, (floor-prevFloor)-1)
		prevFloor = floor
		floor = floor + delta
	}
	return maxEggDrops
}

func TestEggDrop(t *testing.T) {
	for i := 1; i < 100; i++ {
		maxEggDrops := eggDrop(i)
		fmt.Printf("initialFloor=%d, maxEggDrops=%d\n", i, maxEggDrops)
	}
	// initialFloor=14 minimimzes the worst case egg drops.
	// Fail so we see the printf statements
	t.Fail()
}
