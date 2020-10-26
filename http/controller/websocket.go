package controller

import (
	grpcModel "base_grpc/model/charge"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func WebSocketHandler(ctx *gin.Context) {
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}

	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var request grpcModel.Request
		err = json.Unmarshal(message,&request)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err})
			return
		}

		log.Printf("%s send : %+v", conn.RemoteAddr().String(), request)

		statusCode, protocolID, data, err := SendMessage(101001, request)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err})
			return
		}

		log.Printf("Status code : %d", statusCode)
		log.Printf("Protocol ID : %d", protocolID)

		if err := conn.WriteMessage(msgType, data); err != nil {
			log.Println(err)
			return
		}
	}
}
