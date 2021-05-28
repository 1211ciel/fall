package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	C     = new(ViperConfig)
	V     *viper.Viper

)

func NewViper(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%v \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(v.Get("time.status"))
		color.Green("config changed ")
		if err := v.Unmarshal(&C); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%v \n", err))
		}
		color.Red("%v", C.Time.Status)

	})
	V = v
}

// ViperConfig mapstructure
type ViperConfig struct {
	Time struct {
		Status int `mapstructure:"status"`
	} `mapstructure:"time"`
	Mysql struct {
		DataSource string `mapstructure:"DataSource"`
	} `mapstructure:"Mysql"`
	CacheRedis struct {
		Host string `mapstructure:"Host"`
	} `mapstructure:"CacheRedis"`
}


