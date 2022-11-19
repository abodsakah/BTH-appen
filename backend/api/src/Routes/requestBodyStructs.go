// Package routes provides routes
package routes

type regExamBody struct {
	ExamID uint `form:"exam_id" binding:"required" json:"exam_id"`
}
