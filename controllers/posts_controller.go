package controllers

import (
	"errors"
	"github/zaulgin/json_crud_api/initializers"
	"github/zaulgin/json_crud_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type body struct {
	Title string
	Body  string
}

func PostsCreate(c *gin.Context) {
	// get data oof req body
	var b body

	c.Bind(&b)

	// create a post
	post := models.Post{Title: b.Title, Body: b.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "database error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	var b body

	c.Bind(&b)

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "database error",
			})
		}
		return
	}

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: b.Title,
		Body:  b.Body,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "database error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
