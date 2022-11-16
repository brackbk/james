package repositories

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/crawlers/james/domain"
)

type StoreRepository interface {
	Insert(store *domain.Store) (*domain.Store, error)
	InsertMany(store []*domain.Store) ([]*domain.Store, error)
}

type StoreRepositoryDb struct {
	Db *gorm.DB
}

func (repo StoreRepositoryDb) Insert(store *domain.Store) (*domain.Store, error) {

	err := repo.Db.Create(store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (repo StoreRepositoryDb) InsertMany(store []*domain.Store) ([]*domain.Store, error) {

	//Start Transacion
	tx := repo.Db.Begin()
	for _, ts := range store {
		if err := tx.Create(ts).Error; err != nil {
			//Rollback if somthing went rong
			tx.Rollback()
			return nil, err
		}
	}
	//Commit Transaction
	tx.Commit()
	return store, nil
}
