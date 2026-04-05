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

func (repo *ProductRepository) GetAll() ([]ProductResponse, error) {
	var products []Product

	res := repo.Database.DB.
		Omit("deleted_at").
		Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	var response []ProductResponse

	for _, p := range products {
		response = append(response, ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Images:      p.Images,
		})
	}

	return response, nil
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