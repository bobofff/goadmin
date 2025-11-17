package model

import "time"

// Product mirrors the db_product table structure.
type Product struct {
	ID                   int        `gorm:"column:id;primaryKey"`
	SKU                  *string    `gorm:"column:sku"`
	CName                *string    `gorm:"column:cname"`
	Abbr                 *string    `gorm:"column:abbr"`
	Code                 *string    `gorm:"column:code"`
	StyleCode            *string    `gorm:"column:style_code"`
	Category             *string    `gorm:"column:category"`
	Brand                *int       `gorm:"column:brand"`
	ColorSpecs           *string    `gorm:"column:color_specs"`
	Condition            *string    `gorm:"column:condition"`
	ExpirationDate       *int       `gorm:"column:expiration_date"`
	Hashrate             *float64   `gorm:"column:hashrate"`
	ComputingUnit        *int       `gorm:"column:computing_unit"`
	Power                *int       `gorm:"column:power"`
	Weight               *float64   `gorm:"column:weight"`
	GrossWeight          *float64   `gorm:"column:gross_weight"`
	BillingWeight        *float64   `gorm:"column:billing_weight"`
	Dimensions           *string    `gorm:"column:dimensions"`
	GrossDimensions      *string    `gorm:"column:gross_dimensions"`
	Volume               *float64   `gorm:"column:volume"`
	FreightTag           *int       `gorm:"column:freight_tag"`
	ModelID              *uint64    `gorm:"column:model_id"`
	MainPhoto            *int       `gorm:"column:main_photo"`
	Photo                *string    `gorm:"column:photo"`
	Describe             *string    `gorm:"column:describe"`
	Remark               *string    `gorm:"column:remark"`
	MinNum               *int       `gorm:"column:min_num"`
	EstimatedShipDate    *int       `gorm:"column:estimated_ship_date"`
	Cost                 *float64   `gorm:"column:cost"`
	SalesRate            *float64   `gorm:"column:sales_rate"`
	SalesPriceFutures    *float64   `gorm:"column:sales_price_futures"`
	SalesPrice           *float64   `gorm:"column:sales_price"`
	FreightPrice         *float64   `gorm:"column:freight_price"`
	IsAdvance            *string    `gorm:"column:is_advance"`
	AdvanceRatio         *int       `gorm:"column:advance_ratio"`
	Creator              *int       `gorm:"column:creator"`
	CreateTime           *time.Time `gorm:"column:create_time"`
	Updater              *int       `gorm:"column:updater"`
	UpdateTime           *time.Time `gorm:"column:update_time"`
	IsDel                *int       `gorm:"column:is_del"`
	Status               *string    `gorm:"column:status"`
	Type                 *string    `gorm:"column:type"`
	ScID                 *int       `gorm:"column:sc_id"`
	Unit                 *int       `gorm:"column:unit"`
	HasPowersource       *int       `gorm:"column:has_powersource"`
	HasPowerline         *int       `gorm:"column:has_powerline"`
	PowerlineSpecs       *string    `gorm:"column:powerline_specs"`
	PowerlineStandards   *int       `gorm:"column:powerline_standards"`
	MemorySize           *int       `gorm:"column:memory_size"`
	ChipNumber           *int       `gorm:"column:chip_number"`
	WorkingMode          *string    `gorm:"column:working_mode"`
	Algorithm            *string    `gorm:"column:algorithm"`
	WorkingVoltage       *string    `gorm:"column:working_voltage"`
	NoiseLevel           *string    `gorm:"column:noise_level"`
	NetworkConnMethod    *string    `gorm:"column:network_conn_method"`
	OperationTemperature *string    `gorm:"column:operation_temperature"`
	WorkingHumidity      *string    `gorm:"column:working_humidity"`
	CollingMethod        *string    `gorm:"column:colling_method"`
	BareMacheineSize     *string    `gorm:"column:bare_macheine_size"`
	PackageSize          *string    `gorm:"column:package_size"`
	FreeShipping         int        `gorm:"column:free_shipping"`
	ApplyRules           int        `gorm:"column:apply_rules"`
	ContainBattery       *int       `gorm:"column:contain_battery"`
	HasMagnetic          *int       `gorm:"column:has_magnetic"`
	ReleaseDate          *int       `gorm:"column:release_date"`
}

// TableName returns the table name for Product.
func (Product) TableName() string {
	return "db_product"
}
