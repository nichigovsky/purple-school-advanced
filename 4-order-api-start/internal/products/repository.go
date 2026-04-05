package products

import (
	"4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	res := repo.Database.DB.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func (repo *ProductRepository) GetById(id uint64) (*Product, error) {
	var product Product
	res := repo.Database.DB.First(&product, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}

func (repo *ProductRepository) GetAll() (*ProductResponse, error) {
	var product Product
	res := repo.Database.DB.Omit("deleted_at").Clauses(clause.Returning{}).Find(&product)

	if res.Error != nil {
		return nil, res.Error
	}

	return &ProductResponse{
		ID: product.ID,
		Name: product.Name,
		Description: product.Description,
		Images: product.Images,
	}, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)

	if res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) (int64, error) {
	res := repo.Database.DB.Delete(&Product{}, id)

	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}