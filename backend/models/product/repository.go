package product

import (
	"fmt"

	"github.com/schantz/web/go-api-produtos/backend/database"
	"gorm.io/gorm"
)

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	//GetBestSellers() ([]*ProductTop, error)
	//GetTotalVentas() (float64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(databaseConnection *gorm.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {

	product := &Product{}

	//database.ReturnDB().First(&product, productId)

	database.ReturnDB().Raw("SELECT id,product_code,product_name,COALESCE(description,''), standard_cost, list_price, category FROM products WHERE id = ?", productId).Scan(&product)
	fmt.Println("1teste:", product)

	return product, nil
}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	product := []*Product{}
	fmt.Println(params.Limit, params.Offset)

	/*
			json
			{
		    "limit": 10,
		    "offset":0
		}
	*/

	//database.ReturnDB().Select([]string{"id", "product_code", "product_name", "COALESCE(description,' ')", "standard_cost", "list_price", "category"}).Find(&product)

	database.ReturnDB().Raw("SELECT id,product_code,product_name,COALESCE(description,''), standard_cost,list_price, category FROM products	LIMIT ? OFFSET ?", params.Limit, params.Offset).Scan(&product)
	fmt.Println("2teste:", product)
	for _, v := range product {
		fmt.Println(v.Id)
	}

	return product, nil

}

func (repo *repository) GetTotalProducts() (int, error) {

	//database.ReturnDB().Select([]string{"COUNT(*)"}).Find(&product)

	type Total struct {
		sum_total int
	}

	var total Total

	// Raw SQL
	database.ReturnDB().Raw("SELECT COUNT(*) FROM products;").Scan(&total.sum_total)
	fmt.Println("total:", total.sum_total)

	return total.sum_total, nil
}

func (repo *repository) InsertProduct(params *getAddProductRequest) (int64, error) {

	var product Product

	database.ReturnDB().Select(params.ProductCode, ", product_name, category, description, list_price, standard_cost").Create(&product)

	database.ReturnDB().Last(product.Id)
	fmt.Println("4teste:", product)

	return int64(product.Id), nil
}

func (repo *repository) UpdateProduct(params *updateProductRequest) (int64, error) {
	var product Product

	database.ReturnDB().Model(&product).Where("id = ?", params.ID).Updates(Product{
		ProductName:  product.ProductName,
		Category:     product.Category,
		ProductCode:  product.ProductCode,
		ListPrice:    product.ListPrice,
		StandardCost: product.StandardCost,
		Description:  product.Description})
	fmt.Println("5teste:", product)

	return params.ID, nil
}

func (repo *repository) DeleteProduct(params *deleteProductRequest) (int64, error) {
	product := Product{}

	database.ReturnDB().Where("id = ?", params.ProductID).Delete(&product)
	fmt.Println("6teste:", product)

	return 1, nil
}

/*
func (repo *repository) GetBestSellers() ([]*ProductTop, error) {
	const sql = `SELECT
	                od.product_id,
					p.product_name,
					SUM(od.quantity*od.unit_price) vendido
				FROM order_details od
				inner join products p on od.product_id = p.id
				GROUP by od.product_id
				ORDER BY vendido desc
				limit 10`
	results, err := repo.db.Query(sql)
	utils.ErrorPanic(err)

	var products []*ProductTop
	for results.Next() {
		product := &ProductTop{}
		err = results.Scan(&product.ID, &product.ProductName, &product.Vendidos)
		utils.ErrorPanic(err)
		products = append(products, product)
	}

	return products, nil
}
func (repo *repository) GetTotalVentas() (float64, error) {
	const sql = `SELECT SUM(od.quantity*od.unit_price) vendido
	FROM order_details od`
	var total float64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	utils.ErrorPanic(err)
	return total, nil
}
*/
