package test

import (
	"flag"
	"fmt"
	"testing"
)

// 读取用户自定义参数
var configPath string

func init() {
	flag.StringVar(&configPath, "configs", "", "配置路径")
}

func Test_os(t *testing.T) {
	// testing.M.Run来执行测试会自动调用flag.Parse() 因此无需显示声明
	// flag.Parse()
	if configPath != "" {
		fmt.Printf("配置路径存在:%s\n", configPath)
	}
}
