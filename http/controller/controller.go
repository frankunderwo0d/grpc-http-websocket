package controller

import (
	active "base_grpc/active"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	StatusCodeNil = -1
	ProtocolIDNil = 0
)

var (
	clientConn active.CommunicationClient

	upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

func init() {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	//defer conn.Close()

	client := active.NewCommunicationClient(conn)
	clientConn = client
}

func SendMessage(protocolID uint32, req interface{}) (int64, uint32, []byte, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return StatusCodeNil, ProtocolIDNil, []byte{}, err
	}

	response, err := clientConn.Send(context.Background(), &active.CRequest{
		ProtocolId: protocolID,
		Data:       reqData,
	})
	if err != nil {
		return StatusCodeNil, ProtocolIDNil, []byte{}, err
	}

	return response.GetStatusCode(), response.GetProtocolId(), response.GetData(), nil
}

func checkOrigin(request *http.Request) bool {
	return true
}