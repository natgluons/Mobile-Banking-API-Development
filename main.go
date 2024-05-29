package main

import (
	"github.com/dzikrifazahk/task-5-pbi-btpns-DzikriFazaHaunaKusnadi/controllers/photoscontroller"
	"github.com/dzikrifazahk/task-5-pbi-btpns-DzikriFazaHaunaKusnadi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	r.GET("/api/photos", photoscontroller.Index)
	r.POST("/api/photos", photoscontroller.Create)
	r.PUT("/api/photos/:id", photoscontroller.Update)
	r.DELETE("/api/photos", photoscontroller.Delete)

	r.Run()

}
