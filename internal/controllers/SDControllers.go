package controllers

import (
	"godistributor/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateSecDistribInput struct {
	NameSD     string `json:"nameSD"`
	AddressSD  string `json:"addressSD"`
	LocationSD string `json:"locationSD"`
	StockSD    int    `json:"stockSD"`
	IdMD       int
}

type UpdateSecDistribInput struct {
	NameSD     string `json:"nameSD"`
	AddressSD  string `json:"addressSD"`
	LocationSD string `json:"locationSD"`
	StockSD    int    `json:"stockSD"`
	IdMD       int
}

// GET /distribs
// Get all distribs
func GetAllSecDistribs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var secDistribs []entity.SecDistributor
	db.Find(&secDistribs)

	c.JSON(http.StatusOK, gin.H{"data": secDistribs})
}

// POST /distribs
// Create new distrib
func CreateSecDistrib(c *gin.Context) {
	// Validate input
	var input CreateSecDistribInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	secDistrib := entity.SecDistributor{
		NameSD: input.NameSD, AddressSD: input.AddressSD,
		LocationSD: input.LocationSD, StockSD: input.StockSD, IdMD: input.IdMD,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&secDistrib)

	c.JSON(http.StatusOK, gin.H{"data": secDistrib})
}

// GET /distribs/:id
// Find a sdistribs
func FindSecDistribs(c *gin.Context) { // Get model if exist
	var secDistrib entity.SecDistributor

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id_sec_distributor = ?", c.Param("id")).First(&secDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": secDistrib})
}

// PATCH /distribs/:id
// Update a Distrib
func UpdateSecDistrib(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var secDistrib entity.SecDistributor
	if err := db.Where("id_sec_distributor = ?", c.Param("id")).First(&secDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateSecDistribInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	var updatedSecDistrib entity.SecDistributor
	updatedSecDistrib.NameSD = input.NameSD
	updatedSecDistrib.AddressSD = input.AddressSD
	updatedSecDistrib.LocationSD = input.LocationSD
	updatedSecDistrib.StockSD = input.StockSD

	db.Model(&secDistrib).Updates(updatedSecDistrib)

	c.JSON(http.StatusOK, gin.H{"data": secDistrib})
}

// DELETE /distribs/:id
// Delete a distrib
func DeleteSecDistrib(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var secDistrib entity.SecDistributor
	if err := db.Where("id_sec_distributor = ?", c.Param("id")).First(&secDistrib).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&secDistrib)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
