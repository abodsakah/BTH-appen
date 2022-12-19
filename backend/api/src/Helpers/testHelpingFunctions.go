package helpers

import (
	"testing"
	"time"

	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	models "github.com/abodsakah/BTH-appen/backend/api/src/Models"
	"gorm.io/gorm"
)

var (
	// DbGorm variable
	// Stores database pointer for test-enviorment
	DbGorm           *gorm.DB
	additionalTables = []string{"exam_users"}
)

// TestUser variable
// User entry for test-enviorment
var TestUser = &models.User{
	Name:     "Test Testsson",
	Username: "test",
	Password: "pass",
	Role:     "student",
}

// TestExam variable
// Exam entry for test-enviorment
var TestExam = &models.Exam{
	Name:       "test",
	CourseCode: "pa121212",
	StartDate:  time.Now().AddDate(0, 0, 1),
}

// TestNews variable
// News entry for test-enviorment
var TestNews = &models.News{
	Title:       "Test",
	Date:        time.Now(),
	Description: "A test",
	Link:        "test.com",
}

// SetupTables function
func SetupTables(entries ...interface{}) error {
	var err error
	for _, entry := range entries {
		err = DbGorm.Create(entry).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// FixtureWrapCreate function
func FixtureWrapCreate(t *testing.T, entries ...interface{}) (uint, error) {
	err := fixture.CleanUp(DbGorm, additionalTables, &models.User{}, &models.Exam{}, &models.News{}, &models.Token{})
	if err != nil {
		t.Fatal(err)
	}
	err = SetupTables(entries...)
	if err != nil {
		t.Fatal(err)
	}
	return 1, nil
}

// FixtureWrapNonCreate function
func FixtureWrapNonCreate(t *testing.T) error {
	err := fixture.CleanUp(DbGorm, additionalTables, &models.User{}, &models.Exam{}, &models.News{}, &models.Token{})
	if err != nil {
		t.Fatal(err)
		return err
	}
	return nil
}
