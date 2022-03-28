package handlers

import (
	"github.com/django-jane/EmployeeManager/Utils/helpers"
	"github.com/django-jane/EmployeeManager/models"
	"github.com/django-jane/EmployeeManager/models/Employees"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)

func(handler *Handler) UpdateEmployee(context *gin.Context){
	span, ctx := apm.StartSpan(context.Request.Context(),
								"UpdateEmployeeHandler",
								"request")
	id := context.Param("id")

	defer span.End()

	var updateObj Employees.EmployeeUpdate
	log.Info("UpdateEmployeeHandler received a hit")

	if ok := helpers.BindData(context, &updateObj); !ok{

		log.WithFields(log.Fields{
			"!ok":ok,
		}).Error("ERROR : An error occurred when binding an object")

		handler.response.Successful = false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,"Failed")
		handler.response.Code = 500
		handler.response.Dispatch(context)
		return
	}

	log.WithFields(log.Fields{
		"updateObj":updateObj,
	}).Error("Bound query")

	obj := &models.Employee{
		Id: "",
		FirstName:  updateObj.FirstName,
		LastName:   updateObj.LastName,
		Department: updateObj.Department,
EmployeeRole:       updateObj.EmployeeRole,
	}

	accEmployee, err := handler.EmployeeRepository.FindById(context, id)
	if err != nil{
		log.WithFields(log.Fields{
			"err":err,
		}).Error("ERROR : An error occurred when finding employee by ID")
	}
	if len(accEmployee.Id) < 0 {
		log.Info("Employee with that ID cannot be found.")
	}else{
		obj.Id = id

		newEmployee, err := handler.EmployeeRepository.UpdateEmployeeInfo(ctx,obj)
		if err!=nil{
			log.WithFields(log.Fields{
				"err":err,
			}).Error("ERROR : An error occurred while updating employee info")
		}else{
			log.Info("Handler received data from database")

			handler.response.Successful = true
			handler.response.Code = 200
			handler.response.Message = "Employee Info Updated Successfully"
			handler.response.Data = newEmployee


			handler.response.Dispatch(context)
			return
		}
	}

}
