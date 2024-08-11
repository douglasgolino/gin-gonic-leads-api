package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	/*dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	LeadRepository := repository.NewLeadRepository(dbConnection)

	LeadUsecase := usecase.NewLeadUseCase(LeadRepository)*/

	server.Run(":8000")

}
