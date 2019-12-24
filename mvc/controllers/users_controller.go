package controllers

import (
	"encoding/json"
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/services"
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	//c.Query("caller")
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondErr(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
