
package presentation

import (
	"time"
	"github.com/gin-gonic/gin"
	"exercise-api-for-pj1/domain"
	"exercise-api-for-pj1/usecase"	
)
    //APIのレスポンス構造体
    type GreetingResponse struct {
        Code int `json:"code"`
        Message string `json:"message"`
        Response any `json:"response"`
    }

// presentation/handler.go
func GreetingHandler(c *gin.Context) {
    // 別のポートからのアクセスを許可 (CORS)
    c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3030")

    // GETのURLパラメータを構造体にマッピング
    var req struct {
        Name string    `form:"name"`
        Time time.Time `form:"time" time_format:"2006-01-02T15:04:05.000Z"`
    }
	//POSTのbodyを構造体にマッピング
	// var req struct{
	// 	Name string  `json:"name"`
	// 	Time time.Time `json:"time" time_format:"2006-01-02T15:04:05.000Z"`
	// 	}

	//GETのリクエスト不正精査
    if err := c.ShouldBindQuery(&req); err != nil {
        c.JSON(400, GreetingResponse{
            Code:    400,
            Message: "パラメータが不正です",
            Response: nil,
        })
        return
    }

	//POSTのbody部分のリクエスト不正精査
	//  if err := c.ShouldBindJSON(&req); err != nil {
    //     c.JSON(400, GreetingResponse{
    //         Code:    400,
    //         Message: "パラメータが不正です",
    //         Response: nil,
    //     })
    //     return
    // }

    // Domain層のInputに入れ替える
    input := domain.GreetingInput{
        Name: req.Name,
        Time: req.Time,
    }

    //入力値エラー
    if err := input.Validate(); err != nil {
        c.JSON(400, GreetingResponse{
            Code:     400,
            Message:  err.Error(),
            Response: nil,
        })
        return
    }

    output, err := usecase.ExecuteGreeting(input)
    
    if err !=nil {
        c.JSON(500,GreetingResponse{
            Code: 500,
            Message: "サーバーエラー",
            Response: nil,
        })
        return
    }

    c.JSON(200, GreetingResponse{
        Code:     200,
        Message:  "Success", 
        Response: output,    
    })

}
