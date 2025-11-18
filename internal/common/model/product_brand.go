package model

// ProductBrand mirrors the db_product_brand table.
type ProductBrand struct {
	ID    int     `gorm:"column:id;primaryKey" json:"id"`
	Name  *string `gorm:"column:name" json:"name,omitempty"`
	IsDel int     `gorm:"column:is_del" json:"is_del"`
}

// TableName returns the brand table name.
func (ProductBrand) TableName() string {
	return "db_product_brand"
}
