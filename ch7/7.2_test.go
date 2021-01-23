package ch7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Employee is a call-center employee, whether a regular employee,
// a manager, or a director.
type Employee struct {
	name string
	busy bool
}

// dispatchCall chooses a person who is free to take a call, starting with
// employees, then managers, then directors. Returns nil if no-one is available.
func dispatchCall(directors []Employee, managers []Employee, employees []Employee) *Employee {
	for _, employee := range employees {
		if !employee.busy {
			return &employee
		}
	}
	for _, manager := range managers {
		if !manager.busy {
			return &manager
		}
	}
	for _, director := range directors {
		if !director.busy {
			return &director
		}
	}
	return nil
}

// getEmployees returns directors, managers, and regular employees.
func getEmployees() ([]Employee, []Employee, []Employee) {
	return []Employee{Employee{name: "D1"}, Employee{name: "D2"}},
		[]Employee{Employee{name: "M1"}, Employee{name: "M2"}},
		[]Employee{Employee{name: "E1"}, Employee{name: "E2"}}
}

func TestCallCenterNoEscalation(t *testing.T) {
	directors, managers, employees := getEmployees()
	assert.Equal(t, "E1", dispatchCall(directors, managers, employees).name)
}

func TestCallCenterEscalateToManager(t *testing.T) {
	directors, managers, employees := getEmployees()
	employees[0].busy = true
	employees[1].busy = true
	managers[0].busy = true
	assert.Equal(t, "M2", dispatchCall(directors, managers, employees).name)
}

func TestCallCenterEscalateToDirector(t *testing.T) {
	directors, managers, employees := getEmployees()
	employees[0].busy = true
	employees[1].busy = true
	managers[0].busy = true
	managers[1].busy = true
	assert.Equal(t, "D1", dispatchCall(directors, managers, employees).name)
}
