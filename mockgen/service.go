package mockgen

import "github.com/mitooos/thesis/mockgen/model"

type UserRepository interface {
	InsertUser(*model.User) error
}

type SecurityHelper interface {
	HashPassword(string) (string, error)
}

type service struct {
	repository     UserRepository
	securtiyHelper SecurityHelper
}

func (s *service) InsertUser(u *model.User) error {
	var err error
	u.Password, err = s.securtiyHelper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return s.repository.InsertUser(u)
}
