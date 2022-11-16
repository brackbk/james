package crawler

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/crawlers/james/repositories"
)

type Crawler struct {
	Db        *gorm.DB
	Segment   string
	StoreName string
	Category  string
	Scan      string
	Lat       string
	Lng       string
}

func (craw *Crawler) Start() error {

	typeStoreRepo := repositories.TypestoreRepositoryDb{craw.Db}

	typeStoreCraw := TypeStoreCrawlerInit{typeStoreRepo, craw.Scan, craw.Lat, craw.Lng}

	fmt.Println("request TypeStores from james")
	typestore, err := typeStoreCraw.GetData(craw.Segment)

	if err != nil {
		log.Println("Coordenates not have on james:", err)
		return err
	}

	storeRepo := repositories.StoreRepositoryDb{craw.Db}
	storeCraw := StoreCrawlerInit{storeRepo, craw.Lat, craw.Lng, craw.Scan}

	categoryRepo := repositories.CategoryRepositoryDb{craw.Db}
	categoryCraw := CategoryCrawlerInit{categoryRepo, craw.Category, craw.Scan}

	productRepo := repositories.ProductRepositoryDb{craw.Db}
	productCraw := ProductCrawlerInit{productRepo, craw.Scan}

	for _, item := range typestore {
		fmt.Println("Get Stores from: ", item.Name)
		stores, err := storeCraw.GetData(item, craw.StoreName)

		if err != nil {
			log.Println("Error Get Store Data ", err)
			continue
		}

		storeCategories, _ := categoryCraw.GetData(stores)
		productCraw.GetData(storeCategories, item)
	}

	return nil

}
