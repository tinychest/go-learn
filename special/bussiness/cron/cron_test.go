package help

import (
	"testing"
	"time"
)

// cron 表达式语法：[秒] 分 时 日 月 周
// 说明：当前使用的三方类库只支持 五位，即不能传秒
// [详见 cron.New 方法的注释]

// @hourly： Every hour, starting an hour from now
// @every 1h30m：Every hour thirty, starting an hour thirty from now
// cron.New(cron.WithSeconds()) 这样创建定时任务，将无视后续传入的 cron 表达式

// 每分（的零秒）：【* * * * ?】
// 每个小时（的零分）：【0 * * * ?】
// 每天（的零点）：【0 0 * * ?】

// 每两个小时：【0 */2 * * ?】
// */1 = 1
func TestCron(t *testing.T) {
	var err = CronJob{
		{
			Name: "测试 cron",
			Cron: "* * * * * ?",
			Job: func() error {
				println("我是方法...")
				return nil
			},
		},
	}.Start()
	if err != nil {
		t.Log(err)
	}

	time.Sleep(time.Hour)
}
