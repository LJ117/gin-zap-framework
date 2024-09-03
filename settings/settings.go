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
	viper.SetConfigFile("config/config.toml") // ab path, 本地使用建议这种形式, 不受 viper.AddConfigPath() 的影响
	//viper.SetConfigType("yaml")
	//viper.SetConfigName("config")// 只是找同名文件不包含后缀, 因此可能会有文件冲突存在, 同一目录下 config.yaml & config.toml 同时存在会报错
	//viper.SetConfigType("toml")// 专用于从远程获取配置信息时指定配置文件类型, 本地不生效, 本地要指定后缀使用 viper.SetConfigFile()
	//viper.SetConfigName("config-toml")
	viper.AddConfigPath("config") // the project's root path
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
