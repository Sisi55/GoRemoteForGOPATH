package services

import (
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/domain"
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implementation me",
		StatusCode: http.StatusInternalServerError,
	}
}
