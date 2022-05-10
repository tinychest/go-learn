package encode

import (
	"encoding/base64"
	"testing"
)

// [简介]
// Base64 编码是网络传输中最常见的用于传输 8Bit 字节码的编码方式之一。
// Base64 编码是从二进制到字符的过程，可用于在 HTTP 环境下传递较长的标识信息。
// 采用 Base64 编码具有不可读性，需要解码后才能阅读。
//
// [Go 中 base64 包中提供的加解码实例]
// - [StdEncoding]
// 字符集：ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/（填充 =）
// - [RawStdEncoding]
// 字符集同上，不填充
// - [URLEncoding]
// 字符集：ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_（填充 =）
// - [RawURLEncoding]
// 字符集同上，不填充
//
// [相关资料]
// 详见 Base64.md
//
// [前端 Base64 数据的内嵌样例]
// html：<img width="40" height="30" src="data:image/jpg;base64,/9j/4QMZRXhpZgAASUkqAAgAAAAL...." /">
// css：.demoImg{ background-image: url("data:image/jpg;base64,/9j/4QMZRXhpZgAASUkqAAgAAAAL...."); }

func TestBase64DiffRaw(t *testing.T) {
	src := "+"

	// 带不带 Raw 的区别
	dst1 := encoding(base64.URLEncoding, src) // 编码为：Kw==
	decoding(base64.URLEncoding, dst1)        // 能够正常解码出：+

	dst2 := encoding(base64.RawURLEncoding, src) // 编码为：Kw
	decoding(base64.RawURLEncoding, dst2)        // 能够正常解码出：+
}

// 为了得出怎样的字节序列经过 StdEncoding 能够得到 + 的结果（回忆 base64 实现）
func TestHowEncodeToGetPlus(t *testing.T) {
	// [失败案例]
	// 0x3E，没能得到 "+"
	// [案例1]
	// 0 0 0x3F --3 字节按每 6 位 1 取得到 4 字节-→ 0 0 0 3F(63) --经过码表映射-→ "A A A /"
	// [案例2]
	// 很简单就能得出 0, 0, 0x3E 可以得到 "A A A +"
	// [案例3]
	// + + + +（目标）
	// 62 62 62 62（码表下标）
	// 0x3E、0x3E、0x3E、0x3E（16 进制码表下标）
	// [11 1110][11 1110][11 1110][11 1110]（4 * 6）
	// [1111 1011] [1110 1111] [1011 1110]（3 * 8）
	// 251 239 190（10 进制） 或 0xFB 0xEF 0xBE
	// [回过头来]
	// 0x3E → 0011 1110 → 0011 1110 0000（填充 4 位） → Pg==（缺少两个数，填充 =）
	// 所以应该是 1111 1100 → 252 | 0xFB
	// 这个从 ASCII 码表可查不出东西，所以控制台打印了 �

	// src := []byte{0x3E}
	// src := []byte{0, 0, 0x3F}
	// src := []byte{0, 0, 0x3E}
	// src := []byte{251, 239, 190}
	// src := []byte{0xFB, 0xEF, 0xBE}
	src := []byte{0b_1111_1000}

	encoding(base64.StdEncoding, string(src))
}

// 源消息经过特殊定义，为了凸显字符集的不同，带来实际开发中的问题
// 在 Go 源码注释中，URLEncoding 用于 URL 地址栏 和 文件命名中。因为 StdEncoding 码表中的 + 和 / 在 URL 地址规范中属于非法字符
func TestBase64Differ(t *testing.T) {
	src := string([]byte{251, 239, 190})

	// url 编码 → 编码为：----
	// 假设后端返回一段 std 编码的数据
	dst := encoding(base64.URLEncoding, src)

	// std 解码 → 解码失败
	// 前端是使用 window.atob 方法对 base64 加密的数据进行解密的，方法实际使用的码表就是 StdEncoding，是无法解密成功的
	decoding(base64.StdEncoding, dst)

	// [案例]
	// 实际开发中遇到的，前端：
	// → 对静态常量的参数进行 Base64 编码
	// → 放到地址栏上（前端带上参数进行路由跳转，虽然不懂前端工程化，但感觉是很好的做法）
	// → 从地址栏上获取参数进行 Base64 解码
	// → 这个参数作为 JSON 的一部分，JSON 整个使用 crypto.js 加密，并传给后端
	// → 后端解密失败，报出的错误点是对 Base64 编码的数据进行解密，发现了 %
	// → 为什么会有 %，经过前端代码调试，发现 PC 端 和 Mobile 端的 PC 端在地址栏获取 Base64 编码的参数忘记解密了
	// → 那 % 怎么来的呢，实际就是因为前端使用的 Base64 编码方式的码表是包含 + 和 / 的 URL 非法字符的，URL 会分别将 + 和 / 转译成 %2B 和 %2F（可以通过 encodeURIComponent 方法进行测试）
	//
	// [解决（最初尝试，没有成功）]
	// 最初认为是特殊字符导致的，就从前端考虑，调研了解到 Unicode Safe 的 Base64 编解码方式
	// 但是因为实际解密者是后端，unescape、encodeURIComponent 等相关方法在后端都没有现成的实现，所以这样的做法就不太好了
	// 便改为单纯将 Base64 编码后数据中的 + 和 \ 删除掉，虽然，一时好像没问题，但这肯定是有问题的做法
	// （当时，错误情况很难复现，就像上面为了给出能够编码出特殊字符的样例，都只能采用字节数据了）
	// 参见：https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent
	//
	// [优化]
	// 加密的实现也复杂，所以改成了用 Base64 编码就行
	//
	// [最终]
	// 整个，其实是前端逻辑太复杂了，能力差一些的同事都跟不上，应当考虑简化流程
	// → 既然是后端给的常量，后端直接给好 Base64 加密好的常量，来简化前端操作和整个流程（也就是人工的手动操作来简化流程）
	// 后端的 Base64 编码方式是 std 的，因为需要和前端保持一致嘛；正好发现所有常量参数的 Base64 编码的结果中都没有包含特俗字符，所以就不了了之了
	//
	// [其他]
	// - 其实这边是真的复杂，还有一块是没有提及的，就是上面说的 crypto.js 的 AES 加密，后端是无现成的实现，但是项目领导人（给出静态常量参数的开发者）
	// 对应实现了；
	// - 另外一个细节就是后端报错实际是 AES 中 Base64 解密报错了（AES 中包含了对 Base64 算法的引用）
	//
	// - 项目后续，解密步骤都取消了，因为完全没有必要
	//
	// [Unicode Safe 的 Base64 编解码方式]
	// 参见：https://developer.mozilla.org/zh-CN/docs/Web/API/btoa#unicode_%E5%AD%97%E7%AC%A6%E4%B8%B2
	// 参见：https://html.spec.whatwg.org/multipage/webappapis.html#dom-windowbase64-btoa（有说明字符码值范围）
	// （网站负责人当然都明白，并且肯定有其他大佬遇到类似的问题，所以里面给出的解决方法，包含三方的实现）
	// [浏览器地址栏的 百分比编码]
	// URL Standard：https://url.spec.whatwg.org/#percent-encoded-bytes
}
