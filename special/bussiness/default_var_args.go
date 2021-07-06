package bussiness

// 当参数 params 有，就取第一个，不然就取一个默认值
func varArgsHandleLogic(def string, params ...string) string {
	return append(params, def)[0]
}