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

type ResponseSosMed struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint64 `json:"user_id"`
	User           struct {
		ID              uint64 `json:"id"`
		Username        string `json:"username"`
		ProfileImageUrl string `json:"profile_image_url"`
	}
}

func CreateSosmed(c *gin.Context) {
	db := config.DbInit()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID
	err := db.Debug().Create(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.SocialMediaUrl,
		"user_id":          Sosmed.UserID,
		"created_at":       Sosmed.CreatedAt,
	})
}

func UpdateSosmed(c *gin.Context) {
	db := config.DbInit()
	contentType := helpers.GetContentType(c)
	var (
		Sosmed       models.SocialMedia
		SosmedUpdate models.SocialMedia
	)

	SosmedID := c.Param("socialMediaId")

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	err := db.Model(&Sosmed).Where("id = ?", SosmedID).Updates(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}
	db.Where("id = ?", SosmedID).First(&SosmedUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":               SosmedUpdate.ID,
		"name":             SosmedUpdate.Name,
		"social_media_url": SosmedUpdate.SocialMediaUrl,
		"user_id":          SosmedUpdate.UserID,
		"updated_at":       SosmedUpdate.UpdatedAt,
	})
}

func DeleteSosmed(c *gin.Context) {
	db := config.DbInit()
	Sosmed := models.SocialMedia{}
	SosmedID := c.Param("socialMediaId")

	err := db.Where("id = ?", SosmedID).Delete(&Sosmed).Error
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

func GetSosmed(c *gin.Context) {
	db := config.DbInit()
	var (
		sosmed         []models.SocialMedia
		responseSosmed []ResponseSosMed
	)

	err := db.Preload("User").Find(&sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	m := dto.Mapper{}
	m.Map(&responseSosmed, sosmed)

	c.JSON(http.StatusOK, responseSosmed)
}
