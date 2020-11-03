package main

import (
	"github.com/andsholinka/Digitalent-Kominfo_Antrian-Bank-v1/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}
