package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/ARMeeru/time-capsule/config"
	"github.com/ARMeeru/time-capsule/models"
	"github.com/ARMeeru/time-capsule/utils"
)

func ShowRegisterPage(c *gin.Context) {
	utils.RenderTemplate(c, []string{"templates/base.html", "templates/register.html"}, nil)
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/register.html"}, gin.H{"Error": "Failed to register"})
		return
	}

	// Create user
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/register.html"}, gin.H{
			"Error": "Email already registered",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func ShowLoginPage(c *gin.Context) {
	utils.RenderTemplate(c, []string{"templates/base.html", "templates/login.html"}, nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/login.html"}, gin.H{"Error": "Invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/login.html"}, gin.H{
			"Error": "Invalid credentials",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/login.html"}, gin.H{"Error": "Failed to generate token"})
		return
	}

	// Set token as a cookie
	c.SetCookie("token", tokenString, 3600*24*3, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/capsules/new")
}

func Logout(c *gin.Context) {
	// Clear the token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
