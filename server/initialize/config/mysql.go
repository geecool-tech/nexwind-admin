package config

// MySQLConf MySQL conf.
type MySQLConf struct {
	Key      string `json:"key"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}
