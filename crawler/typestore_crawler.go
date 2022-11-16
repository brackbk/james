package crawler

import (
	"fmt"
	"log"

	"gitlab.com/eiprice/crawlers/james/domain"
	"gitlab.com/eiprice/crawlers/james/repositories"
	"gitlab.com/eiprice/crawlers/james/utils"
)

//TypeStoreCrawler interface
type TypeStoreCrawler interface {
	GetData(segment string) ([]*domain.TypeStore, error)
}

//TypeStoreCrawlerInit struct
type TypeStoreCrawlerInit struct {
	TypeStoreRepository repositories.TypeStoreRepository
	List                string
	Lat                 string
	Lng                 string
}

//GetData return typestore
func (craw *TypeStoreCrawlerInit) GetData(segment string) ([]*domain.TypeStore, error) {

	var dados interface{}
	var headers []utils.Header
	var typestore []*domain.TypeStore

	url := `https://api.james.delivery/api/v2/customer/tabs?latitude=` + craw.Lat + `&longitude=` + craw.Lng
	log.Println(url)
	dados, err := utils.Request("GET", url, "", headers)

	if err != nil || dados == nil || dados.(map[string]interface{})["error"] != nil {
		return nil, fmt.Errorf("Error to find coordenates on james")
	}

	for _, item := range dados.(map[string]interface{})["tabs"].([]interface{}) {

		obj, _ := domain.NewTypeStore(
			utils.IntNotNull(item.(map[string]interface{})["id"]),
			item.(map[string]interface{})["name"].(string),
			craw.List,
		)
		if segment != "" {
			if item.(map[string]interface{})["name"].(string) == segment {
				typestore = append(typestore, obj)
			}
		} else {
			typestore = append(typestore, obj)
		}
	}

	craw.TypeStoreRepository.InsertMany(typestore)

	return typestore, nil

}
