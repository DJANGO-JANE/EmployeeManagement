package interfaces

import (

	"context"
	"github.com/django-jane/EmployeeManager/models"
)

type IEmployeeRepository interface {
	SignUpNew(context context.Context, employee *models.Employee) (*models.Employee, error)
	FindById(context context.Context, employeeId string) (*models.Employee, error)
	RetrieveAll(context context.Context) ([]models.Employee, error)
	UpdateEmployeeInfo(context context.Context, employee *models.Employee) (*models.Employee, error)
	RemoveEmployee(context context.Context, employeeId string) (bool, error)
}
