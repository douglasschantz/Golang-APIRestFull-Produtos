package product

import "github.com/schantz/web/go-api-produtos/backend/utils"

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)
	utils.ErrorPanic(err)
	totalProducts, err := s.repo.GetTotalProducts()
	utils.ErrorPanic(err)
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}
