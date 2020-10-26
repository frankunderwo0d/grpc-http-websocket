package test

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	active "grpc/active"
	chargeModel "grpc/model/charge"
	"io"
	"log"
	"sync"
	"testing"
)

func TestAServiceGetFeature(T *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := active.NewAServiceClient(conn)

	response, err := client.GetFeature(context.Background(), &active.Request{
		Path:    "/user",
		Method:  0,
		Headers: map[string]string{"token": "frank-interview"},
		Params:  map[string]string{"result": "let me know the answer after the interview"},
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response.GetStatusCode())
	log.Println(response.GetStatusMsg())
	log.Println(string(response.GetData()))
}

func TestAServiceListFeature(T *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := active.NewAServiceClient(conn)

	lfc, err := client.ListFeature(context.Background(), &active.Request{
		Path:    "/user",
		Method:  0,
		Headers: map[string]string{"token": "frank-interview"},
		Params:  map[string]string{"result": "let me know the answer after the interview"},
	})
	if err != nil {
		log.Println(err)
		return
	}

	for {
		response, err := lfc.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeature(_)=_, %v", client, err)
		}
		log.Println(response.GetStatusCode())
		log.Println(response.GetStatusMsg())
		log.Println(string(response.GetData()))
	}
}

func TestAServiceRecordFeature(T *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := active.NewAServiceClient(conn)

	rfc, err := client.RecordFeature(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 3; i++ {
		request := &active.Request{
			Path:    "/user",
			Method:  0,
			Headers: map[string]string{"token": "frank-interview"},
			Params:  map[string]string{"result": fmt.Sprintf("let me know the answer after the interview - %d", i)},
		}
		if err := rfc.Send(request); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.send(%v) = %v", client, request, err)
		}
	}

	response, err := rfc.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response.GetStatusCode())
	log.Println(response.GetStatusMsg())
	log.Println(string(response.GetData()))
}

func TestAServiceFeatureChat(T *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := active.NewAServiceClient(conn)

	fcc, err := client.FeatureChat(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			response, err := fcc.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Fail to recive response %v", err)
			}
			log.Println(response.GetStatusCode())
			log.Println(response.GetStatusMsg())
			log.Println(string(response.GetData()))
		}
	}()

	for i := 0; i < 3; i++ {
		request := &active.Request{
			Path:    "/user",
			Method:  0,
			Headers: map[string]string{"token": "frank-interview"},
			Params:  map[string]string{"result": fmt.Sprintf("let me know the answer after the interview - %d", i)},
		}
		if err := fcc.Send(request); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.send(%v) = %v", client, request, err)
		}
	}
	if err := fcc.CloseSend(); err != nil {
		log.Println(err)
	}
	wg.Wait()
	log.Println("Done")
}

func TestCommunicationService(T *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := active.NewCommunicationClient(conn)

	req := &chargeModel.Request{
		UserID:       123456789,
		Token:        "ASD8F97AS87FD8A7S",
		Session:      "sa89f6sd5fa6s56",
		Amount:       50000,
		PayType:      1,
		PayChannelID: 52,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := client.Send(context.Background(), &active.CRequest{
		ProtocolId: 101001,
		Data:       reqData,
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response.GetStatusCode())
	log.Println(response.GetProtocolId())
	log.Println(string(response.GetData()))
}
