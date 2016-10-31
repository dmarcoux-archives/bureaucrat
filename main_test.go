package main

import "testing"

var Employees = Employee{
	Id:   1,
	Name: "Claire",
	Subordinates: []*Employee{
		&Employee{
			Id:   2,
			Name: "Roger",
			Subordinates: []*Employee{
				&Employee{
					Id:   3,
					Name: "George",
				},
				&Employee{
					Id:   4,
					Name: "Suzie",
				},
				&Employee{
					Id:   5,
					Name: "Lola",
					Subordinates: []*Employee{
						&Employee{
							Id:   6,
							Name: "Foo",
						},
					},
				},
			},
		},
		&Employee{
			Id:   7,
			Name: "Bar",
		},
		&Employee{
			Id:   8,
			Name: "Paul",
			Subordinates: []*Employee{
				&Employee{
					Id:   9,
					Name: "Jen",
				},
				&Employee{
					Id:   10,
					Name: "Ringo",
				},
			},
		},
	},
}

func TestSearcherEmptyDirectory(t *testing.T) {
	managers := Searcher(nil, 1)

	if <-managers != 0 {
		t.Error("Searcher shouldn't return managers when searching through an empty employee directory")
	}
}

func TestSearcherNonexistentEmployee(t *testing.T) {
	managers := Searcher(&Employees, 9001)

	if <-managers != 0 {
		t.Error("Searcher shoudn't return managers for a nonexistent employee")
	}
}

func TestSearcherCEO(t *testing.T) {
	managers := Searcher(&Employees, 1)

	if <-managers != 0 {
		t.Error("Searcher shouldn't return managers for the CEO")
	}
}

func TestSearcherAnyEmployee(t *testing.T) {
	var expected_manager uint

	expected_managers := []uint{1}
	managers := Searcher(&Employees, 8)
	for manager := range managers {
		// Pop the current expected manager from the slice...
		expected_manager, expected_managers = expected_managers[len(expected_managers)-1], expected_managers[:len(expected_managers)-1]
		if manager != expected_manager {
			t.Error("Searcher should return the CEO as the manager of any employee directly under him/her")
		}
	}

	expected_managers = []uint{1, 8}
	managers = Searcher(&Employees, 10)
	for manager := range managers {
		// Pop the current expected manager from the slice...
		expected_manager, expected_managers = expected_managers[len(expected_managers)-1], expected_managers[:len(expected_managers)-1]
		if manager != expected_manager {
			t.Error("Searcher should return the managers of any employee")
		}
	}

	expected_managers = []uint{1, 2, 5}
	managers = Searcher(&Employees, 6)
	for manager := range managers {
		// Pop the current expected manager from the slice...
		expected_manager, expected_managers = expected_managers[len(expected_managers)-1], expected_managers[:len(expected_managers)-1]
		if manager != expected_manager {
			t.Error("Searcher should return the managers of any employee")
		}
	}
}
