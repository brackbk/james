package domain

import (
	"time"
)

type Store struct {
	Base
	Name          string  `json:"name" gorm:"type:varchar(255)"`
	TypeStore     string  `json:"typestore" gorm:"type:varchar(255)"`
	TypeStoreName string  `json:"typestorename" gorm:"type:varchar(255)"`
	ZoneID        int     `json:"zone_id"`
	StoreID       int     `json:"store_id"`
	Description   string  `json:"description" gorm:"type:varchar(255)"`
	Image         string  `json:"image" gorm:"type:varchar(255)"`
	Address       string  `json:"address" gorm:"type:varchar(255)"`
	CloseTime     string  `json:"close_time" gorm:"type:varchar(255)"`
	OpenTime      string  `json:"open_time" gorm:"type:varchar(255)"`
	TaxDelivery   float64 `json:"tax_delivery"`
	TimeDelivery  float64 `json:"time_delivery"`
	AvgPrice      float64 `json:"avg_price"`
	Rating        string  `json:"rating" gorm:"type:varchar(255)"`
	Lat           string  `json:"lat" gorm:"type:varchar(255)"`
	Lng           string  `json:"lng" gorm:"type:varchar(255)"`
	Scan          string  `json:"scan" gorm:"type:varchar(255)"`
}

func NewStore(
	Name string,
	TypeStore string,
	TypeStoreName string,
	ZoneID int,
	StoreID int,
	Description string,
	Image string,
	Address string,
	CloseTime string,
	OpenTime string,
	TaxDelivery float64,
	TimeDelivery float64,
	AvgPrice float64,
	Rating string,
	Lat string,
	Lng string,
	Scan string,
) (*Store, error) {

	store := &Store{
		StoreID:       StoreID,
		Name:          Name,
		TypeStore:     TypeStore,
		TypeStoreName: TypeStoreName,
		ZoneID:        ZoneID,
		Description:   Description,
		Image:         Image,
		Address:       Address,
		CloseTime:     CloseTime,
		OpenTime:      OpenTime,
		TaxDelivery:   TaxDelivery,
		TimeDelivery:  TimeDelivery,
		AvgPrice:      AvgPrice,
		Rating:        Rating,
		Lat:           Lat,
		Lng:           Lng,
		Scan:          Scan,
	}

	//store.ID = uuid.NewV4().String()
	store.CreatedAt = time.Now()

	return store, nil
}
