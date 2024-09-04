package settings

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// SysCfg
// global variable
// use to save config info
var SysCfg = new(SystemConfig)

/* Manage config file by Viper*/

func Init() (err error) {

	// 方式1: 直接指定配置文件路径（相对路径或绝对路径）
	// 相对路径
	viper.SetConfigFile("config/config.toml")
	// 绝对路径
	//viper.SetConfigFile("/home/seven/LearnSpace/Bilibil/Qimi/web_app/config/config.yaml")

	// 方式2: 指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	//viper.SetConfigName("config")

	// 配置文件位置可配置多个
	//viper.AddConfigPath(".")
	//viper.AddConfigPath("./cong")

	// 专用于从远程获取配置信息时指定配置文件类型, 本地不生效
	// 配合远程配置中心使用，告诉viper当前的数据使用什么格式解析
	//viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed:%v\n", err)
		return err
	}

	// init unmarshal
	if err = viper.Unmarshal(SysCfg); err != nil {
		zap.L().Error("viper.Unmarshal(SysCfg) failed", zap.Error(err))
		return
	}

	// hot load
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed...")
		// unmarshal again
		if err = viper.Unmarshal(SysCfg); err != nil {
			zap.L().Error("viper.Unmarshal(SysCfg) failed", zap.Error(err))
			return
		}
	})
	return nil
}
