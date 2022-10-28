package controllers

import (
	"final-project/config"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
)

type ResponseComments struct {
	ID      uint64 `json:"id"`
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message"`
	User    struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"User"`
	Photo struct {
		ID       uint64 `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   uint64 `json:"user_id"`
	}
}

func CreateComment(c *gin.Context) {
	db := config.DbInit()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func UpdateComment(c *gin.Context) {
	db := config.DbInit()
	contentType := helpers.GetContentType(c)
	var (
		Comment       models.Comment
		CommentUpdate models.Comment
		// Photo         models.Photo
	)

	CommentID := c.Param("commentId")

	if contentType == appJSON {
		c.ShouldBindJSON(&CommentUpdate)
	} else {
		c.ShouldBind(&CommentUpdate)
	}

	err := db.Model(&CommentUpdate).Where("id = ?", CommentID).Updates(&CommentUpdate).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	db.Where("id = ?", CommentID).First(&Comment).Preload("User").Preload("Photo").Find(&Comment)

	// c.JSON(http.StatusOK, Comment)
	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"title":      Comment.Photo.Title,
		"caption":    Comment.Photo.Caption,
		"photo_url":  Comment.Photo.PhotoURL,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := config.DbInit()
	Comment := models.Comment{}
	CommentID := c.Param("commentId")

	err := db.Where("id = ?", CommentID).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been succesfully deleted",
	})
}

func GetComments(c *gin.Context) {
	db := config.DbInit()
	var (
		comments         []models.Comment
		responseComments []ResponseComments
	)

	err := db.Preload("User").Preload("Photo").Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	m := dto.Mapper{}
	m.Map(&responseComments, comments)

	c.JSON(http.StatusOK, responseComments)
}
