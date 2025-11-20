package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"goadmin/internal/admin/dto"
	"goadmin/internal/admin/repository"
	commonModel "goadmin/internal/common/model"
	"goadmin/pkg/utils"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

const (
	defaultProductPageIndex = 1
	defaultProductPageSize  = 20
	maxProductPageSize      = 100
)

// ErrInvalidProductInput indicates validation failure for product creation.
var ErrInvalidProductInput = errors.New("invalid product input")

func (service *ProductService) GetList(ctx context.Context, request dto.ProductListRequest) (*dto.ProductListResponse, error) {
	pageIndex := defaultProductPageIndex
	if request.PageIndex != nil && *request.PageIndex > 0 {
		pageIndex = *request.PageIndex
	}

	pageSize := defaultProductPageSize
	if request.PageSize != nil && *request.PageSize > 0 {
		pageSize = *request.PageSize
		if pageSize > maxProductPageSize {
			pageSize = maxProductPageSize
		}
	}

	params := repository.ProductListParams{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Filters: repository.ProductListFilters{
			Cname:           request.Cname,
			Abbr:            request.Abbr,
			Sku:             request.Sku,
			Code:            request.Code,
			Brand:           request.Brand,
			Condition:       request.Condition,
			Algorithm:       request.Algorithm,
			ModelID:         request.ModelID,
			BillingWeight:   request.BillingWeight,
			HasPowerSource:  request.HasPowerSource,
			HasPowerLine:    request.HasPowerLine,
			IsMergeProduct:  request.IsMergeProduct,
			IsHaveInventory: request.IsHaveInventory,
			Type:            request.Type,
		},
	}

	list, total, err := service.repo.List(ctx, params)
	if err != nil {
		return nil, err
	}

	brandNames := map[int]string{}
	if len(list) > 0 {
		brandIDs := extractBrandIDs(list)
		if len(brandIDs) > 0 {
			brandNames, err = service.repo.GetBrandNames(ctx, brandIDs)
			if err != nil {
				return nil, err
			}
		}
	}

	for i := range list {
		if list[i].Brand != nil {
			if name, ok := brandNames[*list[i].Brand]; ok {
				list[i].BrandValue = stringPtr(name)
			}
		}
	}

	return &dto.ProductListResponse{
		Total:     total,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		List:      list,
	}, nil
}

func extractBrandIDs(list []commonModel.Product) []int {
	idsMap := make(map[int]struct{})
	for _, item := range list {
		if item.Brand != nil && *item.Brand > 0 {
			idsMap[*item.Brand] = struct{}{}
		}
	}
	result := make([]int, 0, len(idsMap))
	for id := range idsMap {
		result = append(result, id)
	}
	return result
}

func stringPtr(value string) *string {
	return &value
}

// Insert creates a new product record.
func (service *ProductService) Insert(ctx context.Context, request dto.ProductCreateRequest) (*commonModel.Product, error) {
	sku := sanitizeStringPtr(request.Sku)
	cname := sanitizeStringPtr(request.Cname)
	if sku == nil || cname == nil {
		return nil, fmt.Errorf("%w: 产品名称或SKU不能为空", ErrInvalidProductInput)
	}

	product := &commonModel.Product{
		SKU:            sku,
		CName:          cname,
		Abbr:           sanitizeStringPtr(request.Abbr),
		Code:           sanitizeStringPtr(request.Code),
		Brand:          request.Brand,
		Condition:      intToStringPtr(request.Condition),
		Algorithm:      intToStringPtr(request.Algorithm),
		ModelID:        intToUint64Ptr(request.ModelID),
		BillingWeight:  intToFloat64Ptr(request.BillingWeight),
		HasPowersource: request.HasPowerSource,
		HasPowerline:   request.HasPowerLine,
		Type:           intToStringPtr(request.Type),
		Creator:        request.Creator,
	}

	now := time.Now()
	product.CreateTime = &now
	product.UpdateTime = &now
	product.IsDel = intPtr(0)

	if err := service.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func sanitizeStringPtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := utils.SafeString(*value)
	if trimmed == "" {
		return nil
	}
	return stringPtr(trimmed)
}

func intPtr(value int) *int {
	return &value
}

func intToStringPtr(value *int) *string {
	if value == nil {
		return nil
	}
	strValue := strconv.Itoa(*value)
	return &strValue
}

func intToFloat64Ptr(value *int) *float64 {
	if value == nil {
		return nil
	}
	floatValue := float64(*value)
	return &floatValue
}

func intToUint64Ptr(value *int) *uint64 {
	if value == nil {
		return nil
	}
	uintValue := uint64(*value)
	return &uintValue
}
