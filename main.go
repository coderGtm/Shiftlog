package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	setUpRoutes(router)

	router.Run(":8000")
}