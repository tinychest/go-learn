package _template

import (
	"errors"
	"fmt"
)

// 背景描述
// 定义一套发送短信的流程：校验基本信息（电话号码是肯定要校验的）、对接指定服务产商请求发送短信
// 实际可能需要对接好几个三方的短信服务产商，所以希望能够有一套好的逻辑流程，并复用其中的一些逻辑

// ISMS 短信发送接口定义
type ISMS interface {
	// Valid 所有短信发送都需要调用的基础信息校验（可以按照需求定义更多的复用逻辑）
	Valid(phone string) error
	// Send 核心的短信发送方法（模板方法骨架）
	Send(msg, phone string) error
}

// BaseSms 短信发送的模板方法
type BaseSms struct {
	ISMS
}

func (s *BaseSms) Valid(phone string) error {
	// 校验手机号码长度
	if len(phone) != 11 {
		return errors.New("phone illegal")
	}
	return nil
}

func (s *BaseSms) Send(msg, phone string) error {
	// 模板方法类是抽象的，不具备单独调用能力
	// （这里直接通过模板方法基础类调用会引发 nil pointer panic）
	if err := s.ISMS.Valid(msg); err != nil {
		return err
	}
	return s.ISMS.Send(msg, phone)
}

// XxxSMS 实现一
type XxxSMS struct {
	base BaseSms
}

func NewXxxSMS() ISMS {
	// 这里可能比较绕，但是是实现模板方法的核心
	xxx := &XxxSMS{}
	xxx.base = BaseSms{xxx}
	return xxx
}

func (s *XxxSMS) Valid(phone string) error {
	// 直接复用模板方法逻辑
	return s.base.Valid(phone)
}

func (s *XxxSMS) Send(msg, phone string) error {
	// 应该调用当前具体的检验方法
	if err := s.Valid(phone); err != nil {
		return err
	}
	// 注意，如果这里还调用基类的方法，将引起死递归
	fmt.Printf("Xxx Send %s to %s\n", msg, phone)
	return nil
}

// YyySMS 实现二
type YyySMS struct {
	base BaseSms
}

func NewYyySMS() ISMS {
	xxx := &YyySMS{}
	xxx.base = BaseSms{xxx}
	return xxx
}

func (s *YyySMS) Valid(phone string) error {
	// 自定义校验逻辑
	if len(phone) != 3 {
		return errors.New("phone illegal")
	}
	return nil
}

func (s *YyySMS) Send(msg, phone string) error {
	if err := s.Valid(phone); err != nil {
		return err
	}
	fmt.Printf("Yyy Send %s to %s\n", msg, phone)
	return nil
}
