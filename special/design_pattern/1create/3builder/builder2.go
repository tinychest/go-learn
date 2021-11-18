package _builder

import "fmt"

/* 构建者 最佳实践 */

// ResourcePoolConfigOption 核心配置
type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

// 作为参数的函数，调用者可自定义配置行为
type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	c := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.maxTotal < 0 || c.maxIdle < 0 || c.minIdle < 0 {
		return nil, fmt.Errorf("args err, c: %v", c)
	}

	if c.maxTotal < c.maxIdle || c.minIdle > c.maxIdle {
		return nil, fmt.Errorf("args err, c: %v", c)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: c.maxTotal,
		maxIdle:  c.maxIdle,
		minIdle:  c.minIdle,
	}, nil
}
