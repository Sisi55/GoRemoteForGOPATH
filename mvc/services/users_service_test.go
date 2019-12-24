package services

import (
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/domain"
	"github.com/Sisi55/GoRemoteForGOPATH/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	UserDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {

	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exists",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	user, err := UsersService.GetUser(0)
	// service 가 dao 에 접근해서 데이터 가져오는 과정이
	// 같은 과정으로 service 가 dao 에 접근해서 가져오는데
	// dao 접근하는 과정을 가로채서 mock 객체를 주입한거였다!

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
