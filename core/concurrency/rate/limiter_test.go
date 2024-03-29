package rate

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 下面两个文章说的很好，博主格局大，说事情有条理，结合实例来说，说的清清楚楚，还有拓展
// https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247488111&idx=1&sn=4e4eb013481d030c0ccfdfd3f143c0bf&chksm=fa80c3f8cdf74aee9314df27c27e4ee4eb38cc7971222f756b546b0de2f74ac76410cc0bd4ca&scene=178&cur_album_id=1326949382503219201#rd
// https://mp.weixin.qq.com/s/qAKQm9CVNxk-ltUxHql1nw

// NewLimiter
// 参数1：向桶中投放令牌的频率，每多少时间放一个，或者每秒放多少个（无论如果传参，执行速率始终是均匀的）
// 参数2：桶的容量

// 并没有单独维护一个 Timer 和队列去真的每隔一段时间向桶中放令牌，而是仅仅通过计数的方式表示桶中剩余的令牌。每次消费取 Token 之前会先根据上次更新令牌数的时间差更新桶中 Token 数
//   通过 time.Timer 去实现等待
// 通过给 ctx 设置 timeout 来指定最大等待时间

// 【Wait、WaitN】是阻塞形式应对并发（最好设置上等待时长）
// 【Allow、AllowN】希望以立即响应形式去应对并发，应该使用 Allow 相关方法
// 【Reserve、ReserveN】如果希望更加定制化的根据需要等待的时间做处理，则可以使用 Reserve 方法

// SetLimit(Limit) 改变放入 Token 的速率
// SetBurst(int) 改变 Token 桶大小

// 拓展：uber-go/ratelimit

var i int

func TestRatePack(t *testing.T) {
	// 每秒 1 次
	// v := rate.Every(time.Second)
	// 1 秒 5 次
	v := rate.Limit(5)

	limiter := rate.NewLimiter(v, 10)

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		err := waitDo(limiter, ctx)
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		err := waitDo(limiter, ctx)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func waitDo(limiter *rate.Limiter, ctx context.Context) error {
	start := time.Now()
	i++

	err := limiter.Wait(ctx)
	if err != nil {
		return err
	}

	// 你的业务逻辑
	fmt.Printf("【index：%3d】 【interval：%d ms】\n", i, time.Since(start).Milliseconds())
	return nil
}
