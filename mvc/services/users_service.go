package services

import (
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/domain"
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
