package model

import "time"

type ShoppingCart struct {
	BaseModel
	User    int32 `gorm:"type:int;index"`
	Goods   int32 `gorm:"type:int;index"`
	Nums    int32 `gorm:"type:int;index"`
	Checked bool
}

type OrderInfo struct {
	BaseModel
	User    int32  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"`
	PayType string `gorm:"type:varchar(20) comment 'alipay,wechat'"`

	Status     string     `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo    string     `gorm:"type:varchar(100) comment '交易号'"` //交易号就是支付宝的订单号 查账
	OrderMount float32    //  订单价格
	PayTime    *time.Time `gorm:"type:datetime"`

	// 快递等服务应该也是一个独立的服务
	Address      string `gorm:"type:varchar(100)"` // 地址
	SignerName   string `gorm:"type:varchar(20)"`  // 下单人
	SingerMobile string `gorm:"type:varchar(11)"`  // 下单手机号
	Post         string `gorm:"type:varchar(20)"`  // 备注
}

// OrderGoods 商品信息关联到了order订单，商品信息做了冗余，相当于上镜像，保存下单时刻的商品信息，
type OrderGoods struct {
	BaseModel
	Order int32 `gorm:"type:int;index"`
	Goods int32 `gorm:"type:int;index"`

	GoodsName  string `gorm:"type:varchar(100)"` // 商品名不做索引，回头看情况再说，这里和老师不一样
	GoodsImage string `gorm:"type:varchar(200)"`
	GoodsPrice float32
	Nums       int32 `gorm:"type:int"`
}
