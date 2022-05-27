package gomongodb_test

import (
	"testing"
	"time"

	"github.com/Vicjocaso/Go-mongodb/models"
	gomongodb "github.com/Vicjocaso/Go-mongodb/services"
)

func TestCreate(t *testing.T) {
	user := models.User{Name: "Karla", Email: "Karla@Calderon.com", CreateAt: time.Now(), UpdateAt: time.Now()}

	err := gomongodb.Create(user)

	if err != nil {
		t.Error("Test fail")
	} else {
		t.Log("Succes Test!!")
	}
}

func TestRead(t *testing.T) {

	users, err := gomongodb.Read()

	if err != nil {
		t.Error("Error Consulting User")
	}

	if len(users) == 0 {
		t.Error("length no valid")
	} else {
		t.Log("Succes Test")
	}

}

func TestUpdate(t *testing.T) {
	user := models.User{Name: "Victor José", Email: "VCalderon@algo.com"}

	err := gomongodb.Update(user, "628fe8558e8b0ab52c50e8c0")
	if err != nil {
		t.Error("Update Error")
		t.Fail()
	} else {
		t.Log("Update Succes")
	}
}

func TestDelete(t *testing.T) {

	err := gomongodb.Delete("628fe8558e8b0ab52c50e8c0")
	if err != nil {
		t.Error("Delete Error")
		t.Fail()
	} else {
		t.Log("Delete Succes")
	}

}
