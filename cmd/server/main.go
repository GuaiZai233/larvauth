package main

import (
	"log"

	"github.com/GuaiZai233/larvauth/internal/auth"
	"github.com/GuaiZai233/larvauth/internal/db"
	"github.com/GuaiZai233/larvauth/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.GetDB()
	if err := database.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.SetupRouter(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
