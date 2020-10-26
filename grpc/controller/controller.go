package controller

import (
	"context"
	"errors"
	active "grpc/active"
	"grpc/controller/charge"
	"log"
)

type (
	HandleFunction func(*active.CRequest) (*active.CResponse, error)
)

var (
	Controller = &Communication {
		make(map[uint32]HandleFunction),
	}
)

func init() {
	err := Controller.RegisterHandler(101001, charge.Charge)
	if err != nil {
		log.Println(err)
		return
	}
}

// 定制商用版
type Communication struct {
	handleMap map[uint32]HandleFunction
}

func (c *Communication) Send(ctx context.Context, request *active.CRequest) (*active.CResponse, error) {
	handler , exist := c.handleMap[request.GetProtocolId()]
	if !exist {
		return nil,errors.New("the handle function of protocol id not exist")
	}

	return handler(request)
}

// 注册单个处理器
func (c *Communication) RegisterHandler(protocolID uint32, handler HandleFunction) error {
	if handler == nil {
		return errors.New("handler can't be nil")
	}
	if _, exist := c.handleMap[protocolID]; exist {
		return errors.New("protocol id has been registered")
	}
	c.handleMap[protocolID] = handler
	return nil
}

// 注册多个处理器
func (c *Communication) RegisterHandlers(handlers map[uint32]HandleFunction) error {
	if handlers == nil {
		return errors.New("handlers can't be nil")
	}
	for protocolID, handleFunc := range handlers {
		if err := c.RegisterHandler(protocolID, handleFunc); err != nil {
			return err
		}
	}
	return nil
}