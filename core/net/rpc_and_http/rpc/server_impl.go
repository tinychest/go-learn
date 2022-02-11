package main

import (
	"go-learn/core/net/rpc_and_http/dto"
)

type Hello struct{}

func (s *Hello) Hello(args *dto.Args, reply *dto.Reply) error {
	reply.Code = 200
	reply.Message = "hello"
	return nil
}
