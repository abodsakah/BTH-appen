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
func SetupRoutes(gormObj *gorm.DB) {
	// setup GORM database object
	gormDB = gormObj

	r := gin.Default()
	// set trusted proxy to localhost
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalln(err)
	}

	// routes
	r.GET("/api/hello", hello)
	r.POST("/api/login", login)
	r.GET("/api/list-exams", listExams)
	r.Get("/api/list-news", listNews)
	r.GET("/api/list-due-exams", listDueExams)
	// auth protected routes
	auth := r.Group("/", authMiddleware)
	auth.GET("/api/auth-hello", hello)
	auth.GET("/api/list-exam-users", listExamUsers) // TODO: should be adminMiddleware protected
	auth.POST("/api/create-user", createUser)       // TODO: should be adminMiddleware protected
	auth.POST("/api/create-exam", createExam)       // TODO: should be adminMiddleware protected
	auth.POST("/api/delete-exam", deleteExam)       // TODO: should be adminMiddleware protected
	auth.POST("/api/register-exam", registerToExam)
	auth.POST("/api/unregister-exam", unregisterFromExam)
	auth.POST("/api/add-user-expo-push-token", addUserExpoPushToken)

	if err = r.Run(":5000"); err != nil {
		log.Fatalln(err)
	}
}

func hello(c *gin.Context) {
	UserID := c.Keys["UserID"]
	if UserID != nil {
		msg := gin.H{"message": fmt.Sprint("Hello there user ", UserID, "!")}
		c.IndentedJSON(http.StatusOK, msg)
		return
	}
	msg := gin.H{"message": "Hello there!"}
	c.IndentedJSON(http.StatusOK, msg)
}

func authMiddleware(c *gin.Context) {
	// check cookie for valid JWT to see if user is already logged in
	cookie, err := c.Cookie("BTH-app")
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	id, err := jwtauth.ValidateJWT(cookie)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
	// create UserID key in c.Keys
	c.Set("UserID", id)
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
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exam)
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
	if err := db.DeleteExam(gormDB, reqObj.ExamID); err != nil {
		log.Println(err.Error())
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Success": fmt.Sprintf("Exam %d was deleted", reqObj.ExamID)})
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
	err := db.AddExpoPushToken(gormDB, userID, expoToken.ExpoPushToken)
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
	err := db.AddUserToExam(gormDB, reqObj.ExamID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"Success": fmt.Sprintf("User %d was registered to exam %d", userID, reqObj.ExamID)})
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
	err := db.RemoveUserFromExam(gormDB, reqObj.ExamID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"Success": fmt.Sprintf("User %d was unregistered from exam %d", userID, reqObj.ExamID)})
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
		c.JSON(401, gin.H{"error": err.Error()})
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

	// create a cookie that's valid for 2 hours to match the JWT 2 hour expiration time
	c.SetCookie("BTH-app", token, 60*60*2, "/", "localhost", true, true)
	c.JSON(200, gin.H{"message": "Success"})
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
