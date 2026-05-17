package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh-21-chauhan/go-auth/helpers"
	"github.com/harsh-21-chauhan/go-auth/config"
	"github.com/harsh-21-chauhan/go-auth/models"
	"golang.org/x/crypto/bcrypt"

)
type SignupInput struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context){

	var input SignupInput

   // Parse body
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  // Hash password
  hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
  if err != nil {
    c.JSON(500, gin.H{"error": "could not hash password"})
    return
  }

  //  Create user in db
  user := models.User{
    Email:    input.Email,
    Password: string(hashed),
    Role:     "user",
  }
  if result := config.DB.Create(&user); result.Error != nil {
    c.JSON(400, gin.H{"error": "email already exists"})
    return
  }

  c.JSON(201, gin.H{"message": "user created", "user": user})

}
func Login(c *gin.Context) {
  var input LoginInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  // find by email
  var user models.User
  if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
    c.JSON(401, gin.H{"error": "invalid credentials"})
    return
  }

  // compare password with hash
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
    c.JSON(401, gin.H{"error": "invalid credentials"})
    return
  }

  // generate jwt
  token, err := helpers.GenerateToken(user.ID, user.Role)
  if err != nil {
    c.JSON(500, gin.H{"error": "could not generate token"})
    return
  }

  c.JSON(200, gin.H{"token": token})
}


func GetAllUsers(c *gin.Context) {
  role, _ := c.Get("role")
  if role.(string) != "admin" {
    c.JSON(403, gin.H{"error": "forbidden: admin access only"})
    return
  }

  var users []models.User
  config.DB.Find(&users)
  c.JSON(200, gin.H{"users": users, "count": len(users)})
}

func GetProfile(c *gin.Context) {
 
  userIDRaw, _ := c.Get("userID")
  userID := uint(userIDRaw.(float64))  

  var user models.User
  if err := config.DB.First(&user, userID).Error; err != nil {
    c.JSON(404, gin.H{"error": "user not found"})
    return
  }

  c.JSON(200, gin.H{"user": user})
  
}
