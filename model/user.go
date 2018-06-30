package model

import (
	// "time"
  // "fmt"

	"github.com/gocraft/dbr"
)

type User struct {
	OpenId    string `json:"open_id" form:"open_id" query:"open_id"`
	Name      string `json:"nick_name" form:"nick_name" query:"nick_name"`
	Avatar    string `json:"avatar" form:"avatar" query:"avatar"`
	Phone     string `json:"tel" form:"tel" query:"tel"`
}

func NewUser(open_id, nick_name, avatar, tel string) *User {
	return &User{
		OpenId:     open_id,
		Name:       nick_name,
		Avatar:     avatar,
  	Phone:      tel,
	}
}

func (u *User) Save(tx *dbr.Tx) error {

  var count = 0
  tempUser := new(User)
  count, err:= tx.Select("*").
		From("User").
		Where("openid = ?", u.OpenId).
		Load(&tempUser)

  if count == 0 {
    // if user not exists, Create
    _, err = tx.InsertInto("User").
      Pair("openid", u.OpenId).
      Pair("name", u.Name).
      Pair("avatar", u.Avatar).
      Pair("phone", u.Phone).
  		Exec()
  } else {
    // if user exists, Update
    _, err = tx.Update("User").
      Set("name", u.Name).
      Set("avatar", u.Avatar).
      Set("phone", u.Phone).
      Where("openid = ?", u.OpenId).
      Exec()
  }

	return err
}

func (u *User) Load(tx *dbr.Tx, open_id string) (int, error) {
	return tx.Select("*").
		From("User").
		Where("openid = ?", open_id).
		Load(u)
}

type Users []User

func (u *Users) Load(tx *dbr.Tx) (int, error) {
	return tx.Select("*").
		From("User").
		Load(u)
}
