package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type dbMock struct {
	mock.Mock
}

func (m *dbMock) Where(username string, password string) *gorm.DB {
	var result *gorm.DB
	result.Error = nil
	return result
}

func TestMock_id1(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	db = new(dbMock)
	AuthUser("test", "test")
	assert.Equal(a, b, "The two ")

}
