package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/crawlers/james/domain"
)

type TypeStoreRepository interface {
	Insert(typestore *domain.TypeStore) (*domain.TypeStore, error)
	InsertMany(typestore []*domain.TypeStore) ([]*domain.TypeStore, error)
	FindAll() ([]*domain.TypeStore, error)
	Find(id int) (*domain.TypeStore, error)
}

type TypestoreRepositoryDb struct {
	Db *gorm.DB
}

func (repo TypestoreRepositoryDb) Insert(typestore *domain.TypeStore) (*domain.TypeStore, error) {

	err := repo.Db.Create(typestore).Error
	if err != nil {
		return typestore, err
	}

	return typestore, nil
}

func (repo TypestoreRepositoryDb) InsertMany(typestore []*domain.TypeStore) ([]*domain.TypeStore, error) {

	//Start Transacion
	tx := repo.Db.Begin()
	for _, ts := range typestore {
		err := tx.Create(ts).Error
		if err != nil {
			//Rollback if somthing went rong
			tx.Rollback()
			return nil, err
		}
	}
	//Commit Transaction
	tx.Commit()
	return typestore, nil
}

func (repo TypestoreRepositoryDb) FindAll() ([]*domain.TypeStore, error) {

	var typestore []*domain.TypeStore
	repo.Db.Find(&typestore)
	if len(typestore) == 0 {
		return nil, fmt.Errorf("typestore is empty")
	}

	return typestore, nil
}

func (repo TypestoreRepositoryDb) Find(id int) (*domain.TypeStore, error) {

	var typestore *domain.TypeStore
	repo.Db.First(&typestore, "id = ?", id)

	if typestore.ID == 0 {
		return nil, fmt.Errorf("typestore does not exist")
	}

	return typestore, nil
}
