package main

import (
	"github.com/gin-gonic/gin"
	
	"ByTeora-Pos-Backend-App/route"
)

func main() {
	r := gin.Default()

	route.RegisterRoutes(r)

	r.Run(":8080")
}