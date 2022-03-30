package handlers

import (
	"fmt"
	"github.com/django-jane/EmployeeManager/Utils/helpers"
	"github.com/django-jane/EmployeeManager/internal/interfaces"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	response helpers.ResponsePayload
	EmployeeRepository interfaces.IEmployeeRepository
}

type Config struct {
	Routing *gin.Engine
	EmployeeRepository interfaces.IEmployeeRepository
	BaseUrl string
}

//Response object

func NewHandler(config *Config) {
		handler := &Handler{

		EmployeeRepository: config.EmployeeRepository,
	}


	//Create routing group
	group := config.Routing.Group(config.BaseUrl)
	var temp = config.BaseUrl
	fmt.Sprintf("Base url is %s",temp)

	group.GET("/All-Employee",handler.RetrieveAllEmployees)
	group.GET("/Find-Employee/:id",handler.FindByID)
	group.POST("/Register-Employee",handler.RegisterEmployee)
	group.PUT("/Update-Employee/:id",handler.UpdateEmployee)
	group.DELETE("/Delete-Employee/:id",handler.DeleteEmployee)

	url := ginSwagger.URL("http://localhost:8090/api/profiler/swagger/doc.json")
	group.GET("/swagger/*any",
					ginSwagger.WrapHandler(
						swaggerFiles.Handler,url))
}