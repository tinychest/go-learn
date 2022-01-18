package main

import (
	"go-learn/core/net/rpc/dto"
)

type IHello interface {
	Hello(*dto.Args, *dto.Reply) error
}

type Hello struct{}

func (s *Hello) Hello(args *dto.Args, reply *dto.Reply) error {
	reply.Code = 200
	reply.Message = "hello"
	return nil
}
