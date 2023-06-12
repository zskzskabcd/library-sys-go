package config

import (
	"os"

	"github.com/jinzhu/configor"
	"gopkg.in/yaml.v2"
)

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Charset  string `yaml:"charset"`
}
type RedisConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type DatabaseConf struct {
	Mysql MysqlConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
}

type Config struct {
	Database DatabaseConf `yaml:"database"`
}

var myConfig Config

func init() {
	myConfig = Config{
		Database: DatabaseConf{
			Mysql: MysqlConf{
				Host:     "127.0.0.1",
				Port:     "3306",
				User:     "root",
				Password: "123456",
				DbName:   "library_sys_go",
				Charset:  "utf8mb4",
			},
			Redis: RedisConf{
				Host:     "127.0.0.1",
				Port:     "6379",
				Password: "",
				Db:       2,
			},
		},
	}
	// 是否存在配置文件
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		// 生成yaml
		data_router, err := yaml.Marshal(myConfig)
		if err != nil {
			panic(err)
		}
		os.WriteFile("config.yaml", data_router, 0666)
	}

	err = configor.Load(&myConfig, "config.yaml")
	if err != nil {
		return
	}
}

func GetConfig() Config {
	return myConfig
}

func GetMysqlConfig() MysqlConf {
	return myConfig.Database.Mysql
}

func (m *MysqlConf) GetMysqlDSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}
