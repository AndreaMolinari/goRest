package controllers

import (
	"aMolinariCom/goRest/database"
	"aMolinariCom/goRest/models"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnagraficaRepo struct {
	Db *gorm.DB
}

func Anagrafica() *AnagraficaRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Anagrafica{})
	return &AnagraficaRepo{Db: db}
}

//create user
func (repository *AnagraficaRepo) CreateAnagrafica(c *gin.Context) {
	var anagrafica models.Anagrafica
	c.BindJSON(&anagrafica)
	err := models.CreateAnagrafica(repository.Db, &anagrafica)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, anagrafica)
}

//get users
func (repository *AnagraficaRepo) GetAnagrafiche(c *gin.Context) {
	var anagrafiche []models.Anagrafica
	err := models.GetAnagrafiche(repository.Db, &anagrafiche)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, anagrafiche)
}

//get user by id
func (repository *AnagraficaRepo) GetAnagrafica(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var anagrafica models.Anagrafica
	err := models.GetAnagrafica(repository.Db, &anagrafica, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, anagrafica)
}

// update user
func (repository *AnagraficaRepo) UpdateAnagrafica(c *gin.Context) {
	var anagrafica models.Anagrafica
	id, _ := c.Params.Get("id")
	err := models.GetAnagrafica(repository.Db, &anagrafica, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&anagrafica)
	err = models.UpdateAnagrafica(repository.Db, &anagrafica)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, anagrafica)
}

// delete user
func (repository *AnagraficaRepo) DeleteAnagrafica(c *gin.Context) {
	var anagrafica models.Anagrafica
	id, _ := c.Params.Get("id")
	err := models.DeleteAnagrafica(repository.Db, &anagrafica, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
