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
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		tx := c.Get("Tx").(*dbr.Tx)

		user := model.NewUser(m.OpenId, m.Name, m.Avatar, m.Phone)

		if err := user.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusCreated,
<<<<<<< HEAD
			NewJSON("OK", "成功创建用户", user))
=======
      NewJSON("OK", "成功创建/更新用户", user))
>>>>>>> 2ba8a575b266b4111ad27ae29e719e7bb93efd75
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		openId := c.Param("open_id")

		tx := c.Get("Tx").(*dbr.Tx)

		user := new(model.User)
		if _, err := user.Load(tx, openId); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "User does not exists.")
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
			return echo.NewHTTPError(fasthttp.StatusNotFound, "User does not exists.")
		}

		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取用户列表", users))
	}
}
