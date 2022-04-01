package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/django-jane/EmployeeManager/internal/persistence"
	log "github.com/sirupsen/logrus"
)

func main() {

	//Seed and initialise data
	data, err := persistence.CreateConnection()

	if err != nil {
		fmt.Printf("Unable to initialise connection")
	}
	if data != nil {
		fmt.Printf("data sources configured")
	}

	//If there was no error in connecting to the database
	//Then proceed to inject the dependencies
	router, err := persistence.Inject(data)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Info("Dependency injection failed")

	} else {
		log.Info("Dependencies successfully injected")
	}

	port := os.Getenv("EMPLOYEE_MAN_PORT")

	if port == "" {
		port = "80"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.WithFields(log.Fields{
			"server": server,
			"port":   port,
			"error":  err,
		}).Info("Tried to open server but failed to open")
	}

}
