package global

import (
	"gorm.io/gorm"
	"hcshop_srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
)
