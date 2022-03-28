package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
)

func(handler *Handler) DeleteEmployee(context *gin.Context){
	span, ctx := apm.StartSpan(context.Request.Context(),
								"DeleteEmployeeHandler",
									"request")

	defer span.End()
	log.Info("DeleteEmployeeHandler received a hit")

	empId := context.Param("id")


	success,err := handler.EmployeeRepository.RemoveEmployee(ctx,empId)
	if err != nil{
		log.WithFields(log.Fields{
			"err":err,
		}).Error("ERROR : An error occurred while deleting employee")

		handler.response.Successful=false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
			err.Error())
		handler.response.Code = 500
		handler.response.Dispatch(context)


		return

	}
	if success == true {
		log.Info("Record successfully deleted")

		handler.response.Successful = true
		handler.response.Code = 200
		handler.response.Data = true
		handler.response.Message = "Successfully deleted"



		handler.response.Dispatch(context)
		return
	}
}
