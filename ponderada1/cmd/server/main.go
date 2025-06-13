package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"ponderada1/config"
	_ "ponderada1/docs"
	"ponderada1/internal/db"
	"ponderada1/internal/handler"
	"ponderada1/internal/model"
	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Weather API Template
// @version     1.0
// @description Exemplo de API em Go + Gin com Swagger e integração externa.
// @contact.name Rodrigo

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http
func main() {
	config.LoadEnv() // lê variáveis de ambiente

	db := db.GetDB()
	_ = db.AutoMigrate(&model.Singleplayer{}, &model.Multiplayer{})

	r := gin.Default()

	r.Use(cors.Default())

	// Documentação
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas de domínio
	v1 := r.Group("/api/v1")

	handler.RegisterDeckRoutes(v1)

	handler.RegisterSinglePlayerGameRoutes(v1)

	handler.RegisterMultiplayerGameRoutes(v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server listening on %s", port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal(err)
	}
}
