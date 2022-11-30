package main

import (
	"github.com/MaeMethas/sa-65-example/controller"
	"github.com/MaeMethas/sa-65-example/entity"
	"github.com/MaeMethas/sa-65-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("/")

	{
		router.Use(middlewares.Authorizes())
		{
			r.GET("/Student", controller.ListStudent)
			r.GET("/Student/:id", controller.GetStudent)
			r.PATCH("/Student", controller.UpdateStudent)
			r.DELETE("/Student/:id", controller.DeleteStudent)

			// Subject Routes

			r.GET("/subjects", controller.ListSubject)
			r.GET("/subjects/:id", controller.GetSubject)
			r.PATCH("/subjects", controller.UpdateSubject)
			r.DELETE("/subjects/:id", controller.DeleteSubject)

			// Group Routes

			r.GET("/states", controller.ListState)
			r.GET("/states/:id", controller.GetState)
			r.PATCH("/states", controller.UpdateState)
			r.DELETE("/states/:id", controller.DeleteState)

			//Registration Routes

			r.GET("/registrations", controller.ListRegistration)
			r.GET("/registrations/:id", controller.GetRegistration)
			r.POST("/registrations", controller.CreateRegistration)
			r.PATCH("/registrations", controller.UpdateRegistration)
			r.DELETE("/registrations/:id", controller.DeleteRegistration)
		}

	}

	r.POST("/signup", controller.CreateStudent)
	// login User Route
	r.POST("/login", controller.Login)

	// // Run the server go run main.go
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}

}
