package dto

import commonModel "goadmin/internal/common/model"

type ProductListRequest struct {
	PageIndex       *int    `json:"pageIndex" form:"pageIndex"`
	PageSize        *int    `json:"pageSize" form:"pageSize"`
	Cname           *string `json:"cname" form:"cname"`
	Abbr            *string `json:"abbr" form:"abbr"`
	Sku             *string `json:"sku" form:"sku"`
	Code            *string `json:"code" form:"code"`
	Brand           *int    `json:"brand" form:"brand"`
	Condition       *int    `json:"condition" form:"condition"`
	Algorithm       *int    `json:"algorithm" form:"algorithm"`
	ModelID         *int    `json:"model_id" form:"model_id"`
	BillingWeight   *int    `json:"billing_weight" form:"billing_weight"`
	HasPowerSource  *int    `json:"has_powersource" form:"has_powersource" binding:"omitempty,oneof=0 1"`
	HasPowerLine    *int    `json:"has_powerline" form:"has_powerline" binding:"omitempty,oneof=0 1"`
	IsMergeProduct  *int    `json:"is_merge_product" form:"is_merge_product" binding:"omitempty,oneof=0 1"`
	IsHaveInventory *int    `json:"is_have_inventory" form:"is_have_inventory" binding:"omitempty,oneof=0 1"`
	Type            *int    `json:"type" form:"type"`
}

type ProductListResponse struct {
	Total     int64                 `json:"total"`
	PageIndex int                   `json:"pageIndex"`
	PageSize  int                   `json:"pageSize"`
	List      []commonModel.Product `json:"list"`
}
