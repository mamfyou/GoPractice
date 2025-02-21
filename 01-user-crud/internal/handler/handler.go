package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"mamfyou/internal/handler/user"
)

const (
	host = "localhost"
	port = "8080"
)

func Start() {
	r := gin.Default()
	registerRoutes(r)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}

func registerRoutes(r *gin.Engine) {
	user.RegisterRoutes(r)
}
