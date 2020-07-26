package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	_ "path/filepath"
	"smiling-blog/tools/slog"
	"time"
)

var Configs *Config
var YamlFile []byte

type Config struct {
	DataBase DataBaseConfig `yaml:"dataSource"`
	Redis    redisConfig    `yaml:"qr_redis"`
	App      AppConfig      `yaml:"app"`
	Log      slog.LogConfig `yaml:"log"`
}

/**
 * database config
 */
type DataBaseConfig struct {
	Host        string        `yaml:"host"`
	User        string        `yaml:"user"`
	Pwd         string        `yaml:"password"`
	DbName      string        `yaml:"database"`
	Port        int           `yaml:"port"`
	MaxIdleCons int           `yaml:"maxIdleConns"`
	MaxOpenCons int           `yaml:"maxOpenConns"`
	MaxLifeTime time.Duration `yaml:"maxLifeTime"`
}

/**
 * redis config
 */
type redisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

/**
 * app config
 */
type AppConfig struct {
	Addr  string `yaml:"addr"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

func Init(fileName string) {
	var err error
	YamlFile, err = ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("load conf error, will exit")
		os.Exit(0)
	}
	conf := &Config{}
	err = yaml.Unmarshal(YamlFile, conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	Configs = conf
}

func GetAppAddr() string {
	return fmt.Sprintf("%s:%d", Configs.App.Addr, Configs.App.Port)
}
