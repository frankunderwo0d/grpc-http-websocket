package controller

import (
	grpcModel "base_grpc/model/charge"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type PayInfo struct {
	Amount     int64 `form:"amount"`
	PayType    int64 `form:"payType"`
	PayChannel int64 `form:"payChannel"`
}

func Lobby(ctx *gin.Context) {
	action := ctx.Param("action")
	log.Println(action)

	var payInfo PayInfo
	if err := ctx.ShouldBind(&payInfo); err != nil {
		ctx.JSON(401, gin.H{"error": err})
		return
	}
	log.Printf("%+v", payInfo)

	userId , err := strconv.Atoi(ctx.GetHeader("userId"))
	if err != nil {
		ctx.JSON(401, gin.H{"error": err})
		return
	}

	request := &grpcModel.Request{
		UserID:       int64(userId),
		Token:        ctx.GetHeader("token"),
		Session:      ctx.GetHeader("session"),
		Amount:       payInfo.Amount,
		PayType:      payInfo.PayType,
		PayChannelID: payInfo.PayChannel,
	}

	rd , _ := json.Marshal(request)
	log.Println(string(rd))

	statusCode, protocolID, data, err := SendMessage(101001, request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	log.Printf("Status code : %d", statusCode)
	log.Printf("Protocol ID : %d", protocolID)

	response := &grpcModel.Response{}
	err = json.Unmarshal(data, response)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	log.Printf("Response : %+v", *response)

	ctx.JSON(int(response.Code), response)
}
