package main

import (
	"fmt"
	"github.com/django-jane/EmployeeManager/internal/persistence"
	_ "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() //Flushes the buffer if any
	logger.Info("Logger successfully configured")

	//Seed and initialise data
	data, err := persistence.CreateConnection()
	logger.Info("Data has been loaded")
	if err != nil {
		fmt.Printf("Unable to initialise connection")
	}
	if data != nil {
		fmt.Printf("data sources configured")
	}
	//If there was no error in connecting to the database
	//Then proceed to inject the dependencies
	router, err := persistence.Inject(data)
	logger.Info("Dependencies successfully injected")

	port := os.Getenv("EMPLOYEE_MAN_PORT")

	if port ==""{
		port = "80"
	}
	server := &http.Server{
		Addr: ":"+port,
		Handler: router,
	}



		if err := server.ListenAndServe(); err!=nil && err != http.ErrServerClosed{
			log.WithFields(log.Fields{
				"server": server,
				"port": port,
			}).Info("Tried to open server but failed to open")
		}



	//Set the router as the default one provided by Gin
/*	router = gin.Default()
	logger.Info("Running gin")

	router.Run()*/
}
