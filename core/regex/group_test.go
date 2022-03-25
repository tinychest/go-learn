package regex

import (
	"encoding/json"
	"regexp"
	"testing"
)

// 分组定义语法：(?P<组名>组内容)
// 表达式中分组引用（按下标引用不行，按名字引用不知道怎么引用）
//
// [api 方法补充说明]
// FindStringSubmatch 只捕获一次匹配正则表达式的结果
// FindAllStringSubmatch 捕获所有匹配正则表达式的结果
// SubexpNames 返回定义的所有捕获组的名称

// 捕获组
func TestCaptureGroup(t *testing.T) {
	// 中文描述价格的表达式匹配
	s1 := `包子一共 10 元`
	//language=GoRegExp
	r1 := "^[\u4E00-\u9FA5]+一共 \\d+ 元$"
	t.Log(regexp.MatchString(r1, s1))

	// 假如描述价格的物件数量不确定，就需要借助捕获组了
	s2 := `包子一共 10 元，油条一共 20 元`
	//language=GoRegExp
	r2 := "^[\u4E00-\u9FA5]+一共 \\d+ 元(，[\u4E00-\u9FA5]+一共 \\d+ 元)?$"
	t.Log(regexp.MatchString(r2, s2))
}

// 非捕获组
// 语法：(?:)
func TestUncaptureGroup(t *testing.T) {
	// 假如我们希望获取所有物件的价格，那我们应该为价格部分的表达式添加捕获组，然后，获取捕获组的匹配结果
	s := `包子一共 10 元，油条一共 20 元`
	//language=GoRegExp
	r1 := "^[\u4E00-\u9FA5]+一共 (\\d+) 元(，[\u4E00-\u9FA5]+一共 (\\d+) 元)?$"
	res := regexp.MustCompile(r1).FindStringSubmatch(s)
	bs, _ := json.Marshal(res)
	t.Log(string(bs))

	// 上面，你会发现出现了意料之外的捕获组内容
	// - 数组的第一个元素就是目标串本身，这是默认会加上的捕获组（跳过第一个元素）
	// - 后边的大组都被捕获了（那就需要用到非捕获组了，表示不希望捕获该组的内容）
	r4 := "^[\u4E00-\u9FA5]+一共 (\\d+) 元(?:，[\u4E00-\u9FA5]+一共 (\\d+) 元)?$"
	res = regexp.MustCompile(r4).FindStringSubmatch(s)
	bs, _ = json.Marshal(res[1:])
	t.Log(string(bs))
}

// 你会发现通过下标去获取捕获组内容显得十分笨拙，这时我们可以通过为捕获组定义名称，然后，按照顺序自己拼装出参数对应的结果
func TestNamedCaptureGroup(t *testing.T) {
	date := "1997-10-05"
	//language=GoRegExp
	r := `(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`

	reg := regexp.MustCompile(r)
	matchs := reg.FindStringSubmatch(date)
	names := reg.SubexpNames()

	m := make(map[string]string, len(names))
	for i := 1; i < len(matchs); i++ {
		m[names[i]] = matchs[i]
	}
	t.Log(m)
}
