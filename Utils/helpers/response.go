package helpers

type RequestResult struct {
	Code int `json:"code" example:"500"`
	Message string `json:"message" example:"Error accessing the requested resource"`
	Data interface{} `json:"data"`
}

/*func response(c *gin.Context, statType int, err error, data interface{}){
	c.JSON(statType, RequestResult{
		Code: statType,
		Message: err.Error(),
		Data: data,
	})

	return
}*/

/*func FailedResponse(c *gin.Context, err error, data interface{}){
	statCode := Status(err)

	//response(c, Status(err), err, data)
	switch statCode {
	case 400:
		response(c, statCode, err, data)
	case 401:
		response(c, statCode, err, data)
	case 404:
		response(c, statCode, err, data)
	case 409:
		response(c, statCode, err, data)
	case 413:
		response(c, statCode, err, data)
	case 415:
		response(c, statCode, err, data)
	case 500:
		response(c, statCode, err, data)
	default:
		response(c, 503, err, data)
	}
}

func SuccessResponse(c *gin.Context, data interface{}){
	c.JSON(200, RequestResult{
		Code: 200,
		Message: "Success",
		Data: data,
	})

	return
}*/