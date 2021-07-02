package controllers

import (
	"net/http"

	"github.com/bayuiqballl/demo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InputPeople struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Create(c *gin.Context) {
	var input InputPeople
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	insert := models.People{
		Name: input.Name,
		Age:  input.Age,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&insert)

	c.JSON(http.StatusOK, gin.H{"data": insert})

}

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var people []models.People
	db.Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})

}

func GetByID(c *gin.Context) {
	var people models.People
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": people})

}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var people models.People

	if err := db.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
	}

	db.Delete(&people)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var people models.People

	if err := db.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
	}

	var input InputPeople
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var update models.People
	update.Name = input.Name
	update.Age = input.Age

	db.Model(&people).Update(&update)

	c.JSON(http.StatusOK, gin.H{"data": update})

}
