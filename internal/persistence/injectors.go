package persistence

import (
	"fmt"
	"os"

	"github.com/django-jane/EmployeeManager/api/handlers"
	"github.com/django-jane/EmployeeManager/internal/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
)

func Inject(ds *dataSources) (*gin.Engine, error) {
	//Set the router as the default one provided by Gin
	router := gin.Default()

	//Configuring router to use a specified middleware
	//apmgin to provide middleware

	router.Use(apmgin.Middleware(router))
	baseUrl := fmt.Sprintf("%s", os.Getenv("EMPLOYEE_MAN_API_URL"))

	//EmployeeRepository constructor dependencies
	employeeRepository := repositories.NewEmployeeRepository(ds.DB)

	log.WithFields(log.Fields{
		"BaseUrl": baseUrl,
	}).Info("Handler initialised")

	handlers.NewHandler(&handlers.Config{
		Routing:            router,
		EmployeeRepository: employeeRepository,
		BaseUrl:            baseUrl,
	})
	return router, nil
}
