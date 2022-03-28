package helpers


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func BindData(context *gin.Context, results interface{}) bool{
/*	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter:  &easy.Formatter{
			TimestampFormat: "15:04:05 2006-01-02",
			LogFormat:       "[%lvl%]: %time% - %msg%",
			ForceFormatting: true,
		},
	}*/
  temp :=context.Request.PostForm.Has("request")
	//log.Formatter = new(prefixed.TextFormatter)
	//log.Level = logrus.DebugLevel
	/*Bind incoming json to struct and check for validation errors
	send error if Content-Type != application/json */
	log.WithFields(log.Fields{
		"context":context.Value("Request"),
		"temp":temp,

	}).Info("Binding the context")
	if context.ContentType() != "application/json"{
		msg := fmt.Sprintf("%s only accepts Content-Type of \"application/json\"",
			context.FullPath())

		
		err := NewUnsupportedMediaType(msg)
		response.Successful=false
		response.ErrorMessages = append(response.ErrorMessages,err.Error())
		response.Code = 500

		response.Dispatch(context)
		return false
	}

		log.WithFields(log.Fields{
			"results":results,
	}).Info("Binding data")

	if err:= context.ShouldBind(results); err != nil{

		if errs, ok := err.(validator.ValidationErrors);ok{
			var invalidArgs []InvalidArgument

			for _,err := range errs{
				invalidArgs = append(invalidArgs, InvalidArgument{
					Field: err.Field(),
					Value : err.Value().(string),
					Tag: err.Tag(),
					Param :err.Param(),
				})
				log.WithFields(log.Fields{
					"field": err.Field(),
				}).Info("ERROR : An error occurred when binding data")

			}

			err := NewBadRequest("invalid request parameters. See invalidArgs")
			response.Successful=false
			response.ErrorMessages = append(response.ErrorMessages,
											err.Error())
			response.Code = 506

			response.Dispatch(context)
			log.WithFields(log.Fields{

				"err" :err,

			}).Info("ERROR : An error occurred ")
			return false
		}
			response.Successful=false
			response.ErrorMessages = append(response.ErrorMessages,
											NewInternal().Message)
			response.Code = 502
		log.WithFields(log.Fields{

			"err" :err,

		}).Info("ERROR : An error occurred ")
			response.Dispatch(context)
		return false
	}
	return true
}