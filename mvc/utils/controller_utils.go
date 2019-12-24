package utils

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondErr(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, body)
		return
	}
	c.JSON(err.StatusCode, body)
}
