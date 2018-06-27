package api

import (
	// "strconv"
	// "fmt"
	"github.com/Sirupsen/logrus"
	"github.com/cli-server/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := new(model.User)
		if err = c.Bind(m); err != nil {
			logrus.Debug(err)
  		return c.JSON(fasthttp.StatusInternalServerError,
  			NewJSON("OK", "内部错误", nil))
		}

		tx := c.Get("Tx").(*dbr.Tx)

		user := model.NewUser(m.OpenId, m.Name, m.Avatar, m.Phone)

		if err := user.Save(tx); err != nil {
			logrus.Debug(err)
  		return c.JSON(fasthttp.StatusBadRequest,
  			NewJSON("OK", "创建/更新用户失败", nil))
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建/更新用户", user))
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		openId := c.Param("open_id")

		tx := c.Get("Tx").(*dbr.Tx)

		user := new(model.User)
    count, err := user.Load(tx, openId)
    if err != nil {
			logrus.Debug(err)
  		return c.JSON(fasthttp.StatusInternalServerError,
  			NewJSON("OK", "内部错误", nil))
		}
    if count == 0 {
     return c.JSON(fasthttp.StatusOK,
       NewJSON("OK", "用户不存在", nil))
    }
    // fix the missing of openId
    user.OpenId = openId
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取用户", user))
	}
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		users := new(model.Users)
		if _, err = users.Load(tx); err != nil {
			logrus.Debug(err)
  		return c.JSON(fasthttp.StatusOK,
  			NewJSON("OK", "用户不存在", nil))
		}

		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取用户列表", users))
	}
}
