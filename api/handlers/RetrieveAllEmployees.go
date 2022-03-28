package handlers

import (
	"github.com/django-jane/EmployeeManager/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)


func(handler *Handler) RetrieveAllEmployees(context *gin.Context){
	span, ctx := apm.StartSpan(context.Request.Context(),
						"RetrieveAllEmployeesHandler",
						"request")
	defer span.End()

	var results []models.Employee

	if err := context.BindQuery(&results); err !=nil{
		handler.response.Successful=false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
												err.Error())
		handler.response.Code = 500
		handler.response.Dispatch(context)



		return
	}

	results, err := handler.EmployeeRepository.RetrieveAll(ctx)
	if err != nil{

		log.WithFields(log.Fields{
			"err":err,
		}).Error("ERROR : An error occurred while retrieving all employees")
		handler.response.Successful=false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
												err.Error())
		handler.response.Code = 500
		handler.response.Dispatch(context)


		return
	} else{

		log.Info("Handler received data from database")

		handler.response.Successful = true
		handler.response.Code = 200
		handler.response.Data = results


		handler.response.Dispatch(context)
		return
	}
}
