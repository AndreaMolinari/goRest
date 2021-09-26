package main

import (
	// "gorm-test/controllers"
	"aMolinariCom/goRest/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.User()
	anagraficaRepo := controllers.Anagrafica()

	r.POST("/login", userRepo.Login)

	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	r.GET("/anagrafica", anagraficaRepo.GetAnagrafiche)
	r.POST("/anagrafica", anagraficaRepo.CreateAnagrafica)
	r.PUT("/anagrafica/:id", anagraficaRepo.UpdateAnagrafica)
	r.GET("/anagrafica/:id", anagraficaRepo.GetAnagrafica)
	r.DELETE("/anagrafica/:id", anagraficaRepo.DeleteAnagrafica)

	return r
}
