// Unit tests for route module
package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func startTest(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, helpers.TestExam, helpers.TestNews)
	temp := *helpers.TestUser
	_ = db.CreateUser(helpers.DbGorm, &temp)
	temp = *helpers.TestAdmin
	_ = db.CreateUser(helpers.DbGorm, &temp)
}

func setupContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// make sure c.Request is not nil
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	return c, w
}

func mockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestDatabaseRoutes(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Log("DEV: Could not load .env file")
	}
	dbP, err := db.SetupDatabase()
	if dbP == nil {
		fmt.Printf("Is nil!!!!!!!!!!!!!!!!!!!!!1111\n")
	}
	assert.Nil(t, err, "Database can not be connected to")
	helpers.DbGorm = dbP
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestSetupRoutes(t *testing.T) {
	var err error
	router, err = SetupRoutes(helpers.DbGorm)
	assert.Nil(t, err, "Trying to set up routes shall not cause any errors")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestLogin1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Username": "test",
		"Password": "pass",
	})

	// call API endpoint
	login(c)
	assert.Equal(t, 200, w.Code, "When trying to log in with correct admin credentials it shall return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestLogin2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Username": "admin",
		"Password": "incorrect-password",
	})

	// call API endpoint
	login(c)
	assert.NotEqual(t, 200, w.Code, "When trying to log in with incorrect admin credentials it shall not return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestExamCreate1(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"name":        "Test exam",
		"course_code": "something",
		"room":        "something",
		"start_date":  time.Now(),
	})

	// call API endpoint
	createExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on create Exam API with no duplicates present, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestExamCreate2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"name":        "Test exam",
		"course_code": "something",
		"room":        "something",
		"start_date":  time.Now(),
	})

	// call API endpoint
	createExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on create Exam API with duplicates present, it should not return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestExamDelete1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})

	// call API endpoint
	deleteExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on delete Exam API with entry present, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestExamDelete2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})

	// call API endpoint
	deleteExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on delete Exam API with no entry present, it should not return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestListExams1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})

	// call API endpoint
	listExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list Exams API with entries present, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestListDueExams1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})

	// call API endpoint
	listDueExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list due Exams API with due exams present, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestListUserExams(t *testing.T) {
	startTest(t)
	_, _ = db.AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"userID":  helpers.TestEntryIndex,
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	listUserExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list User Exams API with exams present for user, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestListExamUser(t *testing.T) {
	startTest(t)
	_, _ = db.AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	listExamUsers(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list Exam Users API with users present for user, it should return status: 200")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestRegisterToExam1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	registerToExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on registerToExam API with user present, it should return status: 200")
}

func TestRegisterToExam2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	registerToExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on registerToExam API with no user present, it should not return status: 200")
}

func TestUnregisterFromExam1(t *testing.T) {
	startTest(t)
	_, _ = db.AddUserToExam(helpers.DbGorm, helpers.TestEntryIndex, helpers.TestEntryIndex)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	unregisterFromExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on Unregister from Exam API with user present in exam2user table, it should return status: 200")
}

func TestUnregisterFromExam2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": helpers.TestEntryIndex,
	})
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	unregisterFromExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on Unregister from Exam API with no user present in exam2user table, it should not return status: 200")
}

func TestListNews1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	assert.Equal(t, 200, w.Code, "When trying to call on ListNews API with news present in news table, it should return status: 200")
}

func TestListNews2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	c.Set("UserID", uint(helpers.TestEntryIndex))
	// call API endpoint
	assert.Equal(t, 200, w.Code, "When trying to call on ListNews API with no news present in news table, it should not return status: 200")
}

func TestCreateUser1(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Name":     "Admin Adminsson",
		"Username": "admin",
		"Password": "pass",
		"Role":     "admin",
	})

	assert.Equal(t, 200, w.Code, "When trying to call on Create User API with correct credentials, it should return status: 200")
}

func TestCreateUser2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Name":     "Admin Adminsson",
		"Username": "admin",
		"Password": "too-loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong-password",
		"Role":     "admin",
	})
	createUser(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on Create User API with correct credentials except a password which is too long, it should not return status: 200")
}

func TestCreateUser3(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Name":     "Admin Adminsson",
		"Password": "pass",
		"Role":     "admin",
	})
	createUser(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on Create User API with missing username, it should not return status: 200")
}

// Middlewares
func TestAuthMiddleware1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Username": "test",
		"Password": "pass",
	})

	login(c)

	authMiddleware(c)
	assert.Equal(t, 200, w.Code, "When trying to call on Auth Middleware API with user being logged in, it should return status: 200")
}

func TestAuthMiddleware2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type

	authMiddleware(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on Auth Middleware API with user not logged in, it should not return status: 200")
}

func TestAdminMiddleware1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type

	c.Set("UserID", uint(helpers.TestEntryIndex+1))
	adminMiddleware(c)
	assert.Equal(t, 200, w.Code, "When trying to call on Admin Middleware API with right ID for admin, it should return status: 200")
}

func TestAdminMiddleware2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type

	c.Set("UserID", uint(helpers.TestEntryIndex))
	adminMiddleware(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on Admin Middleware API with wrong ID for admin, it should not return status: 200")
}
