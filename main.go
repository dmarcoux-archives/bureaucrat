package main

import "fmt"

// An employee can have subordinates, which in turn can have subordinates and so on...
// In terms of data structure, an employee is a node in a general tree
// An employee (a node) will be either:
// - The CEO (the root node)
// - A manager; an employee with at least one subordinate (a parent node with at least one child)
// - A normal employee; an employee without subordinates (a leaf node)
type Employee struct {
	Id           uint
	Name         string
	Subordinates []*Employee
}

// Searches for an employee in an employee tree, sending the id(s) of his/her manager(s) on a channel
func Search(employee *Employee, employeeId uint, channel chan uint) bool {
	if employee == nil {
		return false
	}

	if employee.Id == employeeId {
		// An employee will always be considered his own manager, in addition to any manager(s) he/she has
		// This is done for the edge case of the CEO, which wouldn't have a manager otherwise.
		// We can argue that the CEO is his/her own boss or not... This would be done in the code review...
		channel <- employee.Id
		return true
	}

	for _, subordinate := range employee.Subordinates {
		if Search(subordinate, employeeId, channel) {
			channel <- employee.Id
			return true
		}
	}

	return false
}

// Launches a search in an employee tree for a specific employee with his/her id
// Returns a slice of id(s) from his/her manager(s)
func Searcher(employees *Employee, employeeId uint) []uint {
	managerIds := make(chan uint)

	go func() {
		Search(employees, employeeId, managerIds)
		close(managerIds)
	}()

	var managers []uint
	for managerId := range managerIds {
		managers = append(managers, managerId)
	}

	// The first manager of the employee will always be himself/herself (edge case for the CEO).
	// So if he/she has at least another manager, remove himself/herself from his/her managers
	if len(managers) > 1 {
		managers = managers[1:]
	}

	return managers
}

// Finds the closest common manager (farthest from the top manager) between two employees
func FindCommonManager(employees *Employee, employee1Id, employee2Id uint) uint {
	if employees == nil {
		return 0
	}

	if employee1Id == employees.Id || employee2Id == employees.Id {
		return employees.Id
	}

	managersEmployee1 := Searcher(employees, employee1Id)
	managersEmployee2 := Searcher(employees, employee2Id)

	var commonManager uint
	for _, manager1 := range managersEmployee1 {
		for _, manager2 := range managersEmployee2 {
			if manager1 == manager2 {
				commonManager = manager1
				break
			}
		}

		if commonManager != 0 {
			break
		}
	}

	return commonManager
}

// This is the structure of the employees directory, numbers being the ids of employees
//      1
//    / | \
//   2  7  8
//  /|\   / \
// 3 4 5 9  10
//      \
//       6
var EmployeesDirectory *Employee = &Employee{
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

func main() {
	fmt.Println("Welcome! This is Bureaucr.at's employees directory.")
	fmt.Println("--------------------")
	fmt.Println("To find the common closest manager of 2 employees, please provide their id separated by a space (Ids must be greater than 0)")
	var employee1 uint
	var employee2 uint
	_, err := fmt.Scanf("%d %d", &employee1, &employee2)
	if err == nil && employee1 > 0 && employee2 > 0 {
		commonManager := FindCommonManager(EmployeesDirectory, employee1, employee2)
		fmt.Printf("For the employees #%d and #%d, their common closest manager is the employee #%d\n", employee1, employee2, commonManager)
	} else {
		fmt.Println("Wrong input provided")
	}
	fmt.Println("Exiting now...")
}
