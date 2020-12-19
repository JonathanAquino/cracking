package ch3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// An AnimalShelter takes in dogs and cats and allows a customer to choose
// the oldest dog, the oldest cat, or the oldest animal (either dog or cat).
type AnimalShelter struct {
	dogQueue  *SinglyLinkedList
	catQueue  *SinglyLinkedList
	lastIndex int
}

func NewAnimalShelter() *AnimalShelter {
	return &AnimalShelter{dogQueue: &SinglyLinkedList{}, catQueue: &SinglyLinkedList{}}
}

type Animal struct {
	name    string
	species string
	index   int
}

func (s *AnimalShelter) enqueue(species string, name string) {
	s.lastIndex++
	animal := Animal{name: name, species: species, index: s.lastIndex}
	switch species {
	case "dog":
		s.dogQueue.Add(animal)
	case "cat":
		s.catQueue.Add(animal)
	}
}

func (s *AnimalShelter) dequeueDog() (string, error) {
	if s.dogQueue.IsEmpty() {
		return "", errors.New("dog queue is empty")
	}
	animal, err := s.dogQueue.Shift()
	if err != nil {
		return "", err
	}
	return animal.(Animal).name, nil
}

func (s *AnimalShelter) dequeueCat() (string, error) {
	if s.catQueue.IsEmpty() {
		return "", errors.New("cat queue is empty")
	}
	animal, err := s.catQueue.Shift()
	if err != nil {
		return "", err
	}
	return animal.(Animal).name, nil
}

func (s *AnimalShelter) dequeueAny() (string, error) {
	if s.dogQueue.IsEmpty() && s.catQueue.IsEmpty() {
		return "", errors.New("dog and cat queues are empty")
	}
	var animal interface{}
	var err error
	if s.dogQueue.IsEmpty() {
		animal, err = s.catQueue.Shift()
	} else if s.catQueue.IsEmpty() {
		animal, err = s.dogQueue.Shift()
	} else if s.catQueue.head.data.(Animal).index < s.dogQueue.head.data.(Animal).index {
		animal, err = s.catQueue.Shift()
	} else {
		animal, err = s.dogQueue.Shift()
	}
	if err != nil {
		return "", err
	}
	return animal.(Animal).name, nil
}

func Test3Dot6(t *testing.T) {
	s := NewAnimalShelter()
	s.enqueue("cat", "cat3")
	s.enqueue("dog", "dog5")
	s.enqueue("cat", "cat2")
	s.enqueue("dog", "dog4")
	s.enqueue("dog", "dog3")
	s.enqueue("cat", "cat1")
	s.enqueue("dog", "dog2")
	s.enqueue("dog", "dog1")
	name, _ := s.dequeueDog()
	assert.Equal(t, "dog5", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "cat3", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "cat2", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "dog4", name)
	name, _ = s.dequeueCat()
	assert.Equal(t, "cat1", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "dog3", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "dog2", name)
	name, _ = s.dequeueAny()
	assert.Equal(t, "dog1", name)
}
