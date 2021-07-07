package mockgen

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	mock_mockgen "github.com/mitooos/thesis/mockgen/mocks"
	"github.com/mitooos/thesis/mockgen/model"
)

func TestInsertUser(t *testing.T) {
	t.Run("inserts user succesfully", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}

		ctrl := gomock.NewController(t)

		mockSecurityHelper := mock_mockgen.NewMockSecurityHelper(ctrl)
		mockSecurityHelper.EXPECT().HashPassword(usr.Password).Return("hashed", nil)

		mockRepository := mock_mockgen.NewMockUserRepository(ctrl)
		mockRepository.EXPECT().InsertUser(usr).Return(nil)

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		if got != nil {
			t.Errorf("got not nil reponse: %v", got)
		}

		if "hashed" != usr.Password {
			t.Errorf("got %v, want %v", usr.Password, "hashed")
		}
	})

	t.Run("returns error if hashing password fail", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}
		ctrl := gomock.NewController(t)

		mockSecurityHelper := mock_mockgen.NewMockSecurityHelper(ctrl)
		mockSecurityHelper.EXPECT().HashPassword(usr.Password).Return("hashed", nil)

		mockRepository := mock_mockgen.NewMockUserRepository(ctrl)
		mockRepository.EXPECT().InsertUser(usr).Return(nil)

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		if got != nil && got.Error() != "error hashing" {
			t.Errorf("got %v, want %v", got.Error(), "error hashing")
		}
	})

	t.Run("returns an error if InsertUser from repository returns one", func(t *testing.T) {
		usr := &model.User{
			Email:    "email@email.com",
			Username: "mock-usrname",
			Password: "pwd",
		}

		ctrl := gomock.NewController(t)

		mockSecurityHelper := mock_mockgen.NewMockSecurityHelper(ctrl)
		mockSecurityHelper.EXPECT().HashPassword(usr.Password).Return("hashed", nil)

		mockRepository := mock_mockgen.NewMockUserRepository(ctrl)
		mockRepository.EXPECT().InsertUser(usr).Return(fmt.Errorf("error inserting"))

		service := new(service)
		service.securtiyHelper = mockSecurityHelper
		service.repository = mockRepository

		got := service.InsertUser(usr)

		if got.Error() != "error inserting" {
			t.Errorf("got %v, want %v", got.Error(), "error inserting")
		}
	})
}
