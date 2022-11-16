package crawler

import (
	"fmt"
	"strconv"

	"gitlab.com/eiprice/crawlers/james/domain"
	"gitlab.com/eiprice/crawlers/james/repositories"
	"gitlab.com/eiprice/crawlers/james/utils"
)

type ProductCrawler interface {
	GetData(categories []*domain.Category, typeStore *domain.TypeStore)
	GetById(category *domain.Category, typeStore *domain.TypeStore) ([]*domain.Product, error)
}

type ProductCrawlerInit struct {
	ProductRepository repositories.ProductRepository
	List              string
}

func (craw *ProductCrawlerInit) GetData(categories []*domain.Category, typeStore *domain.TypeStore) {

	for _, item := range categories {
		fmt.Println("Get product from Category: ", item.Name)
		fmt.Println("SubCategory: ", item.SubName)
		product, err := craw.GetById(item, typeStore)
		craw.ProductRepository.InsertMany(product)

		if err != nil {
			fmt.Println("Error to get Products  from Category:", item.Name)
		}

		fmt.Println("Total products inserted from Category "+item.Name+": ", len(product))
	}

}

func (craw *ProductCrawlerInit) GetById(category *domain.Category, typeStore *domain.TypeStore) ([]*domain.Product, error) {

	var dados interface{}
	var product []*domain.Product

	url := "https://api.james.delivery/api/v2/customer/menu/" + strconv.Itoa(category.StoreID) + "?category_id=" + strconv.Itoa(category.SubCategoryID) + "&offset=0&limit=1000000&only_one_aisle=false"

	dados, err := utils.Request("GET", url, "", nil)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["aisles"].([]interface{}) {
		for _, subItem := range item.(map[string]interface{})["products"].([]interface{}) {
			productID := utils.IntNotNull(subItem.(map[string]interface{})["id"])

			obj, _ := domain.NewProduct(
				productID,
				"0",
				true,
				category.CategoryID,
				category.Name,
				category.SubCategoryID,
				category.SubName,
				utils.StringNotNull(subItem.(map[string]interface{})["title"]),
				utils.StringNotNull(subItem.(map[string]interface{})["photo"].(map[string]interface{})["medium"].(map[string]interface{})["url"]),
				true,
				utils.StringNotNull(subItem.(map[string]interface{})["title"]),
				utils.FloatNotNull(subItem.(map[string]interface{})["price"]),
				utils.IntNotNull(subItem.(map[string]interface{})["promotion_customer_quantity_limit"]),
				utils.FloatNotNull(subItem.(map[string]interface{})["original_price"]),
				category.StoreID,
				category.StoreName,
				category.StoreAddress,
				typeStore.Name,
				typeStore.Name,
				craw.List,
				`https://api.james.delivery/api/v1/menus/`+strconv.Itoa(productID),
			)
			product = append(product, obj)

		}
	}

	return product, nil
}
