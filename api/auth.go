package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vigneshganesan008/notes-api/dao"
	"github.com/vigneshganesan008/notes-api/models"
	"github.com/vigneshganesan008/notes-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var signupReq models.User
	if err := c.Bind(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload"})
		return
	}

	if hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(signupReq.Password),
		bcrypt.DefaultCost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	} else {
		signupReq.Password = string(hashedPassword)
	}

	id, err := dao.InsertUser(signupReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func Login(c *gin.Context) {
	var loginReq models.User
	if err := c.Bind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload"})
		return
	}

	user, err := dao.GetUser(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "passowrd wrong"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%+v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": "bearer" + token})
}
