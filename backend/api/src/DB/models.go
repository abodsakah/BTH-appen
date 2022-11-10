package db

import (
	"time"

	"gorm.io/gorm"
)

// Database Models

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
}

// Exam struct
type Exam struct {
	gorm.Model
	CourseCode string    `gorm:"uniqueIndex" form:"course_code" binding:"required" json:"course_code"`
	StartDate  time.Time `form:"start_date" binding:"required" json:"start_date"`
	Users      []User    `gorm:"many2many:exam_users;"`
}

// News struct
type News struct {
	gorm.Model
	Title string `form:"title" binding:"required" json:"title"`
	Body  string `form:"body" binding:"required" json:"body"`
}
