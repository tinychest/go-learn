package regex

import (
	"errors"
	"regexp"
	"strings"
)

func IdsValidate(idsStr string) error {
	// 起始、结束符的作用：这里一定要加上 ^ $ 的前后限定，不然正则表达式的含义就变成了只要包含数据就可以
	if rulePattern := regexp.MustCompile(`^\d+(,\d+)*$`); !rulePattern.MatchString(idsStr) {
		return errors.New("参数格式非法")
	}

	ids := strings.Split(idsStr, ",")
	idM := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		if _, ok := idM[id]; ok {
			return errors.New("参数非法，存在重复元素")
		}
		idM[id] = struct{}{}
	}
	return nil
}
