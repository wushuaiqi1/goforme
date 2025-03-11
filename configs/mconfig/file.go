package mconfig

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configFilePath string

func init() {
	log.Info("file init")
	flag.StringVar(&configFilePath, "config", "", "配置文件路径")
}

type fileProvider struct {
	v *viper.Viper
}

// NewFileProvider 创建提供者
func newFileProvider() (provider IProvider, err error) {
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigFile(configFilePath)
	// 调用viper读取配置函数
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	f := &fileProvider{
		v: v,
	}
	// 接口实现者
	provider = f
	return
}

func (f *fileProvider) Watch(change OnChange) (err error) {
	f.v.OnConfigChange(func(in fsnotify.Event) {
		switch in.Op {
		case fsnotify.Create, fsnotify.Write:
			change(f.v)
		}
	})
	go f.v.WatchConfig()
	return
}

func (f *fileProvider) GetViper() *viper.Viper {
	return f.v
}
