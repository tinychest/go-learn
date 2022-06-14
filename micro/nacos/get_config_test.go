package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"testing"
)

// 使用 Nacos SDK 向本地多个项目配置信息的默认配置启动的 Nacos，获取指定项目的配置文件信息

const (
	ProjectName = "gin-sqlx-base"
	EnvName     = "native_wmc"
)

// 注意下面的文件路径，遇到了生成配置文件缓存未设置 gitignore 导致 git 提交历史泄漏了重要的账号信息
// 假如泄漏，应该按照如下步骤处理
// > 仓库设置私有
// > 修改账号密码
// > 删除历史中的文件（可能较为繁琐）
func TestGetConfig(t *testing.T) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "d924005c-3bba-4187-a29d-a2121a1f5efa",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./target/log",
		CacheDir:            "./target/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos", // 可以通过 Nacos 的启动日志了解 context 是什么
			Port:        8848,
			Scheme:      "http",
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	res, err := configClient.GetConfig(vo.ConfigParam{
		DataId: ProjectName,
		Group:  EnvName,
		// Content:  "",
		// DatumId:  "",
		// Type:     "",
		// OnChange: nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("获取到的配置信息：")
	t.Log(res)
}
