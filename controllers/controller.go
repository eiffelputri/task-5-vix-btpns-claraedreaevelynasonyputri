package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/helpers"
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context, db *gorm.DB) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving it
	hashedPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	fmt.Println("Hashed Password during registration:", hashedPassword)

	newUser.Password = hashedPassword

	// Save the user in the database
	db.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
// LoginUser handles user login
func LoginUser(c *gin.Context, db *gorm.DB) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Email"})
		return
	}

	fmt.Println("Stored Hashed Password:", user.Password) // Debug print

	// Check password
	if !helpers.CheckPasswordHash(loginData.Password, user.Password) {
		fmt.Println("Password comparison failed") // Debug print
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	fmt.Println("Password comparison successful") // Debug print

	// Rest of the code

	// Generate JWT token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// UpdateUser handles user profile update
func UpdateUser(c *gin.Context, db *gorm.DB) {
	userID := c.Param("userID")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user profile in the database
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	// ... Update other fields as needed
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}

// DeleteUser handles user account deletion
func DeleteUser(c *gin.Context, db *gorm.DB) {
	userID := c.Param("userID")

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User account deleted successfully"})
}

// UploadPhoto handles photo upload
func UploadPhoto(c *gin.Context, db *gorm.DB) {
	userIDStr := c.GetString("userID") // Retrieve userID from middleware

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPhoto.UserID = uint(userID)

	// Save the photo in the database
	db.Create(&newPhoto)

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully"})
}

// UpdatePhoto handles photo update
func UpdatePhoto(c *gin.Context, db *gorm.DB) {
	userID := c.GetString("userID") // Retrieve userID from middleware
	photoID := c.Param("photoID")

	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var photo models.Photo
	if err := db.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if fmt.Sprint(photo.UserID) != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You're not allowed to update this photo"})
		return
	}

	// Update photo in the database
	photo.Title = updatedPhoto.Title
	photo.Caption = updatedPhoto.Caption
	// ... Update other fields as needed
	db.Save(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto handles photo deletion
func DeletePhoto(c *gin.Context, db *gorm.DB) {
	userIDStr := c.GetString("userID") // Retrieve userID from middleware

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	photoID := c.Param("photoID")

	var photo models.Photo
	if err := db.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if photo.UserID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You're not allowed to delete this photo"})
		return
	}

	// Delete the photo from the database
	db.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
