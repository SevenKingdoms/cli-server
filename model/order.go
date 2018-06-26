package model

import (
	"time"

	"github.com/gocraft/dbr"
)

// "fmt"

type Order struct {
	Id          int64  `json:"order_id" form:"order_id" query:"order_id"`
	Date        int64  `json:"date" form:"date" query:"date"`
	NumOfPeople int64  `json:"numOfPeople" form:"numOfPeople" query:"numOfPeople"`
	DeskId      int64  `json:"deskId" form:"deskId" query:"deskId"`
	Remark      string `json:"remark" form:"remark" query:"remark"`
	User_OpenId string `json:"user_OpenId" form:"user_OpenId" query:"user_OpenId"`
	Merchant_id int64  `json:"marchant_id" form:"marchant_id" query:"marchant_id"`
	Status      int64  `json:"status" form:"status" query:"status"`
}

func NewOrder(id, numOfPeople, deskId int64, remark, user_OpenId string, merchant_id int64, status int64) *Order {
	return &Order{
		Id:          id,
		Date:        time.Now().Unix(),
		NumOfPeople: numOfPeople,
		DeskId:      deskId,
		Remark:      remark,
		User_OpenId: user_OpenId,
		Merchant_id: merchant_id,
		Status:      status,
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
		// if user not exists, Create
		_, err = tx.InsertInto("Order").
			Pair("id", u.Id).
			Pair("date", u.Date).
			Pair("numOfPeople", u.NumOfPeople).
			Pair("deskId", u.DeskId).
			Pair("remark", u.Remark).
			Pair("User_openId", u.User_OpenId).
			Pair("Merchant_id", u.Merchant_id).
			Pair("status", u.Status).
			Exec()
	} else {
		// if order exists, Update
		_, err = tx.Update("Order").
			Set("date", u.Date).
			Set("numOfPeople", u.NumOfPeople).
			Set("deskId", u.DeskId).
			Set("remark", u.Remark).
			Set("User_openId", u.User_OpenId).
			Set("Merchant_id", u.Merchant_id).
			Set("status", u.Status).
			Where("id = ?", u.Id).
			Exec()
	}

	return err
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
		Where("Id = ?", id).
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
