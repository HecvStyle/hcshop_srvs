package config

type MySqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	Password string `mapstructure:"password" json:"password"`
	User     string `mapstructure:"user" json:"user"`
}

type ServerConfig struct {
	Name      string      `mapstructure:"name"`
	Port      int         `mapstructure:"port"`
	MysqlInfo MySqlConfig `mapstructure:"mysql" json:"mysql"`
}
