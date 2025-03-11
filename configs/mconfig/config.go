package mconfig

import (
	"errors"
	"github.com/spf13/viper"
)

// OnChange 配置变化方法
type OnChange func(v *viper.Viper)

// IProvider 接口 工厂模式是一种创建对象的设计模式，它将对象的创建和使用分离
type IProvider interface {
	GetViper() *viper.Viper
	Watch(change OnChange) error
}

var mfr IProvider

// Init 初始化配置
func Init() (err error) {
	if configFilePath != "" {
		mfr, err = newFileProvider()
	} else {
		err = errors.New("配置路径为空，其他工厂模式暂不支持")
	}
	// 合并viper配置
	if err == nil {
		viper.MergeConfigMap(mfr.GetViper().AllSettings())
	}
	return
}
