package repository

import (
	"jchhay/go-rest-api-gin/pkg/db"
)

// Find
func Find(where interface{}, out interface{}, orders ...string) error {
	db := db.GetDB()

	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Save
func Save(value interface{}) error {
	return db.GetDB().Save(value).Error
}

// Delete
func DeleteByID(model interface{}, id int) (count int64, err error) {
	db := db.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return 0, err
	}
	count = db.RowsAffected
	return count, nil
}

// First
func FirstByID(out interface{}, id int) (err error) {
	err = db.GetDB().First(out, id).Error
	if err != nil {
		return err
	}
	return nil
}
