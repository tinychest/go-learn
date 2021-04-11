package regex

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// 校验字符串的格式是否是逗号隔开自然数的格式，且自然数不能重复
func TestIdsValidate(t *testing.T) {
	type FunctionParam struct {
		ids string
	}
	testcases := []struct {
		FunctionName  string
		FunctionParam FunctionParam
		ExpectError   bool
	}{
		{FunctionName: "IdQueryParamValidate", FunctionParam: FunctionParam{""}, ExpectError: true},
		{FunctionName: "IdQueryParamValidate", FunctionParam: FunctionParam{"1 "}, ExpectError: true},
		{FunctionName: "IdQueryParamValidate", FunctionParam: FunctionParam{"1,2,++"}, ExpectError: true},
		{FunctionName: "IdQueryParamValidate", FunctionParam: FunctionParam{"1,2,2"}, ExpectError: true},
	}

	for _, testcase := range testcases {
		t.Run("参数校验方法", func(t *testing.T) {
			err := IdsValidate(testcase.FunctionParam.ids)
			if (err != nil) != testcase.ExpectError {
				t.Errorf("%s() error = %v, expect %v", testcase.FunctionName, err, testcase.ExpectError)
				return
			}
		})
	}
}

func IdsValidate(ids string) error {
	// 这里一定要加上 ^ $ 的前后限定，不然正则表达式的含义就变成了只要包含数据就可以
	if rulePattern := regexp.MustCompile(`^\d+(,\d+)*$`); !rulePattern.MatchString(ids) {
		fmt.Printf("参数格式非法：【ids：%s】\n", ids)
		return errors.New("参数格式非法")
	}

	idStrSlice := strings.Split(ids, ",")

	idDuplicateMap := make(map[string]int, len(idStrSlice))
	for _, id := range idStrSlice {
		idDuplicateMap[id] = idDuplicateMap[id] + 1
	}
	duplicateIds := make([]string, 0, len(idDuplicateMap))
	for idStr, times := range idDuplicateMap {
		if times > 1 {
			duplicateIds = append(duplicateIds, idStr)
		}
	}
	if len(duplicateIds) > 0 {
		fmt.Printf("参数非法：【%s】 - 重复的id：%v\n", ids, duplicateIds)
		return errors.New(fmt.Sprintf("参数非法，重复的元素：%v", duplicateIds))
	}

	return nil
}
