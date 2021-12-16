package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/schantz/web/go-api-produtos/backend/utils"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct{}

// @Summary Producto by Id
// @Tags Producto
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {object} product.Product "ok"
// @Router /products/{id} [get]
func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		utils.ErrorPanic(err)
		return product, nil
	}

	return getProductByIdEndpoint
}

// @Summary Lista de Productos
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.getProductsRequest true "User Data"
// @Success 200 {object} product.ProductList "ok"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		utils.ErrorPanic(err)
		return result, nil
	}
	return getProductsEndPoint
}

// @Summary Insertar Productos
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.getAddProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [post]
func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		utils.ErrorPanic(err)
		return productId, nil
	}

	return addProductEndpoint
}

// @Summary Update Producto
// @Tags Producto
// @Accept json
// @Produce  json
// @Param request body product.updateProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [put]
func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProduct(&req)
		utils.ErrorPanic(err)
		return r, nil
	}
	return updateProductEndpoint
}

// @Summary Eliminar Producto
// @Tags Producto
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {integer} int "ok"
// @Router /products/{id} [delete]
func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		utils.ErrorPanic(err)
		return result, nil
	}
	return deleteProductEndPoint
}
