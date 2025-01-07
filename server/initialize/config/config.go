package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"nexwind.net/admin/server/constant"
)

// Conf system conf struct.
type Conf struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	SysDb    MySQLConf   `json:"sys_db"`
	MySQLSet []MySQLConf `json:"mySQLSet,omitempty"`
	RedisSet []RedisConf `json:"redisSet,omitempty"`
	IsDev    bool        `json:"isDev"`
	Env      string      `json:"env"`
	Cors     []string    `json:"cors,omitempty"`
}

// InitConfig init config form config file.
func InitConfig(d *bool, env string, logger *zap.Logger) *Conf {
	var (
		configPath = fmt.Sprintf("./config/v1/%s", constant.CONFIG_FILE[env])
		conf       = &Conf{}
	)
	if !*d {
		configPath = constant.CONFIG_FILE[env]
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("An error occurred while reading configuration file", zap.Any("file", configPath))
	}
	if err := viper.Unmarshal(&conf); err != nil {
		logger.Error("An error occurred while parse configuration file", zap.Any("file", configPath))
	}
	conf.IsDev = *d
	conf.Env = env
	return conf
}
