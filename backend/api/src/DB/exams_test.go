
package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestDatabaseExam(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
	assertNoError(t, err, "Database can not be connected to")
	db = dbP
}

func TestCreateExam1(t *testing.T) {
  _, err := fixtureWrapExam(t)
  assertNoError(t, err, "After fixture is run a create will be called that shall return no error if successful")
}

func TestCreateExam2(t *testing.T) {
  _, err := fixtureWrapExam(t)
  err = CreateExam(db, testExam)
  assertError(t, err, "After already having created an exam with the same credentials, trying to create a duplicate shall return an error")
}

func TestDeleteExam1(t *testing.T) {
  id, err := fixtureWrapExam(t)
  err = DeleteExam(db, id)
  assertNoError(t, err, "After having created an exam, trying to delete it should cause no errors")
}

/*
func TestDeleteExam2(t *testing.T) {
  err := fixtureWrapExam(t)
  id, _ := getExamByName(db, testExam.Name)
  DeleteExam(db, uint(id))
  err = DeleteExam(db, uint(id))
  assertError(t, err, "After having created an exam and deleted it, trying to delete it again shall cause errors")
}
*/
func TestAddUserToExam1(t *testing.T) {
  idExam, _ := fixtureWrapExam(t)
  idUser, _ := createUserWrap()
  err := AddUserToExam(db, idExam, idUser)
  assertNoError(t, err, "Adding an existing user to an existing exam, with no duplicates, should create no errors")
}

func TestAddUserToExam2(t *testing.T) {
  idExam, _ := fixtureWrapExam(t)
  idUser, _ := createUserWrap()
  AddUserToExam(db, idExam, idUser)
  err := AddUserToExam(db, idExam, idUser)
  assertError(t, err, "Adding an existing user to an existing exam, with a duplicate, should create cause errors")
}



