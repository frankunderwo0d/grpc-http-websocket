package charge

import (
	chargeModel "grpc/model/charge"
	"log"
	"time"
)

func Charge(request *chargeModel.Request) (*chargeModel.Response, error) {
	log.Println("Service method [charge] in use")
	log.Printf("request : %+v", request)

	time.Sleep(time.Second * time.Duration(2))

	return &chargeModel.Response{
		Code:   200,
		Msg:    "success",
		PayUrl: "a9s87f8s97a69g9sa79s8d",
	}, nil
}
