package repository

import (
	"context"

	"gorm.io/gorm"

	commonModel "goadmin/internal/common/model"
	"goadmin/pkg/utils"
)

// ProductRepository defines access methods for products.
type ProductRepository interface {
	List(ctx context.Context, params ProductListParams) ([]commonModel.Product, int64, error)
	GetBrandNames(ctx context.Context, ids []int) (map[int]string, error)
	Create(ctx context.Context, product *commonModel.Product) error
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a repository backed by GORM.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// ProductListParams contains pagination and filter info for list queries.
type ProductListParams struct {
	PageIndex int
	PageSize  int
	Filters   ProductListFilters
}

// ProductListFilters defines optional filters for listing products.
type ProductListFilters struct {
	Cname           *string
	Abbr            *string
	Sku             *string
	Code            *string
	Brand           *int
	Condition       *int
	Algorithm       *int
	ModelID         *int
	BillingWeight   *int
	HasPowerSource  *int
	HasPowerLine    *int
	IsMergeProduct  *int
	IsHaveInventory *int
	Type            *int
}

func (r *productRepository) List(ctx context.Context, params ProductListParams) ([]commonModel.Product, int64, error) {
	query := r.db.WithContext(ctx).Model(&commonModel.Product{}).Where("is_del = 0")
	filters := params.Filters

	if v := normalizeString(filters.Cname); v != "" {
		query = query.Where("cname LIKE ?", likeValue(v))
	}
	if v := normalizeString(filters.Abbr); v != "" {
		query = query.Where("abbr LIKE ?", likeValue(v))
	}
	if v := normalizeString(filters.Sku); v != "" {
		query = query.Where("sku LIKE ?", likeValue(v))
	}
	if v := normalizeString(filters.Code); v != "" {
		query = query.Where("code LIKE ?", likeValue(v))
	}
	if filters.Brand != nil {
		query = query.Where("brand = ?", *filters.Brand)
	}
	if filters.Condition != nil {
		query = query.Where("`condition` = ?", *filters.Condition)
	}
	if filters.Algorithm != nil {
		query = query.Where("algorithm = ?", *filters.Algorithm)
	}
	if filters.ModelID != nil {
		query = query.Where("model_id = ?", *filters.ModelID)
	}
	if filters.BillingWeight != nil {
		query = query.Where("billing_weight = ?", *filters.BillingWeight)
	}
	if filters.HasPowerSource != nil {
		query = query.Where("has_powersource = ?", *filters.HasPowerSource)
	}
	if filters.HasPowerLine != nil {
		query = query.Where("has_powerline = ?", *filters.HasPowerLine)
	}
	if filters.IsMergeProduct != nil {
		sub := r.db.Table("db_product_relation").Select("pid").Where("is_del = 0")
		if *filters.IsMergeProduct == 1 {
			query = query.Where("id IN (?)", sub)
		} else {
			query = query.Where("id NOT IN (?)", sub)
		}
	}
	if filters.IsHaveInventory != nil {
		sub := r.db.Table("db_inventory").Select("productid").Where("is_del = 0")
		if *filters.IsHaveInventory == 1 {
			query = query.Where("id IN (?)", sub)
		} else {
			query = query.Where("id NOT IN (?)", sub)
		}
	}
	if filters.Type != nil {
		query = query.Where("type = ?", *filters.Type)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []commonModel.Product{}, 0, nil
	}

	offset := (params.PageIndex - 1) * params.PageSize
	if offset < 0 {
		offset = 0
	}
	limit := params.PageSize
	if limit <= 0 {
		limit = 20
	}

	var result []commonModel.Product
	err := query.Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func normalizeString(value *string) string {
	if value == nil {
		return ""
	}
	return utils.SafeString(*value)
}

func likeValue(value string) string {
	return "%" + value + "%"
}

func (r *productRepository) GetBrandNames(ctx context.Context, ids []int) (map[int]string, error) {
	result := make(map[int]string, len(ids))
	if len(ids) == 0 {
		return result, nil
	}

	var brands []commonModel.ProductBrand
	err := r.db.WithContext(ctx).
		Model(&commonModel.ProductBrand{}).
		Where("is_del = 0").
		Where("id IN ?", ids).
		Find(&brands).Error
	if err != nil {
		return nil, err
	}

	for _, brand := range brands {
		if brand.Name != nil {
			result[brand.ID] = *brand.Name
		}
	}

	return result, nil
}

func (r *productRepository) Create(ctx context.Context, product *commonModel.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}
