// Unit tests for route module
package routes

import (
	"testing"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"

	"net/http"
	"net/http/httptest"

	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

var (
  dbGorm *gorm.DB
  router *gin.Engine
)

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
		t.Fatal(err)
	}
	dbP, err := db.SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	dbGorm = dbP
}

func TestSetupRoutes(t *testing.T) {
  var err error
  router, err = SetupRoutes(dbGorm)
  assert.Nil(t, err, "Trying to set up routes shall not cause any errors")
}

func TestPingServer(t *testing.T) {
  w := httptest.NewRecorder()
  req, _ := http.NewRequest(http.MethodGet, "/api/hello", nil)
  router.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin1(t *testing.T) {
    c, w := setupContext()
		// Set Body, Header and Content-Type
		mockJSONPost(c, &gin.H{
			"Username": "admin",
      "Password": "pass",
		})
    

		// call API endpoint
		login(c)
    assert.Equal(t, 200, w.Code, "When trying to log in with correct admin credentials it shall return status: 200")
}

func TestLogin2(t *testing.T) {
    c, w := setupContext()
		// Set Body, Header and Content-Type
		mockJSONPost(c, &gin.H{
			"Username": "admin",
      "Password": "incorrect-password",
		})
    

		// call API endpoint
		login(c)
    assert.NotEqual(t, 200, w.Code, "When trying to log in with incorrect admin credentials it shall not return status: 200")
}


