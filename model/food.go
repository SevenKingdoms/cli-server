package model

import (
	// "time"
	// "fmt"

	"github.com/gocraft/dbr"
)

type Food struct {
	Id           int64   `json:"food_id" form:"food_id" query:"food_id"`
	Name         string  `json:"name" form:"name" query:"name"`
	Image        string  `json:"image" form:"image" query:"image"`
	Type         string  `json:"type" form:"type" query:"type"`
	Price        float64 `json:"price" form:"price" query:"price"`
	HotIndex     int64   `json:"hot_index" form:"hot_index" query:"hot_index"`
	Introduction string  `json:"introduction" form:"introduction" query:"introduction"`
	Merchant_id  int64   `json:"merchant_id" form:"merchant_id" query:"merchant_id"`
	InStock      bool    `json:"in_stock" form:"in_stock" query:"in_stock"`
}

func NewFood(id int64, name, image, food_type string, price float64, hot_index int64, intro string, merchant_id int64, in_stock bool) *Food {
	return &Food{
		Id:           id,
		Name:         name,
		Image:        image,
		Type:         food_type,
		Price:        price,
		HotIndex:     hot_index,
		Introduction: intro,
		Merchant_id:  merchant_id,
		InStock:      in_stock,
	}
}

func (f *Food) Save(tx *dbr.Tx) error {

	var count = 0
	tempFood := new(Food)
	count, err := tx.Select("*").
		From("Food").
		Where("id = ?", f.Id).
		Load(&tempFood)

	if count == 0 {
		// if food not exists, Create
		_, err = tx.InsertInto("Food").
			Pair("id", f.Id).
			Pair("name", f.Name).
			Pair("image", f.Image).
			Pair("type", f.Type).
			Pair("price", f.Price).
			Pair("hot_index", f.HotIndex).
			Pair("introduction", f.Introduction).
			Pair("merchant_id", f.Merchant_id).
			Pair("in_stock", f.InStock).
			Exec()
	} else {
		// if food exists, Update
		_, err = tx.Update("Food").
			Set("name", f.Name).
			Set("image", f.Image).
			Set("type", f.Type).
			Set("price", f.Price).
			Set("hot_index", f.HotIndex).
			Set("introduction", f.Introduction).
			Set("merchant_id", f.Merchant_id).
			Set("in_stock", f.InStock).
			Where("id = ?", f.Id).
			Exec()
	}

	return err
}

func (f *Food) Load(tx *dbr.Tx, id int64) (int, error) {
	return tx.Select("*").
		From("Food").
		Where("id = ?", id).
		Load(f)
}

type Foods []Food

func (f *Foods) MerchantLoad(tx *dbr.Tx, merchant_id int64) (int, error) {
	return tx.Select("*").
		From("Food").
		Where("Merchant_id = ?", merchant_id).
		Load(f)

}

func (f *Food) FoodDelete(tx *dbr.Tx, food_id int64) error {
	_, err := tx.DeleteFrom("Food").Where("id = ?", food_id).Exec()
	return err
}

// func (u *Foods) Load(tx *dbr.Tx) (int, error) {
// 	return tx.Select("*").
// 		From("Food").
// 		Load(u)
// }
