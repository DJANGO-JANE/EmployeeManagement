package handlers

import (
	"github.com/django-jane/EmployeeManager/Utils/helpers"
	"github.com/django-jane/EmployeeManager/internal/interfaces"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	response helpers.ResponsePayload
	//tokenMaker         token.Maker
	EmployeeRepository interfaces.IEmployeeRepository
}

type Config struct {
	Routing            *gin.Engine
	EmployeeRepository interfaces.IEmployeeRepository
	BaseUrl            string
}

//Response object

func NewHandler(config *Config) {
	//tokenMaker, err := token.NewPasetoMaker("")
	handler := &Handler{

		EmployeeRepository: config.EmployeeRepository,
	}

	//Create routing group
	group := config.Routing.Group(config.BaseUrl)

	group.GET("/:id", handler.FindEmployeeByID) //
	//group.GET("/all-employee", handler.RetrieveAllEmployees)  This endpoint doesn't get hit and is still under development
	group.POST("/register-employee", handler.RegisterEmployee)
	group.PUT("/update-employee/:id", handler.UpdateEmployee)
	group.DELETE("/delete-employee/:id", handler.DeleteEmployee)

	url := ginSwagger.URL("http://localhost:8090/api/profiler/swagger/doc.json")
	group.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler, url))
}
