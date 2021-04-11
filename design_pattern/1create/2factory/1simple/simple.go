package _simple

/* 简单工厂：适用于类型改动不频繁的情况 */

// 接口
type IParser interface {
	Parse(data []byte)
}

// 实现1
type jsonParser struct{}

func (j jsonParser) Parse(data []byte) {
	panic("implement me")
}

// 实现2
type yamlParser struct{}

func (Y yamlParser) Parse(data []byte) {
	panic("implement me")
}

// 用于获取 实例
func NewParser(t string) IParser {
	switch t {
	case "json":
		return jsonParser{}
	case "yaml":
		return yamlParser{}
	}
	return nil
}
