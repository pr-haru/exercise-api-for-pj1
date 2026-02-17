package main

import (
	// "net/http"
	// "time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
        // すべてのドメインからのアクセスを許可（開発用として最も確実な設定）
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")

        // プリフライトリクエスト(OPTIONS)への対応
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })
	// Load the HTML template file from the presentation directory
	r.LoadHTMLFiles("presentation/index.html")

	r.GET("/presentation/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":3030")
}
