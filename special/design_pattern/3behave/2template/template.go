package _template

import (
	"errors"
	"fmt"
)

// ISMS 短信发送
type ISMS interface {
	Valid(msg string) error
	Send(msg string, phone int) error
}

// BaseSms 短信发送的模板方法
type BaseSms struct {
	ISMS
}

func NewSms(sms ISMS) ISMS {
	return &BaseSms{ISMS: sms}
}

func (s *BaseSms) Valid(msg string) error {
	if len(msg) > 63 {
		return errors.New("msg too long")
	}
	return nil
}

func (s *BaseSms) Send(msg string, phone int) error {
	if err := s.Valid(msg); err != nil {
		return err
	}
	return s.ISMS.Send(msg, phone)
}

/*
以上将作为模板模式的模板方法

实际可以不同的短信发送方法（比如，可以选用各大云厂商的服务）
*/

type XxxSMS struct {
	baseSms BaseSms
}

func NewXxxSms() ISMS {
	res := &XxxSMS{}
	res.baseSms = BaseSms{ISMS: res}
	return res
}

func (s *XxxSMS) Valid(msg string) error {
	// 直接复用
	return s.baseSms.Valid(msg)
}

func (s *XxxSMS) Send(msg string, phone int) error {
	// 自定义行为
	if err := s.Valid(msg); err != nil {
		return err
	}
	// 如果还调用父类的方法，将引起死递归
	// s.baseSms.Send(msg, phone)
	fmt.Println("Xxx 发送短信")
	return nil
}
