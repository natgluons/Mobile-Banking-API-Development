package photoscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dzikrifazahk/task-5-pbi-btpns-DzikriFazaHaunaKusnadi/database"
	"github.com/dzikrifazahk/task-5-pbi-btpns-DzikriFazaHaunaKusnadi/models"
)

// show data photo
func Index(c *gin.Context) {
	var photos []models.Photos

	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

// tambah data photo
func Create(c *gin.Context) {
	var photos models.Photos

	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})

}

// update data photo
func Update(c *gin.Context) {
	var photos models.Photos
	id := c.Param("id")

	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&photos).Where("id = ?", id).Updates(&photos).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

// hapus data photo
func Delete(c *gin.Context) {
	var photos models.Photos

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if database.DB.Delete(&photos, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}
