package main

import (
	"gin-gonic-leads-api/controller"
	"gin-gonic-leads-api/db"
	"gin-gonic-leads-api/repository"
	"gin-gonic-leads-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	LeadRepository := repository.NewLeadRepository(dbConnection)

	LeadUsecase := usecase.NewLeadUseCase(LeadRepository)

	LeadController := controller.NewLeadController(LeadUsecase)

	server.POST("/lead", LeadController.Create)

	server.Run(":8000")

}
