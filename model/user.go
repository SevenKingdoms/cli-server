package model

import (
	// "time"

	"github.com/gocraft/dbr"
)

type User struct {
	OpenId    string `json:"open_id" form:"open_id" query:"open_id"`
	Name      string `json:"nick_name" form:"nick_name" query:"nick_name"`
	Avatar    string `json:"avatar" form:"avatar" query:"avatar"`
	Phone     string `json:"tel" form:"tel" query:"tel"`
	// CreatedAt string  `json:"createdAt"`
}

func NewUser(open_id, nick_name, avatar, tel string) *User {
	return &User{
		OpenId:     open_id,
		Name:       nick_name,
		Avatar:     avatar,
  	Phone:      tel,
		// CreatedAt:  time.Now().Unix(), // to string
	}
}

func (u *User) Save(tx *dbr.Tx) error {

  // TODO: if user exists, Update

	_, err := tx.InsertInto("User").
    Pair("openId", u.OpenId).
    Pair("name", u.Name).
    Pair("avatar", u.Avatar).
    Pair("phone", u.Phone).
		Exec()
    // Error 1064: You have an error in your SQL syntax; check the manual that corresponds
    // to your MySQL server version for the right syntax to use near '' at line 1
    // ----
    // Columns("openId", "name", "avatar", "phone").
    // Record(&u).
    // Exec()

	return err
}

func (u *User) Load(tx *dbr.Tx, open_id string) (int, error) {
  // TODO: wnat is the int in (int, error)
	return tx.Select("*").
		From("User").
		Where("openId = ?", open_id).
		Load(u)
}

type Users []User

func (u *Users) Load(tx *dbr.Tx) (int, error) {
	return tx.Select("*").
		From("User").
		Load(u)
}
