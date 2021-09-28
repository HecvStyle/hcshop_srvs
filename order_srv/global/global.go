package global

import (
	"gorm.io/gorm"
	"hcshop_srvs/order_srv/config"
	"hcshop_srvs/order_srv/proto"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}

	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)
