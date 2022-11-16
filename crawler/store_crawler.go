package crawler

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/eiprice/crawlers/james/domain"
	"gitlab.com/eiprice/crawlers/james/repositories"
	"gitlab.com/eiprice/crawlers/james/utils"
)

type StoreCrawler interface {
	GetData(typeStore string, StoreName string) ([]*domain.Store, error)
	GetDataByPage(typeStore *domain.TypeStore, StoreName string, page int, limit int) ([]*domain.Store, error)
}

type StoreCrawlerInit struct {
	StoreRepository repositories.StoreRepository
	Lat             string
	Lng             string
	List            string
}

func (craw *StoreCrawlerInit) GetData(typeStore *domain.TypeStore, StoreName string) ([]*domain.Store, error) {

	var stores []*domain.Store
	var listStores []*domain.Store
	var dados interface{}

	var page int = 1
	var limit int = 20

	url := `https://api.james.delivery/api/v2/customer/restaurants?latitude=` +
		craw.Lat +
		`&longitude=` +
		craw.Lng +
		`&tab_id=` +
		strconv.Itoa(typeStore.TypeStoreID) +
		`&page=` + strconv.Itoa(page) +
		`&limit=` + strconv.Itoa(limit)
	dados, err := utils.Request("GET", url, "", nil)

	if err != nil || dados == nil {
		return nil, err
	}

	limit = 20
	page = 1
	totalStores := utils.IntNotNull(dados.(map[string]interface{})["meta"].(map[string]interface{})["total"])

	if totalStores > limit {
		totalStores = (totalStores / limit) + 1
	} else {
		totalStores = 1
	}

	for page <= totalStores {
		stores, _ = craw.GetDataByPage(typeStore, StoreName, page, limit)
		if stores != nil {
			craw.StoreRepository.InsertMany(stores)
			listStores = append(listStores, stores...)
		}
		page = page + 1
	}

	return listStores, nil

}

func (craw *StoreCrawlerInit) GetDataByPage(typeStore *domain.TypeStore, StoreName string, page int, limit int) ([]*domain.Store, error) {

	var stores []*domain.Store
	var store *domain.Store
	var dados interface{}

	url := `https://api.james.delivery/api/v2/customer/restaurants?latitude=` +
		craw.Lat +
		`&longitude=` +
		craw.Lng +
		`&tab_id=` +
		strconv.Itoa(typeStore.TypeStoreID) +
		`&page=` + strconv.Itoa(page) +
		`&limit=` + strconv.Itoa(limit)
	dados, err := utils.Request("GET", url, "", nil)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["restaurants"].([]interface{}) {
		dado := item.(map[string]interface{})
		fmt.Println("get store : ", utils.StringNotNull(dado["name"]))

		store, _ = domain.NewStore(
			utils.StringNotNull(dado["name"]),
			typeStore.Name,
			typeStore.Name,
			utils.IntNotNull(dado["gpa_id"]),
			utils.IntNotNull(dado["id"]),
			utils.StringNotNull(dado["name"]),
			" ",
			utils.StringNotNull(dado["address"]),
			" ",
			" ",
			utils.FloatNotNull(dado["original_delivery_fee"]),
			utils.FloatNotNull(dado["delivery_time"]),
			0,
			" ",
			fmt.Sprint(utils.FloatNotNull(dado["latitude"])),
			fmt.Sprint(utils.FloatNotNull(dado["longitude"])),
			craw.List,
		)

		if StoreName != "" {
			if strings.Contains(dado["name"].(string), StoreName) {
				stores = append(stores, store)
			}
		} else {
			stores = append(stores, store)
		}
	}

	return stores, nil

}
