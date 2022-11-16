package domain

import (
	"time"
)

type Product struct {
	Base
	Product_id        int     `json:"product_id"`
	Ean               string  `json:"ean" gorm:"type:varchar(255)" `
	In_stock          bool    `json:"In_stock`
	Corridor_id       int     `json:"corridor_id"`
	Corridor_name     string  `json:"corridor_name" gorm:"type:varchar(255)"`
	Sub_corridor_id   int     `json:"sub_corridor_id"`
	Sub_corridor_name string  `json:"sub_corridor_name" gorm:"type:varchar(255)"`
	Description       string  `json:"description" gorm:"type:text"`
	Image             string  `json:"image" gorm:"type:text"`
	Is_available      bool    `json:"is_available"`
	Name              string  `json:"name" gorm:"type:varchar(255)"`
	Price             float64 `json:"price"`
	Quantity          int     `json:"quantity"`
	Original_price    float64 `json:"original_price" gorm:"type:varchar(255)"`
	Store_id          int     `json:"store_id"`
	Store_name        string  `json:"store_name" gorm:"type:varchar(255)"`
	Store_Address     string  `json:"store_address" gorm:"type:varchar(255)"`
	TypeStore         string  `json:"typestore" gorm:"type:varchar(255)"`
	TypeStoreName     string  `json:"typestorename" gorm:"type:varchar(255)"`
	Scan              string  `json:"scan" gorm:"type:varchar(255)"`
	Url               string  `json:"url" gorm:"type:varchar(255)"`
}

func NewProduct(
	Product_id int,
	Ean string,
	In_stock bool,
	Corridor_id int,
	Corridor_name string,
	Sub_corridor_id int,
	Sub_corridor_name string,
	Description string,
	Image string,
	Is_available bool,
	Name string,
	Price float64,
	Quantity int,
	Original_price float64,
	Store_id int,
	Store_name string,
	Store_Address string,
	TypeStore string,
	TypeStoreName string,
	Scan string,
	Url string,
) (*Product, error) {

	product := &Product{
		Product_id:        Product_id,
		Ean:               Ean,
		In_stock:          In_stock,
		Corridor_id:       Corridor_id,
		Corridor_name:     Corridor_name,
		Sub_corridor_id:   Sub_corridor_id,
		Sub_corridor_name: Sub_corridor_name,
		Description:       Description,
		Image:             Image,
		Is_available:      Is_available,
		Name:              Name,
		Price:             Price,
		Quantity:          Quantity,
		Original_price:    Original_price,
		Store_id:          Store_id,
		Store_name:        Store_name,
		Store_Address:     Store_Address,
		TypeStore:         TypeStore,
		TypeStoreName:     TypeStoreName,
		Scan:              Scan,
		Url:               Url,
	}

	//product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	return product, nil
}
