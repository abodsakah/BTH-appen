// Package routes provides routes
package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/abodsakah/BTH-appen/backend/api/src/JWTAuth"
)

// variables
var (
	dBase *gorm.DB
	err   error
)

// SetupRoutes function
func SetupRoutes() {
	dBase, err = db.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}
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
	// auth protected routes
	auth := r.Group("/", authMiddleware)
	auth.GET("/api/auth-hello", hello)
	auth.GET("/api/list-exam-users", listExamUsers) // TODO: should be adminMiddleware protected
	auth.POST("/api/create-user", createUser)       // TODO: should be adminMiddleware protected
	auth.POST("/api/create-exam", createExam)       // TODO: should be adminMiddleware protected
	auth.POST("/api/delete-exam", deleteExam)       // TODO: should be adminMiddleware protected
	auth.POST("/api/register-exam", registerToExam)
	auth.POST("/api/unregister-exam", unregisterFromExam)

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
		fmt.Println("user NOT logged in")
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	fmt.Println("Cookie:", cookie)

	id, err := jwtauth.ValidateJWT(cookie)
	if err != nil {
		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create exam
	if err := db.CreateExam(dBase, &exam); err != nil {
		fmt.Println(err.Error())
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exam)
}

func deleteExam(c *gin.Context) {
	var reqObj examReqBody

	// bind body data or return error if it fails
	if err := c.ShouldBind(&reqObj); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// delete exam
	if err := db.DeleteExam(dBase, reqObj.ExamID); err != nil {
		fmt.Println(err.Error())
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Success": fmt.Sprintf("Exam %d was deleted", reqObj.ExamID)})
}

func listExams(c *gin.Context) {
	// get exams from database
	exams, err := db.ListExams(dBase)
	if err != nil {
		fmt.Println(err.Error())
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
	users, err := db.ListExamUsers(dBase, reqObj.ExamID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, users)
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
	err := db.AddUserToExam(dBase, reqObj.ExamID, userID)
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
	err := db.RemoveUserFromExam(dBase, reqObj.ExamID, userID)
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
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": "Missing user credentials"})
		return
	}

	// Create user
	fmt.Printf("User: %#v\n", user)
	if err := db.CreateUser(dBase, &user); err != nil {
		fmt.Println(err.Error())
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"username": user.Username, "creation-date": user.CreatedAt})
}

func login(c *gin.Context) {
	var user db.User

	// bind body data or return error if it fails
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": "Missing user credentials"})
		return
	}

	// Try to authenticate user
	fmt.Printf("User: %#v\n", user)
	userID, err := db.AuthUser(dBase, user.Username, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(401, gin.H{"error": "Failed to authenticate user, username or password is wrong."})
		return
	}

	// generate a JWT for the user with user ID
	token, err := jwtauth.GenerateJWT(userID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// create a cookie that's valid for 2 hours to match the JWT 2 hour expiration time
	c.SetCookie("BTH-app", token, 60*60*2, "/", "localhost", true, true)

	fmt.Println("Token:", token)

	c.JSON(200, gin.H{"message": "Success"})
}
