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
// Returns a read-only channel of id(s) from his/her manager(s)
func Searcher(employees *Employee, employeeId uint) <-chan uint {
	managers := make(chan uint)

	go func() {
		Search(employees, employeeId, managers)
		close(managers)
	}()

	return managers
}

func main() {

}
