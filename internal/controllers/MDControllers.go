package controllers

import (
	"godistributor/internal/entity"
	"net/http"

"github.com/gin-gonic/gin"
"github.com/jinzhu/gorm"
)

type CreateDistribInput struct {
	NameMD				string 	`json:"nameMD"`
	AddressMD			string	`json:"addressMD"`
	LocationMD			string	`json:"locationMD"`
	StockMD				int		`json:"stockMD"`
}

type UpdateDistribInput struct {
	NameMD				string 	`json:"nameMD"`
	AddressMD			string	`json:"addressMD"`
	LocationMD			string	`json:"locationMD"`
	StockMD				int		`json:"stockMD"`
}

// GET /distribs
// Get all distribs
func GetAllDistribs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mainDistribs []entity.MainDistributor
	db.Find(&mainDistribs)

	c.JSON(http.StatusOK, gin.H{"data": mainDistribs})
}

// POST /distribs
// Create new distrib
func CreateDistrib(c *gin.Context) {
	// Validate input
	var input CreateDistribInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	mainDistrib := entity.MainDistributor {
		NameMD: input.NameMD, AddressMD: input.AddressMD,
		LocationMD: input.LocationMD, StockMD: input.StockMD,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&mainDistrib)

	c.JSON(http.StatusOK, gin.H{"data": mainDistrib})
}

// GET /distribs/:id
// Find a distribs
func FindDistribs(c *gin.Context) { // Get model if exist
	var mainDistrib entity.MainDistributor

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id_main_distributor = ?", c.Param("id")).First(&mainDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mainDistrib})
}

// PATCH /distribs/:id
// Update a Distrib
func UpdateDistrib(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var mainDistrib entity.MainDistributor
	if err := db.Where("id_main_distributor = ?", c.Param("id")).First(&mainDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateDistribInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	var updatedMainDistrib entity.MainDistributor
	updatedMainDistrib.NameMD = input.NameMD
	updatedMainDistrib.AddressMD = input.AddressMD
	updatedMainDistrib.LocationMD = input.LocationMD
	updatedMainDistrib.StockMD = input.StockMD

	db.Model(&mainDistrib).Updates(updatedMainDistrib)

	c.JSON(http.StatusOK, gin.H{"data": mainDistrib})
}

// DELETE /distribs/:id
// Delete a distrib
func DeleteDistrib(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var mainDistrib entity.MainDistributor
	if err := db.Where("id_main_distributor = ?", c.Param("id")).First(&mainDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&mainDistrib)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
