package repositories

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/crawlers/james/domain"
)

type ProductRepository interface {
	Insert(product *domain.Product) (*domain.Product, error)
	InsertMany(product []*domain.Product) ([]*domain.Product, error)
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (repo ProductRepositoryDb) Insert(product *domain.Product) (*domain.Product, error) {

	err := repo.Db.Create(product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (repo ProductRepositoryDb) InsertMany(product []*domain.Product) ([]*domain.Product, error) {

	//Start Transacion
	tx := repo.Db.Begin()
	for _, ts := range product {
		if err := tx.Create(ts).Error; err != nil {
			//Rollback if somthing went rong
			tx.Rollback()
			return nil, err
		}
	}
	//Commit Transaction
	tx.Commit()
	return product, nil
}
