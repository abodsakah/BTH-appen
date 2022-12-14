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
	err := helpers.FixtureWrapCreate(t, &helpers.TestExam)
	assert.Nil(t, err, "After fixture is run a create will be called that shall return no error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestCreateExam2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	err := CreateExam(helpers.DbGorm, helpers.TestExam)
	assert.NotNil(t, err, "Trying to create a duplicate exam shall return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestDeleteExam1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	_, err := DeleteExam(helpers.DbGorm, helpers.TestEntryIndex)
	assert.Nil(t, err, "When deleting the object it shall not return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestGetExamsDueSoon(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	exams, _ := GetExamsDueSoon(helpers.DbGorm)
	assert.Less(t, 0, len(exams), "After an exam has been created with the current date plus one day, it should come up in the array of due exams")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestAddUserToExam1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, err := AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	assert.Nil(t, err, "Adding an existing user to an existing exam, with no duplicates, should create no errors")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestAddUserToExam2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	_, err := AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	assert.NotNil(t, err, "Trying to add a duplicate entry should return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestRemoveUserFromExam1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	_, err := RemoveUserFromExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	assert.Nil(t, err, "Removing an existing entry shall not return any errors")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestRemoveUserFromExam2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, err := RemoveUserFromExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	assert.NotNil(t, err, "Removing a non-existent entry shall return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestGetExamUsers1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	_, _ = AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	users, _ := GetExamUsers(helpers.DbGorm, helpers.TestEntryIndex)
	assert.Less(t, 0, len(users), "GetExamUsers should return 1 entry in the array after a user has been added to the exam")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestGetExamUsers2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestExam)
	helpers.DbGorm.Create(&helpers.TestUser)
	users, _ := GetExamUsers(helpers.DbGorm, helpers.TestEntryIndex)
	assert.Equal(t, 0, len(users), "GetExamUsers should return 0 entries in the array when no user has been added to exam")
	_ = helpers.FixtureWrapNonCreate(t)
}
