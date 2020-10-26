package main

import (
	"github.com/gin-gonic/gin"
	"http/controller"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware())

	engine.GET("/ws", websocketVerification(), controller.WebSocketHandler)
	engine.POST("/recharge/:action", controller.Lobby, verification())
	engine.POST("/upload/:type", controller.Upload)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: engine,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Printf("Before : %s", time.Now().Format("2006-01-02 15:04:05.000000000"))

		context.Next()

		log.Printf("After  : %s", time.Now().Format("2006-01-02 15:04:05.000000000"))
	}
}

func verification() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.GetHeader("userId")
		token := context.GetHeader("token")
		session := context.GetHeader("session")

		// redis 中应当缓存用户的信息，[userId:业主Id:token] 为field，各种用户信息为value，其中包含用户的session
		// 用户在登录时会从mysql中拿到用户数据，哈希缓存到redis中，因此session不对的用户验证不通过

		log.Printf("userId  : %s", userId)
		log.Printf("token   : %s", token)
		log.Printf("session : %s", session)

		context.Next()
	}
}

func websocketVerification() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Query("userId")
		token := context.Query("token")
		session := context.Query("session")

		log.Printf("userId  : %s", userId)
		log.Printf("token   : %s", token)
		log.Printf("session : %s", session)

		if userId != "123456" {
			context.String(400,"userId wrong")
			return
		}

		context.Next()
	}
}
