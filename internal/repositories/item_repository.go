package repositories

import (
	"github.com/salom12/memcache/internal/database"
	"github.com/salom12/memcache/internal/models"
)

type ItemRepository struct {
	db *database.DbHandler
}

func NewItemRepository(db *database.DbHandler) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Get() ([]models.Item, error) {
	rows := []models.Item{
		{ID: 1, Name: "Item 1", Price: 100, Quantity: 1},
		{ID: 2, Name: "Item 2", Price: 200, Quantity: 1},
		{ID: 3, Name: "Item 3", Price: 300, Quantity: 1},
	}
	return rows, nil
}

func (r *ItemRepository) GetByID(id uint64) (*models.Item, error) {
	row := r.db.GetByID(id)
	return &models.Item{ID: row.ID, Name: "test", Price: 100, Quantity: 1}, nil
}

func (r *ItemRepository) Insert(item *models.Item) (uint64, error) {
	i := &models.Item{ID: 2, Name: "test", Price: 100, Quantity: 1}
	id := r.db.Insert(i)
	return id, nil
}

func (r *ItemRepository) Update(id uint64, item *models.Item) bool {
	i := &models.Item{ID: 2, Name: "test", Price: 100, Quantity: 1}
	return r.db.Update(2, i)
}

func (m *ItemRepository) Delete(id uint64) bool {
	return m.db.Delete(1)
}
