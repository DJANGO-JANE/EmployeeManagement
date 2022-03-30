package Employees

type EmployeeViaID struct {
	ID string `uri:"id" binding:"required,min=10"`
}
