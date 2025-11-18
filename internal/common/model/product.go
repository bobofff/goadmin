package model

import "time"

// Product mirrors the db_product table structure.
type Product struct {
	ID                   int        `gorm:"column:id;primaryKey" json:"id"`
	SKU                  *string    `gorm:"column:sku" json:"sku,omitempty"`
	CName                *string    `gorm:"column:cname" json:"cname,omitempty"`
	Abbr                 *string    `gorm:"column:abbr" json:"abbr,omitempty"`
	Code                 *string    `gorm:"column:code" json:"code,omitempty"`
	StyleCode            *string    `gorm:"column:style_code" json:"style_code,omitempty"`
	Category             *string    `gorm:"column:category" json:"category,omitempty"`
	Brand                *int       `gorm:"column:brand" json:"brand,omitempty"`
	ColorSpecs           *string    `gorm:"column:color_specs" json:"color_specs,omitempty"`
	Condition            *string    `gorm:"column:condition" json:"condition,omitempty"`
	ExpirationDate       *int       `gorm:"column:expiration_date" json:"expiration_date,omitempty"`
	Hashrate             *float64   `gorm:"column:hashrate" json:"hashrate,omitempty"`
	ComputingUnit        *int       `gorm:"column:computing_unit" json:"computing_unit,omitempty"`
	Power                *int       `gorm:"column:power" json:"power,omitempty"`
	Weight               *float64   `gorm:"column:weight" json:"weight,omitempty"`
	GrossWeight          *float64   `gorm:"column:gross_weight" json:"gross_weight,omitempty"`
	BillingWeight        *float64   `gorm:"column:billing_weight" json:"billing_weight,omitempty"`
	Dimensions           *string    `gorm:"column:dimensions" json:"dimensions,omitempty"`
	GrossDimensions      *string    `gorm:"column:gross_dimensions" json:"gross_dimensions,omitempty"`
	Volume               *float64   `gorm:"column:volume" json:"volume,omitempty"`
	FreightTag           *int       `gorm:"column:freight_tag" json:"freight_tag,omitempty"`
	ModelID              *uint64    `gorm:"column:model_id" json:"model_id,omitempty"`
	MainPhoto            *int       `gorm:"column:main_photo" json:"main_photo,omitempty"`
	Photo                *string    `gorm:"column:photo" json:"photo,omitempty"`
	Describe             *string    `gorm:"column:describe" json:"describe,omitempty"`
	Remark               *string    `gorm:"column:remark" json:"remark,omitempty"`
	MinNum               *int       `gorm:"column:min_num" json:"min_num,omitempty"`
	EstimatedShipDate    *int       `gorm:"column:estimated_ship_date" json:"estimated_ship_date,omitempty"`
	Cost                 *float64   `gorm:"column:cost" json:"cost,omitempty"`
	SalesRate            *float64   `gorm:"column:sales_rate" json:"sales_rate,omitempty"`
	SalesPriceFutures    *float64   `gorm:"column:sales_price_futures" json:"sales_price_futures,omitempty"`
	SalesPrice           *float64   `gorm:"column:sales_price" json:"sales_price,omitempty"`
	FreightPrice         *float64   `gorm:"column:freight_price" json:"freight_price,omitempty"`
	IsAdvance            *string    `gorm:"column:is_advance" json:"is_advance,omitempty"`
	AdvanceRatio         *int       `gorm:"column:advance_ratio" json:"advance_ratio,omitempty"`
	Creator              *int       `gorm:"column:creator" json:"creator,omitempty"`
	CreateTime           *time.Time `gorm:"column:create_time" json:"create_time,omitempty"`
	Updater              *int       `gorm:"column:updater" json:"updater,omitempty"`
	UpdateTime           *time.Time `gorm:"column:update_time" json:"update_time,omitempty"`
	IsDel                *int       `gorm:"column:is_del" json:"is_del,omitempty"`
	Status               *string    `gorm:"column:status" json:"status,omitempty"`
	Type                 *string    `gorm:"column:type" json:"type,omitempty"`
	ScID                 *int       `gorm:"column:sc_id" json:"sc_id,omitempty"`
	Unit                 *int       `gorm:"column:unit" json:"unit,omitempty"`
	HasPowersource       *int       `gorm:"column:has_powersource" json:"has_powersource,omitempty"`
	HasPowerline         *int       `gorm:"column:has_powerline" json:"has_powerline,omitempty"`
	PowerlineSpecs       *string    `gorm:"column:powerline_specs" json:"powerline_specs,omitempty"`
	PowerlineStandards   *int       `gorm:"column:powerline_standards" json:"powerline_standards,omitempty"`
	MemorySize           *int       `gorm:"column:memory_size" json:"memory_size,omitempty"`
	ChipNumber           *int       `gorm:"column:chip_number" json:"chip_number,omitempty"`
	WorkingMode          *string    `gorm:"column:working_mode" json:"working_mode,omitempty"`
	Algorithm            *string    `gorm:"column:algorithm" json:"algorithm,omitempty"`
	WorkingVoltage       *string    `gorm:"column:working_voltage" json:"working_voltage,omitempty"`
	NoiseLevel           *string    `gorm:"column:noise_level" json:"noise_level,omitempty"`
	NetworkConnMethod    *string    `gorm:"column:network_conn_method" json:"network_conn_method,omitempty"`
	OperationTemperature *string    `gorm:"column:operation_temperature" json:"operation_temperature,omitempty"`
	WorkingHumidity      *string    `gorm:"column:working_humidity" json:"working_humidity,omitempty"`
	CollingMethod        *string    `gorm:"column:colling_method" json:"colling_method,omitempty"`
	BareMacheineSize     *string    `gorm:"column:bare_macheine_size" json:"bare_macheine_size,omitempty"`
	PackageSize          *string    `gorm:"column:package_size" json:"package_size,omitempty"`
	FreeShipping         int        `gorm:"column:free_shipping" json:"free_shipping"`
	ApplyRules           int        `gorm:"column:apply_rules" json:"apply_rules"`
	ContainBattery       *int       `gorm:"column:contain_battery" json:"contain_battery,omitempty"`
	HasMagnetic          *int       `gorm:"column:has_magnetic" json:"has_magnetic,omitempty"`
	ReleaseDate          *int       `gorm:"column:release_date" json:"release_date,omitempty"`
	BrandValue           *string    `gorm:"-" json:"brand_value,omitempty"`
}

// TableName returns the table name for Product.
func (Product) TableName() string {
	return "db_product"
}
