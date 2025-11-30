package database

import "errors"

type Item struct {
	ID   int
	Data any
}

type DatabaseInMemory struct {
	items map[int]Item
}

var dbInstance *DatabaseInMemory

func GetDatabaseInstance() *DatabaseInMemory {

	if dbInstance != nil {
		return dbInstance
	}

	dbInstance = &DatabaseInMemory{
		items: make(map[int]Item),
	}

	return dbInstance
}

func (db *DatabaseInMemory) CreateItem(data any) (Item, error) {
	id := len(db.items) + 1
	db.items[id] = Item{ID: id, Data: data}

	return db.items[id], nil

}

func (db *DatabaseInMemory) GetItem(id int) (Item, error) {

	item, ok := db.items[id]
	if !ok {
		return Item{}, errors.New("item not found")
	}

	return item, nil
}

func (db *DatabaseInMemory) UpdateItem(id int, data any) (Item, error) {
	db.items[id] = Item{
		ID:   id,
		Data: data,
	}
	return db.items[id], nil
}

func (db *DatabaseInMemory) DeleteItem(id int) error {
	delete(db.items, id)
	return nil
}

func (db *DatabaseInMemory) GetAllItems() map[int]Item {
	return db.items
}
