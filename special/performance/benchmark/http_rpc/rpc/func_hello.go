package main

import "go-learn/special/performance/benchmark/http_rpc/dto"

type IHello interface {
	Hello(p *dto.Param, r *dto.Reply) error
}

type Hello struct{}

func (s *Hello) Hello(param *dto.Param, reply *dto.Reply) error {
	reply.Code = 200
	reply.Message = "hello"
	return nil
}
