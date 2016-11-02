package main

import "testing"

var Employees = Employee{
	Id:   1,
	Name: "Claire",
	Subordinates: []*Employee{
		{
			Id:   2,
			Name: "Roger",
			Subordinates: []*Employee{
				{
					Id:   3,
					Name: "George",
				},
				{
					Id:   4,
					Name: "Suzie",
				},
				{
					Id:   5,
					Name: "Lola",
					Subordinates: []*Employee{
						{
							Id:   6,
							Name: "Foo",
						},
					},
				},
			},
		},
		{
			Id:   7,
			Name: "Bar",
		},
		{
			Id:   8,
			Name: "Paul",
			Subordinates: []*Employee{
				{
					Id:   9,
					Name: "Jen",
				},
				{
					Id:   10,
					Name: "Ringo",
				},
			},
		},
	},
}

func TestSearcherEmptyDirectory(t *testing.T) {
	managers := Searcher(nil, 1)

	if managers != nil {
		t.Error("Searcher shouldn't return managers when searching through an empty employee directory")
	}
}

func TestSearcherNonexistentEmployee(t *testing.T) {
	managers := Searcher(&Employees, 9001)

	if managers != nil {
		t.Error("Searcher shoudn't return managers for a nonexistent employee")
	}
}

func TestSearcherCEO(t *testing.T) {
	managers := Searcher(&Employees, 1)

	if len(managers) > 1 && managers[0] != 1 {
		t.Error("Searcher should return the CEO as its own manager")
	}
}

func TestSearcherAnyEmployee(t *testing.T) {
	// First level...
	expected_managers := []uint{1}
	managers := Searcher(&Employees, 8)
	for i, manager := range managers {
		if manager != expected_managers[i] {
			t.Error("Searcher should return the managers of any employee, including himself/herself")
		}
	}

	// Second level employee...
	expected_managers = []uint{8, 1}
	managers = Searcher(&Employees, 10)
	for i, manager := range managers {
		if manager != expected_managers[i] {
			t.Error("Searcher should return the managers of any employee, including himself/herself")
		}
	}

	// Third level employee...
	expected_managers = []uint{5, 2, 1}
	managers = Searcher(&Employees, 6)
	for i, manager := range managers {
		if manager != expected_managers[i] {
			t.Error("Searcher should return the managers of any employee, including himself/herself")
		}
	}
}

func TestFindCommonManagerEmptyDirectory(t *testing.T) {
	if FindCommonManager(nil, 1, 2) != 0 {
		t.Error("FindCommonManager shouldn't return a manager when searching through an empty employee directory")
	}
}

func TestFindCommonManagerOneNonexistentEmployee(t *testing.T) {
	if FindCommonManager(&Employees, 9001, 2) != 0 {
		t.Error("FindCommonManager shouldn't return a manager for 2 employees, when one of them is nonexistent")
	}
}

func TestFindCommonManagerOneEmployeeIsCEO(t *testing.T) {
	var expected_manager uint = 1
	manager := FindCommonManager(&Employees, 1, 10)
	if manager != expected_manager {
		t.Error("FindCommonManager should return the CEO as the closest common manager for 2 employees, when one of them is the CEO")
	}
}

func TestFindCommonManagerSameEmployee(t *testing.T) {
	var expected_manager uint = 8
	manager := FindCommonManager(&Employees, 9, 9)
	if manager != expected_manager {
		t.Error("FindCommandManager should return the closest common manager for twice the same employee")
	}
}

func TestFindCommonManagerDifferentEmployees(t *testing.T) {
	var expected_manager uint = 8
	manager := FindCommonManager(&Employees, 9, 10)
	if manager != expected_manager {
		t.Error("FindCommandManager should return the closest common manager for 2 different employees")
	}

	expected_manager = 2
	manager = FindCommonManager(&Employees, 3, 6)
	if manager != expected_manager {
		t.Error("FindCommandManager should return the closest common manager for 2 different employees")
	}

	expected_manager = 1
	manager = FindCommonManager(&Employees, 6, 9)
	if manager != expected_manager {
		t.Error("FindCommandManager should return the closest common manager for 2 different employees")
	}

}
