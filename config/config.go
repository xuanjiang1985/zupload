package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

var Conf = new(Config)

type Config struct {
	App      string
	Version  string
	Env      string
	DataBase DataBase
}

type DataBase struct {
	Sqlite3 Sqlite3
}

type Sqlite3 struct {
	DBName string
}

func InitConfig(path string) (*Config, error) {
	if path == "" {
		viper.SetConfigFile("./config/dev.yaml")
	} else {
		viper.SetConfigFile(path)
	}

	// 加载配置
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// 解析
	if err := viper.Unmarshal(Conf); err != nil {
		return nil, err
	}

	// 热加载
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
	})

	viper.WatchConfig()
	v, _ := json.Marshal(Conf)
	fmt.Printf("%s \n", pretty.Pretty(v))
	return Conf, nil
}

func InitForTest() error {
	abPath := ""
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}

	_, err := InitConfig(abPath + string(os.PathSeparator) + "dev.yaml")
	if err != nil {
		return err
	}

	// other init

	time.Sleep(time.Millisecond * 1000)
	return nil
}
