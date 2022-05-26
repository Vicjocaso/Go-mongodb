package repo

import (
	"github.com/Vicjocaso/Go-mongodb/models"
)

func Create(user models.User) error {
	return nil
}

func Read() ([]models.User, error) {

	var US []models.User
	res := models.User{ID: "123", Name: "hleas", Email: "vi@gdf.com"}
	algo := append(US, res)
	return algo, nil
}

func Update(user models.User, userId string) error {
	return nil
}

func Delete(userId string) error {
	return nil
}
