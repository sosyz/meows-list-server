package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

var Config *config
var Logger *logger

func init() {
	Config = &config{}
	Logger = &logger{
		level: LevelInfo,
	}
	err := ReadConfig()
	if err != nil {
		Logger.Error("Read config failed, err:", err)
		return
	}
	switch Config.RunConfig.LogLevel {
	case "debug":
		Logger.level = LevelDebug
	case "info":
		Logger.level = LevelInfo
	case "warn":
		Logger.level = LevelWaring
	case "error":
		Logger.level = LevelError
	default:
		Logger.level = LevelInfo
	}
}

func ReadConfig() error {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		// 从文件读取配置失败 尝试从环境变量获取
		Logger.Waring("Read config from env failed, err: %v, try to read from env", err)
		//设置环境变量名前缀
		return ReadConfigFromEnv(Config, "MEOWS")
	} else {
		if err := v.Unmarshal(Config); err != nil {
			return err
		}
	}
	return nil
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
				Logger.Waring("can't read [%s] value", envPrefix+"_"+f.Name)
			} else {
				// 修改其值
				objValue.Elem().Field(i).SetString(val)
			}
		}
	}
	return nil
}
