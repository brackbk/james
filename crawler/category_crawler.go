package crawler

import (
	"fmt"
	"strconv"

	"gitlab.com/eiprice/crawlers/james/domain"
	"gitlab.com/eiprice/crawlers/james/repositories"
	"gitlab.com/eiprice/crawlers/james/utils"
)

type CategoryCrawler interface {
	GetData(stores []*domain.Store) ([]*domain.Category, error)
	GetCategoryById(stores *domain.Store) ([]*domain.Category, error)
	GetSubCategoryById(store *domain.Store, categoryID int, title string) ([]*domain.Category, error)
}
type CategoryCrawlerInit struct {
	CategoryRepository repositories.CategoryRepository
	Category           string
	List               string
}

func (craw *CategoryCrawlerInit) GetData(stores []*domain.Store) ([]*domain.Category, error) {
	var storeCategories []*domain.Category

	for _, item := range stores {

		fmt.Println("Get Categories from store: ", item.Name)

		categories, err := craw.GetCategoryById(item)
		if err != nil {
			fmt.Println("Error to get Categories from Store:", item.Name)
		}
		craw.CategoryRepository.InsertMany(categories)
		fmt.Println("Total Categories inserted from Store "+item.Name+": ", len(categories))
		storeCategories = append(storeCategories, categories...)
	}

	return storeCategories, nil
}

func (craw *CategoryCrawlerInit) GetCategoryById(store *domain.Store) ([]*domain.Category, error) {

	var dados interface{}

	var category []*domain.Category
	var categories []*domain.Category

	url := "https://api.james.delivery/api/v2/customer/menu/" + strconv.Itoa(store.StoreID) + "?offset=0&limit=10"

	dados, err := utils.Request("GET", url, "", nil)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["categories"].([]interface{}) {

		dado := item.(map[string]interface{})

		categoryID := utils.IntNotNull(dado["id"])
		title := utils.StringNotNull(dado["title"])

		category, _ = craw.GetSubCategoryById(store, categoryID, title)

		if category == nil {
			continue
		}

		if craw.Category != "" {
			if title == craw.Category {
				categories = append(categories, category...)
			}
		} else {
			categories = append(categories, category...)
		}

	}

	return categories, nil
}

func (craw *CategoryCrawlerInit) GetSubCategoryById(store *domain.Store, categoryID int, title string) ([]*domain.Category, error) {

	var dados interface{}

	var subCategory []*domain.Category
	url := "https://api.james.delivery/api/v2/customer/menu/" + strconv.Itoa(store.StoreID) + "?category_id=" + strconv.Itoa(categoryID) + "&offset=0&limit=100000"

	dados, err := utils.Request("GET", url, "", nil)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["aisles"].([]interface{}) {
		dado := item.(map[string]interface{})

		subTitle := utils.StringNotNull(dado["title"])

		obj, _ := domain.NewCategory(
			store.StoreID,
			categoryID,
			utils.IntNotNull(dado["id"]),
			title,
			subTitle,
			store.Name,
			store.Address,
			craw.List,
		)
		if obj != nil {
			subCategory = append(subCategory, obj)
		}

	}

	return subCategory, nil
}
