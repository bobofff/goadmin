package service

import "goadmin/internal/admin/repository"

type ProductService struct {
	// Add necessary fields here
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
