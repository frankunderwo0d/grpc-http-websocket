package charge

import (
	"encoding/json"
	"errors"
	active "grpc/active"
	"grpc/model/charge"
	chargeService "grpc/service/charge"
	"log"
)

func Charge(request *active.CRequest) (*active.CResponse, error) {
	log.Println("101001 be in use")
	if len(request.GetData()) == 0 {
		return nil, errors.New("request.data can't be nil")
	}

	req := &charge.Request{}
	err := json.Unmarshal(request.GetData(), req)
	if err != nil {
		return nil, err
	}

	rsp, err := chargeService.Charge(req)
	if err != nil {
		return nil, err
	}

	rspData, err := json.Marshal(rsp)
	if err != nil {
		return nil, err
	}

	response := &active.CResponse{
		StatusCode: 200,
		ProtocolId: request.GetProtocolId() + 1,
		Data:       rspData,
	}

	return response, nil
}
