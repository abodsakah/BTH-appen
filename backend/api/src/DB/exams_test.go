package db

import (
	"testing"

	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseExam(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
	fixture.AssertNoError(t, err, "Database can not be connected to")
	db = dbP
}

func TestCreateExam1(t *testing.T) {
	_, err := fixtureWrap(t, &testExam)
	fixture.AssertNoError(t, err, "After fixture is run a create will be called that shall return no error if successful")
}

func TestCreateExam2(t *testing.T) {
	_, err := fixtureWrap(t, &testExam)
	err = CreateExam(db, testExam)
	fixture.AssertError(t, err, "After already having created an exam with the same credentials, trying to create a duplicate shall return an error")
}

func TestDeleteExam1(t *testing.T) {
	id, _ := fixtureWrap(t, &testExam)
	DeleteExam(db, id)
	var exam Exam
	res, _ := fixture.CheckIfDeleted(db, id, &exam)
	assert.Equal(t, true, res, "After having deleted an exam, it should not be in database anymore")
}

func TestAddUserToExam1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, err := AddUserToExam(db, idExam, idUser)
	fixture.AssertNoError(t, err, "Adding an existing user to an existing exam, with no duplicates, should create no errors")
}

func TestAddUserToExam2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	AddUserToExam(db, idExam, idUser)
	_, err := AddUserToExam(db, idExam, idUser)
	fixture.AssertError(t, err, "Adding an existing user to an existing exam, with a duplicate, should cause errors")
}

func TestRemoveUserFromExam1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	AddUserToExam(db, idExam, idUser)
	_, err := RemoveUserFromExam(db, idExam, idUser)
	fixture.AssertNoError(t, err, "Removing an existing entry in exam2user table shall not return any errors")
}

func TestRemoveUserFromExam2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	_, err := RemoveUserFromExam(db, idExam, idUser)
	fixture.AssertError(t, err, "Removing a non-existent entry in exam2user table shall return an error")
}

func TestGetExamsDueSoon(t *testing.T) {
	fixtureWrap(t, &testExam)
	exams, _ := GetExamsDueSoon(db)
	assert.Less(t, 0, len(exams), "After an exam has been created with the current date, it should come up in the array of due exams")
}

func TestGetExamUsers1(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	idUser, _ := createUserWrap()
	AddUserToExam(db, idExam, idUser)
	users, _ := GetExamUsers(db, idExam)
	assert.Less(t, 0, len(users), "GetExamUsers should return 1 entry in the array after a user has been added to the exam")
}

func TestGetExamUsers2(t *testing.T) {
	idExam, _ := fixtureWrap(t, &testExam)
	createUserWrap()
	users, _ := GetExamUsers(db, idExam)
	assert.Equal(t, 0, len(users), "GetExamUsers should return 0 entries in the array when no user has been added to exam")
}
