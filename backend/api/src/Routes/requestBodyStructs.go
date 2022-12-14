// Package routes provides routes
package routes

type examReqBody struct {
	ExamID uint `form:"exam_id" binding:"required" json:"exam_id"`
}

type newsReqBody struct {
	NewsID uint `form:"news_id" binding:"required" json:"news_id"`
}

type authReqBody struct {
	Jwt string `header:"jwt" form:"jwt" binding:"required" json:"jwt"`
}
