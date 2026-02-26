package main

import (
    "exercise-api-for-pj1/presentation"
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	    r := gin.Default()
	r.Use(cors.New(cors.Config{
		// React アプリのホストURLを指定
		AllowOrigins:     []string{"http://localhost:3030"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    // GETリクエストで名前と時刻を受け取る
    r.GET("/api/greet", presentation.GreetingHandler)
	// POSTリクエストで名前と時刻を受け取る
	//r.POST("/api/greet", presentation.GreetingHandler)
    r.Run(":8081")
}
