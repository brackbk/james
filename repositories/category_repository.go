package repositories

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/crawlers/james/domain"
)

type CategoryRepository interface {
	Insert(category *domain.Category) (*domain.Category, error)
	InsertMany(category []*domain.Category) ([]*domain.Category, error)
}

type CategoryRepositoryDb struct {
	Db *gorm.DB
}

func (repo CategoryRepositoryDb) Insert(category *domain.Category) (*domain.Category, error) {

	err := repo.Db.Create(category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (repo CategoryRepositoryDb) InsertMany(category []*domain.Category) ([]*domain.Category, error) {

	//Start Transacion
	tx := repo.Db.Begin()
	for _, ts := range category {
		if err := tx.Create(ts).Error; err != nil {
			//Rollback if somthing went rong
			tx.Rollback()
			return nil, err
		}
	}
	//Commit Transaction
	tx.Commit()
	return category, nil
}
