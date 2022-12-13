package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseExam(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Log("DEV: Could not load .env file")
	}
	dbP, err := SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	db = dbP
}

func TestCreateExam1(t *testing.T) {
	_, err := fixtureWrap(t, &testExam)
	assert.Nil(t, err, "After fixture is run a create will be called that shall return no error")
}

func TestCreateExam2(t *testing.T) {
	_, _ = fixtureWrap(t, &testExam)
	err := CreateExam(db, testExam)
	assert.NotNil(t, err, "Trying to create a duplicate exam shall return an error")
}

func TestDeleteExam1(t *testing.T) {
	id, _ := fixtureWrap(t, &testExam)
	_, err := DeleteExam(db, id)
	assert.Nil(t, err, "When deleting the object it shall not return an error")
}

func TestAddUserToExam1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, err := AddUserToExam(db, idExam, idUser)
	assert.Nil(t, err, "Adding an existing user to an existing exam, with no duplicates, should create no errors")
}

func TestAddUserToExam2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, _ = AddUserToExam(db, idExam, idUser)
	_, err := AddUserToExam(db, idExam, idUser)
	assert.NotNil(t, err, "Trying to add a duplicate entry should return an error")
}

func TestRemoveUserFromExam1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, _ = AddUserToExam(db, idExam, idUser)
	_, err := RemoveUserFromExam(db, idExam, idUser)
	assert.Nil(t, err, "Removing an existing entry shall not return any errors")
}

func TestRemoveUserFromExam2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, err := RemoveUserFromExam(db, idExam, idUser)
	assert.NotNil(t, err, "Removing a non-existent entry shall return an error")
}

func TestGetExamsDueSoon(t *testing.T) {
	_, _ = fixtureWrap(t, &testExam)
	exams, _ := GetExamsDueSoon(db)
	assert.Less(t, 0, len(exams), "After an exam has been created with the current date plus one day, it should come up in the array of due exams")
}

func TestGetExamUsers1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, _ = AddUserToExam(db, idExam, idUser)
	users, _ := GetExamUsers(db, idExam)
	assert.Less(t, 0, len(users), "GetExamUsers should return 1 entry in the array after a user has been added to the exam")
}

func TestGetExamUsers2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	_, _ = createUserWrap()
	users, _ := GetExamUsers(db, idExam)
	assert.Equal(t, 0, len(users), "GetExamUsers should return 0 entries in the array when no user has been added to exam")
}
