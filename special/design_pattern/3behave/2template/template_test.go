package _template

import "testing"

func TestTemplate(t *testing.T) {
	xxxSms := NewXxxSms()

	sms := NewSms(xxxSms)

	_ = sms.Send("123", 110)

	_ = xxxSms.Send("123", 110)
}
