package domain

import (
	"time"
)

//TypeStore struct
type TypeStore struct {
	Base
	TypeStoreID int    `json:"typestore_id"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Scan        string `json:"list" gorm:"type:varchar(255)"`
}

//NewTypeStore return typestore
func NewTypeStore(typeStoreID int, name string, scan string) (*TypeStore, error) {

	typestore := &TypeStore{
		TypeStoreID: typeStoreID,
		Name:        name,
		Scan:        scan,
	}

	//typestore.ID = uuid.NewV4().String()
	typestore.CreatedAt = time.Now()

	return typestore, nil
}
