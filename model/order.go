package model

import (
	"time"
  "fmt"

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
	User_openId   string    `json:"open_id" form:"open_id" query:"open_id"`
	Merchant_id   int64     `json:"merchant_id" form:"merchant_id" query:"merchant_id"`
	Foods         string    `json:"foods" form:"foods" query:"foods"`
  // datetime := time.Now().Format(time.RFC3339)
  // https://stackoverflow.com/questions/23415612/insert-datetime-using-now-with-go
	Created_at    string    `json:"created_at" form:"created_at" query:"created_at"`
	Merchant_name string    `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	Merchant_tel  string    `json:"merchant_tel" form:"merchant_tel" query:"merchant_tel"`
}

func NewOrder(id, num_of_people, deskId int64, remark string, paid bool, open_id string, merchant_id int64, foods, merchant_name, merchant_tel string) *Order {
	return &Order{
		Id:            id,
		NumOfPeople:   num_of_people,
		DeskId:        deskId,
		Remark:        remark,
		Paid:          paid,
		User_openId:   open_id,
		Merchant_id:   merchant_id,
		Foods:         foods,
		Merchant_name: merchant_name,
		Merchant_tel:  merchant_tel,
		Created_at:    time.Now().Format(time.RFC3339),
	}
}

func (u *Order) Save(tx *dbr.Tx) error {

	var count = 0
	tempOrder := new(Order)
	count, err := tx.Select("*").
		From("KOrder").
		Where("id = ?", u.Id).
		Load(&tempOrder)
  fmt.Println(u)
	if count == 0 {
		//fmt.Println("here error")
		// if user not exists, Create
		_, err = tx.InsertInto("KOrder").
			Pair("num_of_people", u.NumOfPeople).
			Pair("deskid", u.DeskId).
			Pair("remark", u.Remark).
			Pair("paid", u.Paid).
			Pair("user_openid", u.User_openId).
			Pair("merchant_id", u.Merchant_id).
			Pair("foods", u.Foods).
			Pair("create_at", u.Created_at).
			Pair("merchant_name", u.Merchant_name).
			Pair("merchant_tel", u.Merchant_tel).
			Exec()
		return err
	} else {
		// if order exists, Update
		_, err = tx.Update("KOrder").
			Set("num_of_people", u.NumOfPeople).
			Set("deskid", u.DeskId).
			Set("remark", u.Remark).
			Set("paid", u.Paid).
			Set("user_openid", u.User_openId).
			Set("merchant_id", u.Merchant_id).
			Set("foods", u.Foods).
      // Set("create_at", u.Created_at.Format("2006-01-02 15:04:05")).
      Set("create_at", u.Created_at).
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
		From("KOrder").
		Where("user_openid = ?", open_id).
		Load(u)
}

func (u *Order) Load(tx *dbr.Tx, id int64) (int, error) {
	return tx.Select("*").
		From("KOrder").
		Where("id = ?", id).
		Load(u)
}

func (u *Orders) MerchantIdLoad(tx *dbr.Tx, merchant_id int64) (int, error) {
	return tx.Select("*").
		From("KOrder").
		Where("merchant_id = ?", merchant_id).
		Load(u)
}

func (f *Order) OrderDelete(tx *dbr.Tx, order_id int64) error {
	_, err := tx.DeleteFrom("Order").Where("id = ?", order_id).Exec()
	return err
}
