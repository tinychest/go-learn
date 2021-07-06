package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"testing"
)

var reqFunc = func(wg *sync.WaitGroup, i int) {
	params := url.Values{
		"phone": []string{"18942350318"},
	}
	resp, err := http.PostForm("http://127.0.0.1:8651/v2010/sms/send", params)
	if err != nil {
		fmt.Printf("第 %d 次请求发送验证码失败 %s\n", i, err)
		goto end
	}
	if bs, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("第 %d 次请求发送验证码，读取响应失败 %s\n", i, err)
		goto end
	} else {
		fmt.Printf("第 %d 次请求发送验证码，响应结果 %s\n", i, string(bs))
		goto end
	}

end:
	wg.Done()
}

// 短信验证码接口对电话号码有做每 30 秒才能请求一次短信验证码，这里在短时间发送多个请求，验证验证是否有效
func TestValidLimit(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reqFunc(wg, i+1)
	}

	wg.Wait()
}
