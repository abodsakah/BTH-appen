package helpers

import (
	"testing"
	"time"

	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	models "github.com/abodsakah/BTH-appen/backend/api/src/Models"
	"gorm.io/gorm"
)

var (
	DbGorm           *gorm.DB
	additionalTables = []string{"exam_users"}
)

var TestUser = &models.User{
	Name:     "Test Testsson",
	Username: "test",
	Password: "pass",
	Role:     "student",
}

var TestExam = &models.Exam{
	Name:       "test",
	CourseCode: "pa121212",
	StartDate:  time.Now().AddDate(0, 0, 1),
}

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
