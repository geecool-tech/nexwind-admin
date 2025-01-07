package config

// RedisConf redis conf struct.
type RedisConf struct {
	Key  string `json:"key"`
	Port int    `json:"port"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
}
