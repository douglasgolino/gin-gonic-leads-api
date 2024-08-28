package main

import (
	"gin-gonic-leads-api/controller"
	"gin-gonic-leads-api/db"
	"gin-gonic-leads-api/repository"
	"gin-gonic-leads-api/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Configuração do CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},            // Permitir o frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},     // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Cabeçalhos permitidos
		AllowCredentials: true,                                         // Permitir cookies, autenticação e outras credenciais
	}))

	// Conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Instanciação do repositório, caso de uso e controlador
	LeadRepository := repository.NewLeadRepository(dbConnection)
	LeadUsecase := usecase.NewLeadUseCase(LeadRepository)
	LeadController := controller.NewLeadController(LeadUsecase)

	// Configuração das rotas
	server.POST("/lead", LeadController.Create)

	// Inicialização do servidor
	server.Run(":8000")
}
