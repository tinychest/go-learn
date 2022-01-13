package regex

import (
	"fmt"
	"regexp"
	"testing"
)

/*
方法名规律（尖括号代表重点）
Find...           返回<第一个>匹配的结果
Find...Index      返回匹配的<下标>
FindAll...        返回<所有>匹配的结果，参数 n 的意思是从目标串中捕获多少个（0 代表不捕获；负数 代表全部）
Find...String...  返回值类型的元素类型是 string，而不是 []byte

Find...Submatch   不仅返回匹配的结果，还包含匹配结果中 子捕获组 的内容
SubexpNames       单独拿出来说，获取正则表达式中所有定义捕获组名称
*/

func TestRegex(t *testing.T) {
	wordExpressTest(t)

	// quoteTest(t)
	// singleTest(t)
	//
	// r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	// t.Logf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	// t.Logf("%#v\n", r.SubexpNames())
	//
	// t.Log(matchTest("\\d{4}-\\d{2}-\\d{2}", "2020-01-01"))
	// t.Log(findAllStringSubmatchTest("`^(\\d+)([hd])-(\\d+)$`", "2d-4"))
	// t.Log(findStringSubMatchTest(`\d+-(\d+)-`, "1-22-"))
}

// 技术：regexp.QuoteMeta
// 业务：如果，参数将会被作为正则表达式中的常量部分，那么为了严谨，因为对特殊字符进行转义处理
func quoteTest(t *testing.T) {
	reg := `*`
	reg = regexp.QuoteMeta(reg)
	t.Log(reg)
}

// 技术：正则表达式不加 开始、结束符 带来的逻辑上的差异
func singleTest(t *testing.T) {
	var regex string
	// regex1 := `^\d$`         // false
	// regex2 := `\d+(,\d+)?`   // true
	// regex3 := `^\d+(,\d+)?$` // false
	compile, _ := regexp.Compile(regex)
	t.Log(compile.MatchString("-1"))
}

// 技术：\b 在正则表达式的作用 单词边界
func wordExpressTest(t *testing.T) {
	tableName := "xxx"

	r := fmt.Sprintf(`\b` + tableName + `(_\d+)?\b`)
	q := `SELECT * FROM xxx JOIN xxx_1 ON id_xxx_2`

	matches := regexp.MustCompile(r).FindAllString(q, -1)
	t.Log(matches)
}

// 技术：Match
// 业务：某个字符串是否符合正则表达式定义规则结构
func matchTest(regex, target string) bool {
	return regexp.MustCompile(regex).MatchString(target)
}

// 技术：FindStringSubmatch
// 业务：捕获 1-22- 中的 22
func findStringSubMatchTest(regex, target string) []string {
	compile, _ := regexp.Compile(regex)
	return compile.FindStringSubmatch(target)
}

// 技术：FindAllStringSubmatch
// 业务：测试 "((\d+)-)+" 等同于 "^((\d+)-)+$" 的全匹配能够获取第指定下标的 "(\d+)-" 么
//    全匹配 - FindStringSubmatch 不行，无法进行叠词的下标筛选，打印捕获组匹配的所有内容，也只能得到叠词捕获组的最后一次捕获内容
//    查找匹配 - FindAllStringSubmatch 可以
func reduplicationTest(text string) [][]string {
	// 1、非捕获组起作用，没有被 "?:" 修饰的 "(\d+)-" 没有再被捕获
	// 2、但是无法排除完整正则表达式对应的捕获组
	// compile, _ := regexp.Compile(`(?:(\d+)-)`)
	compile, _ := regexp.Compile(`(\d+)-`)
	return compile.FindAllStringSubmatch(text, -1)
}

/*
技术：FindAllStringSubmatch
业务：返回 target 匹配 regex 的所有捕获组内容

自己概括
[0][0]：正则表达式的在目标串第一个完整匹配（默认的完全正则作为的完整正则表达式）
[n1][n2]（n > 1）：第 n1 + 1 个完整匹配组的第 n2 + 1 个子捕获组的捕获内容

直接上示例
re := regexp.MustCompile(`a(x*)b`)
fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))

Output:
[["ab" ""]]
[["axxb" "xx"]]
[["ab" ""] ["axb" "x"]]
[["axxb" "xx"] ["ab" ""]]
*/
func findAllStringSubmatchTest(regex, target string) [][]string {
	// 数字1个或多个 h或d - 数字1个或多个（定义了 3 个匹配组）
	return regexp.MustCompile(regex).FindAllStringSubmatch(target, -1)
}
