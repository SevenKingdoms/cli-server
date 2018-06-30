package model

import (
	// "time"
	"fmt"

	"github.com/gocraft/dbr"
)

type Comment struct {
	Id          int64   `json:"id" form:"id" query:"id"`
	Description string  `json:"description" form:"description" query:"description"`
	Images      string  `json:"images" form:"images" query:"images"`
	User_openId string  `json:"user_openId" form:"user_openId" query:"user_openId"`
	Merchant_id int64   `json:"merchant_id" form:"merchant_id" query:"merchant_id"`
	Score       float64 `json:"score" form:"score" query:"score"`
}

func NewComment(id int64, description, images, user_openId string, merchant_id int64, score float64) *Comment {
	return &Comment{
		Id:          id,
		Description: description,
		Images:      images,
		User_openId: user_openId,
		Merchant_id: merchant_id,
		Score:       score,
	}
}

func (m *Comment) Save(tx *dbr.Tx) error {
	var count = 0
	tempMerchant := new(Comment)
	fmt.Println(m)
	count, _ = tx.Select("*").From("Comment").Where("id = ?", m.Id).Load(&tempMerchant)

	fmt.Println(count)
	if count == 0 {
		_, err := tx.InsertInto("Comment").
			Pair("id", m.Id).
			Pair("description", m.Description).
			Pair("images", m.Images).
			Pair("user_openId", m.User_openId).
			Pair("Merchant_id", m.Merchant_id).
			Pair("score", m.Score).
			Exec()
		return err
	} else {
		_, err := tx.Update("Comment").
			Set("description", m.Description).
			Set("images", m.Images).
			Set("user_openid", m.User_openId).
			Set("merchant_id", m.Merchant_id).
			Set("score", m.Score).
			Where("id = ?", m.Id).
			Exec()
		return err
	}

}

func (m *Comment) Load(tx *dbr.Tx, id int64) (int, error) {
	return tx.Select("*").
		From("Comment").
		Where("id = ?", id).
		Load(m)
}

type Comments []Comment

func (m *Comments) MerchantLoad(tx *dbr.Tx, id int64) (int, error) {
	return tx.Select("*").
		From("Comment").
		Where("Merchant_id= ?", id).
		Load(m)
}

func (m *Comments) UserIdLoad(tx *dbr.Tx, id string) (int, error) {
	return tx.Select("*").
		From("Comment").
		Where("User_openId= ?", id).
		Load(m)
}

func (f *Comment) CommentDelete(tx *dbr.Tx, comment_id int64) error {
	_, err := tx.DeleteFrom("Comment").Where("id = ?", comment_id).Exec()
	return err
}
