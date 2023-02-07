package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Girilaxman000/auth_go/database"
	"github.com/Girilaxman000/auth_go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// c is a pointer to an instance of the gin.Context struct. Context is a struct that contains the request and response data for a request.
func SignUp(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil { //reference to varibale
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	//handle an error

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{})
}

func SignIn(c *gin.Context) {
	//get email and password from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil { //reference to varibale
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//retrieve from database
	user := models.User{}
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password doesn't match",
		})
		return
	}
	//genereate a jwt token

	//alternative to pass not only id but others to sub here practive for own.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get the complete encoded token as string using secret key

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed To Create Token",
		})
		return
	}
	//send it back

	//reason behing setting it in a cookie and sending it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
