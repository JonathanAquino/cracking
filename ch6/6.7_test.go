package ch5

import (
	"math/rand"
	"testing"
	"time"
)

// Simulates number of boys and girls given this rule: family keeps having
// children until they have a girl.
func apocalypse() (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	boys := 0
	girls := 0
	for families := 0; families < 100000000; families++ {
		for true {
			if rand.Int31n(2) == 0 {
				boys++
			} else {
				girls++
				break
			}
		}
	}
	return boys, girls
}

func TestApocalypse(t *testing.T) {
	boys, girls := apocalypse()
	println("foo")
	t.Log("foo")
	t.Logf("Boys: %d\n", boys)
	t.Logf("Girls: %d\n", girls)
	t.Logf("Girl-Boy Ratio: %f\n", float32(girls)/float32(boys))
	t.Fail() // So Logf statements are displayed in VS Code
}
