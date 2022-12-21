package db

import (
	"testing"

	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"
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
	helpers.DbGorm = dbP
  _ = helpers.FixtureWrapNonCreate(t)
}

func TestCreateExam1(t *testing.T) {
	_, err := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	assert.Nil(t, err, "After fixture is run a create will be called that shall return no error")
}

func TestCreateExam2(t *testing.T) {
	_, _ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	err := CreateExam(helpers.DbGorm, helpers.TestExam)
	assert.NotNil(t, err, "Trying to create a duplicate exam shall return an error")
}

func TestDeleteExam1(t *testing.T) {
	id, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	_, err := DeleteExam(helpers.DbGorm, id)
	assert.Nil(t, err, "When deleting the object it shall not return an error")
}

func TestAddUserToExam1(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, err := AddUserToExam(helpers.DbGorm, idExam, helpers.TestUser.ID)
	assert.Nil(t, err, "Adding an existing user to an existing exam, with no duplicates, should create no errors")
}

func TestAddUserToExam2(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, idExam, helpers.TestUser.ID)
	_, err := AddUserToExam(helpers.DbGorm, idExam, helpers.TestExam.ID)
	assert.NotNil(t, err, "Trying to add a duplicate entry should return an error")
}

func TestRemoveUserFromExam1(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, idExam, helpers.TestUser.ID)
	_, err := RemoveUserFromExam(helpers.DbGorm, idExam, helpers.TestExam.ID)
	assert.Nil(t, err, "Removing an existing entry shall not return any errors")
}

func TestRemoveUserFromExam2(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, err := RemoveUserFromExam(helpers.DbGorm, idExam, helpers.TestUser.ID)
	assert.NotNil(t, err, "Removing a non-existent entry shall return an error")
}

func TestGetExamsDueSoon(t *testing.T) {
	_, _ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	exams, _ := GetExamsDueSoon(helpers.DbGorm)
	assert.Less(t, 0, len(exams), "After an exam has been created with the current date plus one day, it should come up in the array of due exams")
}

func TestGetExamUsers1(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, idExam, helpers.TestUser.ID)
	users, _ := GetExamUsers(helpers.DbGorm, idExam)
	assert.Less(t, 0, len(users), "GetExamUsers should return 1 entry in the array after a user has been added to the exam")
}

func TestGetExamUsers2(t *testing.T) {
	idExam, _ := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	users, _ := GetExamUsers(helpers.DbGorm, idExam)
	assert.Equal(t, 0, len(users), "GetExamUsers should return 0 entries in the array when no user has been added to exam")
}
