package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/cli-server/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetAllMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
    tx := c.get("Tx").(*dbr.Tx)

    merchants := new(model.Merchants)
    if _, err = merchants.Load(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Merchants does not exists.")
		}

		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家列表", merchants))
  }
}

func PostMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
    m := new(model.Merchant)
		if err = c.Bind(m); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		tx := c.Get("Tx").(*dbr.Tx)

		merchant := model.NewMerchant(m.ID, m.Name, m.HotIndex,
       m.Introduction, m.Logo, m.Images, m.Account, m.Password)

		if err := merchant.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建商家", merchant))


  }
}

func GetMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
    id := c.Param("id")

		tx := c.Get("Tx").(*dbr.Tx)

		merchant := new(model.Merchant)
		if _, err := merchant.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Merchant does not exists.")
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家", merchant))

  }
}

func UpdateMerchant() func GetAllMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {


  }
}
