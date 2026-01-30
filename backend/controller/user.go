package controller

import (
	"log"
	"net/http"

	"github.com/ApexPlayground/Linkkit/model"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/ApexPlayground/Linkkit/util"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body model.User

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse user body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	//call service
	user, err := service.SignUp(body)
	if err != nil {
		log.Println("Failed to create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"is_admin":  user.IsAdmin,
		"createdAt": user.CreatedAt,
	})

}

func Login(c *gin.Context) {
	// get the email and password from request body
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse user body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user, err := service.Login(body.Email, body.Password)
	if err != nil {
		log.Println("Login Failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := util.GenerateJWT(user)
	if err != nil {
		log.Println("Failed to generate token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"is_admin":  user.IsAdmin,
		"token":     token,
		"createdAt": user.CreatedAt,
	})

}

func GetUser(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDValue.(uint)

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context) {

	isAdminValue, exists := c.Get("is_admin")
	if !exists || !isAdminValue.(bool) {
		c.JSON(403, gin.H{"error": "forbidden: admin only"})
		return
	}

	// Admin verified, fetch all users
	users, err := service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UpdateUser(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDValue.(uint)

	var body service.UpdateUserInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedUser, err := service.UpdateUser(userID, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        updatedUser.ID,
		"name":      updatedUser.Name,
		"email":     updatedUser.Email,
		"is_admin":  updatedUser.IsAdmin,
		"updatedAt": updatedUser.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDValue.(uint)

	err := service.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
