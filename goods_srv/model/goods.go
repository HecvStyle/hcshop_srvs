package model

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null"`
	Level            int32  `gorm:"type:int;not null;default:1"`
	IsTab            bool   `gorm:"default:false;not null"`
	ParentCategoryID int32
	ParentCategory   *Category
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	//  这里的名有要求，必须是 tablename + id ，不然在自动建表就失败
	//BrandId int32 `gorm:"type:int;index:idx_category_brand,unique"` // 建表失败
	BrandsId int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null;"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default,not null"`
}

type Goods struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsId    int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name       string `gorm:"type:varchar(50);not null"`
	GoodsSn    string `gorm:"type:varchar(50);not null"`
	GoodsBrief string `gorm:"type:varchar(100);not null"`

	ClickNum int32 `gorm:"type:int;default:0;not null"`
	SoldNum  int32 `gorm:"type:int;default:0;not null"`
	FavNum   int32 `gorm:"type:int;default:0;not null"`

	MarketPrice float32 `gorm:"not null"`
	ShopPrice   float32 `gorm:"not null"`

	Images          GormList `gorm:"type:varchar(2000);not null"`
	DesImages       GormList `gorm:"type:varchar(2000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`
}
