package domain

import (
	"time"
)

type Category struct {
	Base
	StoreID       int    `json:"store_id"`
	CategoryID    int    `json:"category_id"`
	SubCategoryID int    `json:"sub_category_id"`
	Name          string `json:"name" gorm:"type:varchar(255)"`
	SubName       string `json:"sub_name" gorm:"type:varchar(255)"`
	StoreName     string `gorm:"-"`
	StoreAddress  string `json:"store_address" gorm:"type:varchar(255)"`
	Scan          string `json:"scan" gorm:"type:varchar(255)"`
}

func NewCategory(
	StoreID int,
	CategoryID int,
	SubCategoryID int,
	Name string,
	SubName string,
	StoreName string,
	StoreAddress string,
	Scan string,
) (*Category, error) {

	category := &Category{
		StoreID:       StoreID,
		CategoryID:    CategoryID,
		SubCategoryID: SubCategoryID,
		Name:          Name,
		SubName:       SubName,
		StoreName:     StoreName,
		StoreAddress:  StoreAddress,
		Scan:          Scan,
	}

	//category.ID = uuid.NewV4().String()
	category.CreatedAt = time.Now()

	return category, nil
}
