package mockery

import (
	"fmt"
	"testing"

	"github.com/mitooos/thesis/mockery/mocks"
	"github.com/mitooos/thesis/mockery/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {
	t.Run("inserts user succesfully", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}

		mockSecurityHelper := new(mocks.SecurityHelper)
		mockSecurityHelper.On("HashPassword", usr.Password).Return("hashed", nil)

		mockRepository := new(mocks.UserRepository)
		mockRepository.On("InsertUser", usr).Return(nil)

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		mockSecurityHelper.AssertExpectations(t)
		mockRepository.AssertExpectations(t)

		assert.Nil(t, got)
		assert.Equal(t, "hashed", usr.Password)
	})

	t.Run("returns error if hashing password fail", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}

		mockSecurityHelper := new(mocks.SecurityHelper)
		mockSecurityHelper.On("HashPassword", usr.Password).Return("hashed", fmt.Errorf("error hashing"))

		mockRepository := new(mocks.UserRepository)

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		mockRepository.AssertNotCalled(t, "InsertUser", mock.Anything)
		mockSecurityHelper.AssertExpectations(t)
		mockRepository.AssertExpectations(t)

		assert.EqualError(t, got, "error hashing")
	})

	t.Run("returns an error if InsertUser from repository returns one", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}

		mockSecurityHelper := new(mocks.SecurityHelper)
		mockSecurityHelper.On("HashPassword", usr.Password).Return("hashed", nil)

		mockRepository := new(mocks.UserRepository)
		mockRepository.On("InsertUser", usr).Return(fmt.Errorf("error inserting"))

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		mockSecurityHelper.AssertExpectations(t)
		mockRepository.AssertExpectations(t)

		assert.EqualError(t, got, "error inserting")
	})
}
