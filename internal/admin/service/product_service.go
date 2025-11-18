package service

import (
	"context"

	"goadmin/internal/admin/dto"
	"goadmin/internal/admin/repository"
	commonModel "goadmin/internal/common/model"
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
