package _method

/* 工厂方法：适用于对象创建逻辑比较复杂的情况 */

// 目标接口
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

// 工厂方法接口
type IParserFactory interface {
	CreateParser() IParser
}

// 创建 目标接口实现1 的工厂
type yamlParserFactory struct{}

func (y yamlParserFactory) CreateParser() IParser {
	return yamlParser{}
}

// 创建 目标接口实现2 的工厂
type jsonParserFactory struct{}

func (j jsonParserFactory) CreateParser() IParser {
	return jsonParser{}
}

// 用于获取 创建实例的工厂
func NewIParserFactory(t string) IParserFactory {
	switch t {
	case "json":
		return jsonParserFactory{}
	case "yaml":
		return yamlParserFactory{}
	}
	return nil
}
