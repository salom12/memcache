package database

type row struct {
	ID uint64
}

type DbHandler struct {
}

func (db DbHandler) GetByID(id uint64) (record *row) {
	return &row{
		ID: 1,
	}
}

func (db DbHandler) Insert(updatedRow any) (id uint64) {
	return 1
}

func (db DbHandler) Update(id uint64, updatedRow any) (success bool) {
	return true
}

func (db DbHandler) Delete(id uint64) (success bool) {
	return true
}

var db *DbHandler

func InitDB() error {
	db = &DbHandler{}
	return nil
}

func GetDB() *DbHandler {
	return db
}
