package model

import (
	// "time"
	"github.com/gocraft/dbr"
)

type Merchant struct {
	ID           int64  `json:"id" form:"id" query:"id"`
	Name         string `json:"name" form:"name" query:"name"`
	HotIndex     int64  `json:"hotIndex" form:"hotIndex" query:"hotIndex"`
	Introduction string `json:"introduction" form:"introduction" query:"introduction"`
	Logo         string `json:"logo" form:"logo" query:"logo"`
	Images       string `json:"iamges" form:"iamges" query:"images"`
	Account      string `json:"account" form:"account" query:"account"`
	Password     string `json:"password" form:"password" query:"password"`
}

func NewMerchant(id int64, name string, hotIndex int64, introduction, logo, images, account, password string) *Merchant {
	return &Merchant{
		ID:           id,
		Name:         name,
		HotIndex:     hotIndex,
		Introduction: introduction,
		Logo:         logo,
		Images:       images,
		Account:      account,
		Password:     password,
		// CreatedAt:  time.Now().Unix(), // to string
	}
}

func (m *Merchant) Save(tx *dbr.Tx) error {
	var count = 0
	tempMerchant := new(Merchant)
	count, _ = tx.Select("*").From("Merchant").Where("id = ?", m.ID).Load(&tempMerchant)
	if count == 0 {
		_, err := tx.InsertInto("Merchant").
			Pair("id", m.ID).
			Pair("name", m.Name).
			Pair("hotIndex", m.HotIndex).
			Pair("introduction", m.Introduction).
			Pair("logo", m.Logo).
			Pair("images", m.Images).
			Pair("account", m.Account).
			Pair("password", m.Password).
			Exec()
		return err
	} else {
		_, err := tx.Update("Merchant").
			Set("name", m.Name).
			Set("hotIndex", m.HotIndex).
			Set("introduction", m.Introduction).
			Set("logo", m.Logo).
			Set("images", m.Images).
			Set("account", m.Account).
			Set("password", m.Password).
			Where("id = ?", m.ID).
			Exec()
		return err
	}

}

func (m *Merchant) Load(tx *dbr.Tx, id int64) (int, error) {
	// TODO: wnat is the int in (int, error)
	return tx.Select("*").
		From("Merchant").
		Where("id = ?", id).
		Load(m)
}

type Merchants []Merchant

func (m *Merchants) Load(tx *dbr.Tx) (int, error) {
	return tx.Select("*").
		From("Merchant").
		Load(m)
}
