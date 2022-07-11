package mock

// 1.创建 db.go → 编写要测试方法
// 2."mock" Open In Terminal
// 3. 生成包含打桩的相关辅助方法的文件 db_mock.go：mockgen -source=db.go -destination=db_mock.go -package=mock
// 4.创建 db_test.go → 编写测试用例
// 5.执行测试：go test . -cover -v
