package main

// DO NOT modify the declaration of the Employee struct!
type Employee struct {
	string
	float64
}

func main() {
	var emp Employee
	// Assign values to the two anonymous fields of 'emp' below:
	emp.string = "James"
	emp.float64 = 42.0
	checkAnonymousField(emp) // DO NOT delete this line!
}
