package controller

import (
	"context"
	"fmt"
	active "grpc/active"
	"io"
	"log"
	"time"
)

// 简单的重写AService服务
type AService struct {
}

// 1.简单的rpc请求
func (as *AService) GetFeature(ctx context.Context, request *active.Request) (*active.Response, error) {
	log.Println("AService's method [GetFeature] be in use")
	log.Printf("Request.Path : %s", request.GetPath())
	log.Printf("Request.Method : %s", active.Method_name[int32(request.GetMethod())])
	log.Printf("Request.Headers : %v", request.GetHeaders())
	log.Printf("Request.Params : %v", request.GetParams())

	time.Sleep(time.Duration(2) * time.Second)

	response := &active.Response{
		StatusCode: 200,
		StatusMsg:  "success",
		Data:       []byte("Congratulations , you are pass the interview"),
	}

	return response, nil
}

// 2.服务器端流式rpc，接受一个请求，持续性发送响应
func (as *AService) ListFeature(request *active.Request, lfs active.AService_ListFeatureServer) error {
	log.Println("AService's method [ListFeature] be in use")
	log.Printf("Request.Path : %s", request.GetPath())
	log.Printf("Request.Method : %s", active.Method_name[int32(request.GetMethod())])
	log.Printf("Request.Headers : %v", request.GetHeaders())
	log.Printf("Request.Params : %v", request.GetParams())

	for i := 0; i < 3; i++ {
		response := &active.Response{
			StatusCode: 200,
			StatusMsg:  "success",
			Data:       []byte(fmt.Sprintf("Congratulations , you are pass the interview - %d",i)),
		}

		if err := lfs.Send(response); err != nil {
			log.Println(err)
			break
		}
	}

	return nil
}

// 3.客户端流式rpc，持续性接受请求，在请求结束时发送一个响应回去
func (as *AService) RecordFeature(rfs active.AService_RecordFeatureServer) error {
	log.Println("AService's method [RecordFeature] be in use")

	for {
		request , err := rfs.Recv()
		if err == io.EOF {
			return rfs.SendAndClose(&active.Response{
				StatusCode: 200,
				StatusMsg:  "success",
				Data:       []byte("Congratulations , you are pass the interview"),
			})
		}
		if err != nil {
			return err
		}

		log.Printf("Request.Path : %s", request.GetPath())
		log.Printf("Request.Method : %s", active.Method_name[int32(request.GetMethod())])
		log.Printf("Request.Headers : %v", request.GetHeaders())
		log.Printf("Request.Params : %v", request.GetParams())
	}
}

// 4.双向流式RPC，持续性接收请求及发送响应，类比全双工socket通讯
func (as *AService) FeatureChat(fcs active.AService_FeatureChatServer) error {
	log.Println("AService's method [FeatureChat] be in use")

	for {
		request , err := fcs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Request.Path : %s", request.GetPath())
		log.Printf("Request.Method : %s", active.Method_name[int32(request.GetMethod())])
		log.Printf("Request.Headers : %v", request.GetHeaders())
		log.Printf("Request.Params : %v", request.GetParams())

		err = fcs.Send(&active.Response{
			StatusCode: 200,
			StatusMsg:  "success",
			Data:       []byte("Congratulations , you are pass the interview"),
		})
		if err != nil {
			return err
		}
	}
}
