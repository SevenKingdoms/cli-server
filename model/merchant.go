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
  Address      string `json:"address" form:"address" query:"address"`
  Images       string `json:"images" form:"images" query:"images"`
	Tel          string `json:"tel" form:"tel" query:"tel"`
  Password     string `json:"password" form:"password" query:"password"`
  OpenTime     string `json:"openTime" form:"openTime" query:"openTime"`
	Open         bool   `json:"open" form:"open" query:"open"`
}

func NewMerchant(id, hotIndex int64, name, introduction, logo, address, images, tel, password, openTime string, open bool) *Merchant {
	return &Merchant{
		ID:           id,
		Name:         name,
		HotIndex:     hotIndex,
		Introduction: introduction,
		Logo:         logo,
    Address:      address,
		Images:       images,
		Tel:          tel,
		Password:     password,
    Open:         open,
    OpenTime:     openTime,
		// CreatedAt:  time.Now().Unix(), // to string
	}
}

func (m *Merchant) Save(tx *dbr.Tx) error {
	var count = 0
	tempMerchant := new(Merchant)
	count, _ = tx.Select("*").From("Merchant").Where("id = ?", m.ID).Load(&tempMerchant)
	if count == 0 {
		_, err := tx.InsertInto("Merchant").
			Pair("name", m.Name).
			Pair("hotIndex", m.HotIndex).
			Pair("introduction", m.Introduction).
      Pair("logo", m.Logo).
			Pair("address", m.Address).
			Pair("images", m.Images).
			Pair("tel", m.Tel).
      Pair("password", m.Password).
      Pair("openTime", m.OpenTime).
			Pair("open", m.Open).
			Exec()
		return err
	} else {
		_, err := tx.Update("Merchant").
			Set("name", m.Name).
			Set("hotIndex", m.HotIndex).
			Set("introduction", m.Introduction).
			Set("logo", m.Logo).
			Set("address", m.Address).
			Set("images", m.Images).
			Set("tel", m.Tel).
      Set("password", m.Password).
      Set("openTime", m.OpenTime).
			Set("open", m.Open).
			Where("id = ?", m.ID).
			Exec()
		return err
	}

}

func (m *Merchant) Load(tx *dbr.Tx, tel string) (int, error) {
	// TODO: wnat is the int in (int, error)
	return tx.Select("*").
		From("Merchant").
		Where("tel = ?", tel).
		Load(m)
}

type Merchants []Merchant

func (m *Merchants) Load(tx *dbr.Tx) (int, error) {
	return tx.Select("*").
		From("Merchant").
		Load(m)
}
