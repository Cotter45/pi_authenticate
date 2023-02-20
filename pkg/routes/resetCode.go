package routes

import (
	"os"
	"fmt"
	"log"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"auth/model"
)

type ResetCodeRequest struct {
	Email string `json:"email" binding:"required"`
}

// Reset password code
func ResetCode(c *gin.Context) {
	appName := c.GetHeader("AppName")

	db := c.MustGet("db").(*gorm.DB)

	var request ResetCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	// find user
	var user model.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// create a random api key uuid string
	uid := uuid.New().String()

	// update user app key to new key
	db.Model(&user).Update("reset_token", uid)

	// send email with uid
	from := mail.NewEmail("Sean Cotter", "cotter.github45@gmail.com")
	subject := fmt.Sprintf("Reset password for %s", appName)
	to := mail.NewEmail(request.Email, request.Email)
	plainTextContent := fmt.Sprintf("Your reset code is: %s", uid)
	htmlContent := fmt.Sprintf("<div><p>Your reset code is: </p></div> <div><strong>%s</strong></div>", uid)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	c.JSON(201, gin.H{
		"message": "success",
	})
}