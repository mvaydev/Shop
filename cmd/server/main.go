package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	pongHandler "github.com/mvaydev/shop/internal/handlers"
)

func main() {
	_ = run()
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	r := gin.Default()

	pong := r.Group("/ping")
	{
		pong.GET("", pongHandler.Get)
	}

	return r.Run(os.Getenv("APP_ADDR"))
}
