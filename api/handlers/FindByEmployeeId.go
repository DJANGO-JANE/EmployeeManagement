package handlers

import (
	"github.com/django-jane/EmployeeManager/models"
	"github.com/django-jane/EmployeeManager/models/Employees"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)



func(handler *Handler)FindEmployeeByID(context *gin.Context){
	span, ctx := apm.StartSpan(context.Request.Context(),
							"FindEmployeeByIDHandler",
							"request")

	obj,_ := context.Get("Id")
	objectID := obj.(*models.Employee).Id

	defer span.End()

	log.WithFields(log.Fields{

		"id":obj,
	}).Info("Find employee with")

	employee := Employees.EmployeeViaID{}

	if err := context.ShouldBind(&employee);err !=nil{
		handler.response.Successful=false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
												err.Error())
		handler.response.Code = 501
		handler.response.Dispatch(context)
		return
	}


	result, err := handler.EmployeeRepository.FindById(ctx,objectID)
	if err != nil{
			handler.response.Successful=false
			handler.response.ErrorMessages = append(handler.response.ErrorMessages,
														err.Error())
			handler.response.Code = 502

			handler.response.Dispatch(context)

			return
	}
	log.Info("Handler received data from database")
	handler.response.Successful = true
	handler.response.Code = 200
	handler.response.Data = result
	handler.response.Dispatch(context)

}