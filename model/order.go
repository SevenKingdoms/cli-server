package model

import (
	"time"

	"github.com/gocraft/dbr"
)

// "fmt"
// type _Food struct {
// 	name   string
// 	price  float64
// 	number int64
// }
// type _Foods []_Food

type Order struct {
	Id            int64     `json:"order_id" form:"order_id" query:"order_id"`
	NumOfPeople   int64     `json:"num_of_people" form:"num_of_people" query:"num_of_people"`
	DeskId        int64     `json:"desk_id" form:"desk_id" query:"desk_id"`
	Remark        string    `json:"remark" form:"remark" query:"remark"`
	Paid          bool      `json:"paid" form:"paid" query:"paid"`
	User_openId   string    `json:"user_openId" form:"user_openId" query:"user_openId"`
	Merchant_id   int64     `json:"merchant_id" form:"merchant_id" query:"merchant_id"`
	Foods         []string  `json:"foods" form:"foods" query:"foods"`
  // TODO: string datetime := time.Now().Format(time.RFC3339)
  // https://stackoverflow.com/questions/23415612/insert-datetime-using-now-with-go
	Created_at    time.Time `json:"created_at" form:"created_at" query:"created_at"`
	Merchant_name string    `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	Merchant_tel  string    `json:"merchant_tel" form:"merchant_tel" query:"merchant_tel"`
}

func NewOrder(id, numOfPeople, deskId int64, remark string, paid bool, user_OpenId string, merchant_id int64, foods []string, merchant_name, merchant_tel string, time_temp time.Time) *Order {
	return &Order{
		Id:            id,
		NumOfPeople:   numOfPeople,
		DeskId:        deskId,
		Remark:        remark,
		Paid:          paid,
		User_openId:   user_OpenId,
		Merchant_id:   merchant_id,
		Foods:         foods,
		Created_at:    time_temp,
		Merchant_name: merchant_name,
		Merchant_tel:  merchant_tel,
	}
}

func (u *Order) Save(tx *dbr.Tx) error {

	var count = 0
	tempOrder := new(Order)
	count, err := tx.Select("*").
		From("Order").
		Where("id = ?", u.Id).
		Load(&tempOrder)

	if count == 0 {
		//fmt.Println("here error")
		// if user not exists, Create
		_, err = tx.InsertInto("Order").
			Pair("id", u.Id).
			Pair("numOfPeople", u.NumOfPeople).
			Pair("deskId", u.DeskId).
			Pair("remark", u.Remark).
			Pair("paid", u.Paid).
			Pair("user_openId", u.User_openId).
			Pair("Merchant_id", u.Merchant_id).
			Pair("foods", u.Foods).
			Pair("create_at", u.Created_at.Format("2006-01-02 15:04:05")).
			Pair("merchant_name", u.Merchant_name).
			Pair("merchant_tel", u.Merchant_tel).
			Exec()
		return err
	} else {
		// if order exists, Update
		_, err = tx.Update("Order").
			Set("numOfPeople", u.NumOfPeople).
			Set("deskId", u.DeskId).
			Set("remark", u.Remark).
			Set("paid", u.Paid).
			Set("user_openId", u.User_openId).
			Set("Merchant_id", u.Merchant_id).
			Set("foods", u.Foods).
			Set("create_at", u.Created_at.Format("2006-01-02 15:04:05")).
			Set("merchant_name", u.Merchant_name).
			Set("merchant_tel", u.Merchant_tel).
			Where("id = ?", u.Id).
			Exec()
		return err
	}

}

type Orders []Order

func (u *Orders) OpenIdLoad(tx *dbr.Tx, open_id string) (int, error) {
	return tx.Select("*").
		From("Order").
		Where("User_openId = ?", open_id).
		Load(u)
}

func (u *Order) Load(tx *dbr.Tx, id int64) (int, error) {
	return tx.Select("*").
		From("Order").
		Where("id = ?", id).
		Load(u)
}

func (u *Orders) MerchantIdLoad(tx *dbr.Tx, merchant_id int64) (int, error) {
	return tx.Select("*").
		From("Order").
		Where("Merchant_id = ?", merchant_id).
		Load(u)
}

func (f *Order) OrderDelete(tx *dbr.Tx, order_id int64) error {
	_, err := tx.DeleteFrom("Order").Where("id = ?", order_id).Exec()
	return err
}
