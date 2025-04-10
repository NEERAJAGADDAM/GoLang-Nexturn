package main

import (
	"log"
	"os"
	"restcrudapi/Internals/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		frontendURL := os.Getenv("FRONTEND_URL")
		if frontendURL == "" {
			frontendURL = "http://localhost:5173"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", frontendURL)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(" No .env file found. Using system environment variables.")
	}
	services.InitDB()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Static("/views", "./views")

	r.GET("/", serveHTML("./views/index.html"))
	r.GET("/create", serveHTML("./views/create.html"))
	r.GET("/update", serveHTML("./views/update.html"))
	r.GET("/delete", serveHTML("./views/delete.html"))
	r.GET("/getbyid", serveHTML("./views/getbyid.html"))
	r.GET("/getall", serveHTML("./views/getall.html"))

	// Routes
	r.POST("/add_user", services.AddUser)
	r.GET("/users", services.GetAllUsers)
	r.GET("/get_all_users", services.GetAllUsers)
	r.GET("/get_user", services.GetUserByID)
	r.PUT("/update_user", services.UpdateUser)
	r.POST("/update_user", services.UpdateUser)
	r.DELETE("/delete_user", services.DeleteUser)
	r.POST("/delete_user", services.DeleteUser)

	log.Println("Server running on port 9000...")
	if err := r.Run(":9000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func serveHTML(filepath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File(filepath)
	}
}
