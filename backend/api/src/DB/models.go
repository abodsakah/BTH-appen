package db

import (
	"time"

	"gorm.io/gorm"
)

// Database Models

// User struct
type User struct {
	gorm.Model
	Name     string  `form:"name" json:"name,omitempty"`
	Username string  `gorm:"uniqueIndex" form:"username" binding:"required" json:"username,omitempty"`
	Password string  `form:"password" binding:"required" json:"password,omitempty"`
	Role     string  `form:"role" gorm:"default:student" json:"role,omitempty"`
	Exams    []*Exam `gorm:"many2many:exam_users;" json:"exams,omitempty"`
	Tokens   []Token `json:"tokens,omitempty"`
}

// Exam struct
type Exam struct {
	gorm.Model
	Name       string    `form:"name" binding:"required" json:"name,omitempty"`
	CourseCode string    `form:"course_code" binding:"required" json:"course_code,omitempty"`
	Room       string    `form:"room" binding:"required" json:"room,omitempty"`
	StartDate  time.Time `form:"start_date" binding:"required" json:"start_date,omitempty"`
	Users      []*User   `gorm:"many2many:exam_users;" json:"users,omitempty"`
}

// News struct
type News struct {
	gorm.Model
	Title       string    `gorm:"uniqueIndex" form:"title" binding:"required" json:"title,omitempty"`
	Date        time.Time `form:"date" binding:"required" json:"date,omitempty"`
	Description string    `form:"description" binding:"required" json:"description,omitempty"`
	Link        string    `form:"link" binding:"required" json:"link,omitempty"`
}

// Token struct
type Token struct {
	gorm.Model
	ExpoPushToken string `gorm:"uniqueIndex" form:"expo_push_token" binding:"required" json:"expo_push_token,omitempty"`
	UserID        uint   `json:"user_id,omitempty"`
}
