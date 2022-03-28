package helpers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ResponsePayload struct {
	Successful bool
	Code int
	Message string
	Data interface{}
	ErrorMessages[] string
}
var response ResponsePayload


//Implementing an extension for the ResponsePayload type
//To enable dispatching a response a little more flexibly
func (resp *ResponsePayload) Dispatch(context *gin.Context) {
	log.Info("Dispatching a json message")
	context.JSON(resp.Code,resp)
return
}


