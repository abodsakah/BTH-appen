// Package routes provides routes
package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	jwtauth "github.com/abodsakah/BTH-appen/backend/api/src/JWTAuth"
)

// variables
var (
	gormDB *gorm.DB
)

// SetupRoutes function
func SetupRoutes(gormObj *gorm.DB) (*gin.Engine, error) {
	// setup GORM database object
	gormDB = gormObj

	r := gin.Default()
	// set trusted proxy to localhost
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil, err
	}

	// routes
	r.GET("/api/hello", hello)
	r.POST("/api/login", login)
	r.GET("/api/list-exams", listExams)
	r.GET("/api/list-news", listNews)
	r.GET("/api/list-due-exams", listDueExams)
	// auth protected routes
	auth := r.Group("/", authMiddleware)
	{
		auth.GET("/api/auth-hello", hello)
		auth.GET("/api/list-user-exams", listUserExams)
		auth.POST("/api/register-exam", registerToExam)
		auth.POST("/api/unregister-exam", unregisterFromExam)
		auth.POST("/api/add-user-expo-push-token", addUserExpoPushToken)
		adminAuth := auth.Group("/", adminMiddleware)
		{
			adminAuth.GET("/api/list-exam-users", listExamUsers)
			adminAuth.POST("/api/create-user", createUser)
			adminAuth.POST("/api/create-exam", createExam)
			adminAuth.DELETE("/api/delete-exam", deleteExam)
			adminAuth.DELETE("/api/delete-news", deleteNews)
		}
	}
  return r, nil
}

func hello(c *gin.Context) {
	userID := c.Keys["UserID"]
	if userID != nil {
		msg := gin.H{"message": fmt.Sprint("Hello there user ", userID, "!")}
		c.IndentedJSON(http.StatusOK, msg)
		return
	}
	msg := gin.H{"message": "Hello there!"}
	c.IndentedJSON(http.StatusOK, msg)
}

func authMiddleware(c *gin.Context) {
	// check cookie for valid JWT to see if user is already logged in
	var id uint
	var h authReqBody
	cookieJwt, err := c.Cookie("BTH-app")
	if err != nil {
		// if no cookie is found, try to bind header.
		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		// if we could bind header
		id, err = jwtauth.ValidateJWT(h.Jwt)
	} else {
		// if cookie exists
		id, err = jwtauth.ValidateJWT(cookieJwt)
	}
	// check error from ValidateJWT
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
	// create UserID key in c.Keys
	c.Set("UserID", id)
}

func adminMiddleware(c *gin.Context) {
	id := c.Keys["UserID"].(uint)
	isAdmin, err := db.IsRole(gormDB, id, "admin")
	if err != nil || !isAdmin {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
}

func createExam(c *gin.Context) {
	var exam db.Exam

	// bind body data or return error if it fails
	if err := c.ShouldBind(&exam); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create exam
	if err := db.CreateExam(gormDB, &exam); err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exam)
}

func deleteNews(c *gin.Context) {
	var reqObj newsReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// delete news article
	news, err := db.DeleteNews(gormDB, reqObj.NewsID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Success": fmt.Sprintf("News article %d was deleted", reqObj.NewsID),
		"news":    news,
	})
}

func deleteExam(c *gin.Context) {
	var reqObj examReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// delete exam
	exam, err := db.DeleteExam(gormDB, reqObj.ExamID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Success": fmt.Sprintf("Exam %d was deleted", reqObj.ExamID),
		"exam":    exam,
	})
}

func listExams(c *gin.Context) {
	// get exams from database
	exams, err := db.GetExams(gormDB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, exams)
}

func listDueExams(c *gin.Context) {
	// get exams from database
	exams, err := db.GetExamsDueSoon(gormDB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, exams)
}

// list exams a user is registered to
func listUserExams(c *gin.Context) {
	userID := c.Keys["UserID"].(uint)
	// get users registered exams from database
	users, err := db.GetUserExams(gormDB, userID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, users)
}

// list a exams registered users
func listExamUsers(c *gin.Context) {
	var reqObj examReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		c.JSON(400, gin.H{"error": "No exam ID provided"})
		return
	}

	// get users from database
	users, err := db.GetExamUsers(gormDB, reqObj.ExamID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, users)
}

func addUserExpoPushToken(c *gin.Context) {
	var expoToken db.Token

	// bind body data or return error if it fails
	if err := c.ShouldBind(&expoToken); err != nil {
		c.JSON(400, gin.H{"error": "No Expo token provided"})
		return
	}
	// add expo push token to user
	userID := c.Keys["UserID"].(uint)
	_, err := db.AddExpoPushToken(gormDB, userID, expoToken.ExpoPushToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"Success": fmt.Sprintf("Token: %s was added to user %d", expoToken.ExpoPushToken, userID)})
}

func registerToExam(c *gin.Context) {
	var reqObj examReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		c.JSON(400, gin.H{"error": "No exam ID provided"})
		return
	}
	// register user to exam
	userID := c.Keys["UserID"].(uint)
	user, err := db.AddUserToExam(gormDB, reqObj.ExamID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"Success": fmt.Sprintf("User %d was registered to exam %d", userID, reqObj.ExamID),
		"user":    user,
	})
}

func unregisterFromExam(c *gin.Context) {
	var reqObj examReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		c.JSON(400, gin.H{"error": "No exam ID provided"})
		return
	}
	// unregister user to exam
	userID := c.Keys["UserID"].(uint)
	user, err := db.RemoveUserFromExam(gormDB, reqObj.ExamID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"Success": fmt.Sprintf("User %d was unregistered from exam %d", userID, reqObj.ExamID),
		"user":    user,
	})
}

func createUser(c *gin.Context) {
	var user db.User
 	// bind body data or return error if it fails
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "Missing user credentials"})
		return
	}

	// Create user
	if err := db.CreateUser(gormDB, &user); err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"username": user.Username, "creation-date": user.CreatedAt})
}

func login(c *gin.Context) {
	var user db.User
	// bind body data or return error if it fails
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "Missing user credentials"})
		return
	}

	// Try to authenticate user
	userID, err := db.AuthUser(gormDB, user.Username, user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(401, gin.H{"error": "Failed to authenticate user, username or password is wrong."})
		return
	}

	// generate a JWT for the user with user ID
	token, err := jwtauth.GenerateJWT(userID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// create JSON to send to client.
	userInfo, err := db.GetUser(gormDB, userID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	jsonUserInfo := gin.H{
		"status": "success",
		"jwt":    token,
		"user":   userInfo,
	}

	// create a cookie that's valid for 2 hours to match the JWT 2 hour expiration time
	c.SetCookie("BTH-app", token, 60*60*2, "/", "localhost", true, true)
	c.JSON(200, jsonUserInfo)
}

func listNews(c *gin.Context) {
	// get news from database
	news, err := db.GetNews(gormDB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, news)
}
