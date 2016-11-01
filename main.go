package main

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
	manager_ids := make(chan uint)

	go func() {
		Search(employees, employeeId, manager_ids)
		close(manager_ids)
	}()

	var managers []uint
	for manager_id := range manager_ids {
		managers = append(managers, manager_id)
	}

	// The first manager of the employee will be himself/herself (edge case for the CEO).
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

func main() {

}
