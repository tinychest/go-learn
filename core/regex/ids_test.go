package regex

import (
	"errors"
	"regexp"
	"strings"
	"testing"
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

func TestIdsValidate(t *testing.T) {
	type args struct {
		idsStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test01",
			args:    args{idsStr: ""},
			wantErr: true,
		},
		{
			name:    "test02",
			args:    args{idsStr: "1 "},
			wantErr: true,
		},
		{
			name:    "test03",
			args:    args{idsStr: "1,2,++"},
			wantErr: true,
		},
		{
			name:    "test04",
			args:    args{idsStr: "1,2,2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IdsValidate(tt.args.idsStr); (err != nil) != tt.wantErr {
				t.Errorf("IdsValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
