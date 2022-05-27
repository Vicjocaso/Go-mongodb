package gomongodb

import (
	"github.com/Vicjocaso/Go-mongodb/models"
	repo "github.com/Vicjocaso/Go-mongodb/repos"
)

func Create(user models.User) error {
	err := repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (models.Users, error) {

	users, err := repo.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user models.User, userId string) error {

	err := repo.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {

	err := repo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
