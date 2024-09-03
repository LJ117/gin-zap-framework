package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

/* Manage config file by Viper*/

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // 项目根目录下
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed:%v\n", err)
		return err
	}
	fmt.Printf("The config viper load success, the value is %v\n", viper.AllSettings())

	// hot load
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return nil
}
