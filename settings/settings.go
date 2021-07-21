package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AllConfig)

type AllConfig struct {
	*AppConfig   `mapstructure:"app"`
	*AuthConfig  `mapstructure:"auth"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type AuthConfig struct {
	JwtExpire int `mapstructure:"jwt_expire"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//viper.SetConfigFile("./conf/config.yaml") // 配置文件(有扩展名)
	//viper.SetConfigFile("./conf/config.json") // 配置文件(有扩展名)
	viper.SetConfigFile("./conf/config.toml") // 配置文件(有扩展名)
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("配置文件读取失败, err:%v\n", err)
		return
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("配置文件反序列化失败, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生改变时会调用此回调函数
		fmt.Println("配置文件发生了改变..:", e.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("配置文件反序列化失败, err:%v\n", err)
		}
	})
	return
}
