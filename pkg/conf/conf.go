package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"sonui.cn/meows-list-server/pkg/logger"
	"strings"
)

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type run struct {
	Host     string
	Port     string
	LogLevel string
}

type redis struct {
	Host     string
	Port     string
	Password string
	Database int
}

type config struct {
	Run      run
	Database database
	Redis    redis
}

var Run *run
var Database *database
var Redis *redis
var conf config

func Init(path string) {
	ReadConfig(path)

	Run = &conf.Run
	Database = &conf.Database
	Redis = &conf.Redis
}

func ReadConfig(path string) {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath(path)
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		// 从文件读取配置失败 尝试从环境变量获取
		logger.Waring("Read config from env failed, err: %v, try to read from env", err)
		//设置环境变量名前缀
		err = ReadConfigFromEnv(conf, "MEOWS")
		if err != nil {
			logger.Error("Read config from env failed, err: %v", err)
		}
	} else {
		if err := v.Unmarshal(&conf); err != nil {
			logger.Error("unmarshal config failed, err: %v", err)
		} else {
			logger.Debug("Read config from file success")
		}
	}
}

// ReadConfigFromEnv 从环境变量读取配置
func ReadConfigFromEnv(obj interface{}, envPrefix string) error {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	if objType.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	objType = objType.Elem()
	for i := 0; i < objType.NumField(); i++ {
		f := objType.Field(i)
		// 判断类型是否属于interface
		if f.Type.Kind() == reflect.Struct {
			// 如果是结构体，递归调用 传递指针
			if err := ReadConfigFromEnv(objValue.Elem().Field(i).Addr().Interface(), strings.ToUpper(envPrefix+"_"+f.Name)); err != nil {
				return err
			}
		} else {
			// 获取环境变量值
			val := os.Getenv(strings.ToUpper(envPrefix + "_" + f.Name))
			if val == "" {
				logger.Waring("can't read [%s] value", envPrefix+"_"+f.Name)
			} else {
				// 修改其值
				objValue.Elem().Field(i).SetString(val)
			}
		}
	}
	return nil
}
