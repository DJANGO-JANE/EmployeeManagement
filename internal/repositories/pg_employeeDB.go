package repositories

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"strconv"
	"time"
	"unicode"

	"github.com/django-jane/EmployeeManager/internal/interfaces"
	"github.com/django-jane/EmployeeManager/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type pgEmployeeRepository struct {
	Db *sqlx.DB
}

func (p pgEmployeeRepository) UpdateEmployeeInfo(context context.Context, employee *models.Employee) (*models.Employee, error) {
	log.Info("Updating Employee Information")

	sqlStatement := `
					UPDATE employees
					SET FirstName = $1, LastName = $2,
						Department = $3, employeeRole = $4
					WHERE id = $5 RETURNING Id, FirstName, LastName, Department, employeeRole`
	err := p.Db.QueryRow(sqlStatement,
		employee.FirstName,
		employee.LastName,
		employee.Department,
		employee.EmployeeRole,
		employee.Id,
	).Scan(&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.Department,
		&employee.EmployeeRole)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("ERROR : There was an error while updating a row in the database")
	} else {
		log.WithFields(log.Fields{
			"employee": employee,
		}).Info("Row was successfully updated")
		return employee, err
	}
	return employee, err
}

func (p pgEmployeeRepository) RemoveEmployee(context context.Context, employeeId string) (bool, error) {
	sqlStatement := `
					DELETE FROM employees WHERE id = $1;`
	_, err := p.Db.Exec(sqlStatement, employeeId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("ERROR : There was an error while deleting a row in the database")

		return false, err
	} else {
		log.Info("Row deleted from the database")
		return true, err
	}
}

func NewEmployeeRepository(db *sqlx.DB) interfaces.IEmployeeRepository {
	return &pgEmployeeRepository{Db: db}
}

//const letterBytes = "ABCDEFGHJKLMNPQRSTUVWXYZ"

func (p pgEmployeeRepository) GenerateKey(employee *models.Employee) *models.Employee {
	year := time.Now().Year()
	var tempString = randString(5)

	employee.Id = "VGC" + strconv.Itoa(year) + tempString
	return employee
}

/*	func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
	b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
	return string(b)
}*/
func randString(n int) string {
	g := big.NewInt(0)
	max := big.NewInt(130)
	bs := make([]byte, n)

	for i, _ := range bs {
		g, _ = rand.Int(rand.Reader, max)
		r := rune(g.Int64())
		for !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			g, _ = rand.Int(rand.Reader, max)
			r = rune(g.Int64())
		}
		bs[i] = byte(g.Int64())
	}
	return string(bs)
}

func (p pgEmployeeRepository) SignUpNew(context context.Context, employee *models.Employee) (*models.Employee, error) {
	//First setup primary key
	log.Info("Performing an INSERT operation on the database. Generating key first")
	p.GenerateKey(employee)
	log.WithFields(log.Fields{
		"id": employee.Id,
	}).Info("Generated employee id")
	sqlStatement := `INSERT INTO employees (id, firstName, lastName, department,employeeRole) 
		VALUES ($1, $2, $3,$4,$5) RETURNING Id, firstName, lastName, department;`
	if err := p.Db.GetContext(context,
		employee,
		sqlStatement,
		employee.Id,
		employee.FirstName,
		employee.LastName,
		employee.Department,
		employee.EmployeeRole); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Database insert failed")
	}
	return employee, nil
}

func (p pgEmployeeRepository) FindById(context context.Context, employeeId string) (*models.Employee, error) {
	sqlStatement := `SELECT * FROM employees WHERE Id = $1`

	var employee models.Employee

	row := p.Db.QueryRow(sqlStatement, employeeId)

	err := row.Scan(&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.Department,
		&employee.EmployeeRole)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return &employee, nil
	case nil:
		return &employee, nil
	default:
		log.Fatal("Unable to scan the row.")
	}
	return &employee, err
}

func (p pgEmployeeRepository) RetrieveAll(context context.Context) ([]models.Employee, error) {
	sqlStatement := `SELECT Id,
							FirstName,
							LastName,
							Department 
							FROM employees 
							order by LastName DESC;`
	rows, err := p.Db.Query(sqlStatement)
	log.Info("Retrieving all rows from database")
	var employees []models.Employee

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Unable to execute the query")
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err = rows.Scan(&employee.Id,
			&employee.FirstName,
			&employee.LastName,
			&employee.Department)

		if err != nil {
			log.Fatal("Unable to scan row. [%v]")
		}

		employees = append(employees, employee)
	}

	return employees, err
}
