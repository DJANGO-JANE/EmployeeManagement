package handlers

import (
	"github.com/django-jane/EmployeeManager/Utils/helpers"
	"github.com/django-jane/EmployeeManager/models"
	"github.com/django-jane/EmployeeManager/models/Employees"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)

func(handler *Handler) RegisterEmployee(context *gin.Context){
	span, ctx := apm.StartSpan(context.Request.Context(),"RegisterEmployeeHandler","request")
	defer span.End()

	var request Employees.SignUpPayLoad

	if ok:= helpers.BindData(context, &request); !ok{
		return
	}
	employee := models.Employee{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Department:   request.Department,
		EmployeeRole: request.EmployeeRole,
	}

	log.WithFields(log.Fields{

		"employee" :employee,

	}).Info("Registration submission payload")

	employeeFromDb, err := handler.EmployeeRepository.SignUpNew(ctx,&employee)
	if err != nil{

		handler.response.Successful=false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
												err.Error())
		handler.response.Code = 500

		log.WithFields(log.Fields{
			"employee": employeeFromDb,
			"err":err,
		}).Error("An error occurred while trying to register an employee")
		handler.response.Dispatch(context)
		return
	} else {

		log.Info("Just finished a registration")
		handler.response.Successful = true
		handler.response.Code = 200
		handler.response.Data = employeeFromDb
		handler.response.Message = "Employee Successfully registered."
		handler.response.Dispatch(context)
	}
}