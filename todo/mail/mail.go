package mail

import (
	_ "embed"
	"encoding/base64"
	"go-learn/tool"
	"gopkg.in/gomail.v2"
	"strings"
	"text/template"
)

// From 一般会带有邮箱地址的中文说明，java 中可以直接处理，但是这里出现了问题，大致原理是这样的
// SetHeader 的底层中，有对中文进行 UTF-8 的编码处理，但是处理的有问题
// 所以只能像下面这样手动，对中文进行 UTF8 的编码处理

//go:embed mail.tmpl
var mailTmpl string

var d = gomail.NewDialer("smtphz.qiye.163.com", 25, "flow03@gaojidata.com", "JEwCkl66Ta")

func SendTo(to, username, password string) {
	m := gomail.NewMessage()

	// email：邮件标题
	m.SetHeader("Subject", "高绩用户创建通知")

	// email：发送方
	prefix := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte("高绩专业建设平台工作组")) + "?="
	m.SetHeader("From", prefix+" <flow03@gaojidata.com>")

	// email：发送给
	m.SetHeader("To", to)

	// email：邮件内容
	b := &strings.Builder{}
	b.Grow(300) // username：小明 password：12346 得到的 builder.Len() 为 290

	tpl := template.Must(template.New("mail").Parse(mailTmpl))
	params := Content{
		Username: username,
		Password: password,
		Datetime: tool.NowDateCN(),
	}
	if err := tpl.Execute(b, params); err != nil {
		panic(err)
	}
	m.SetBody("text/html", b.String())

	// 发送邮件
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
