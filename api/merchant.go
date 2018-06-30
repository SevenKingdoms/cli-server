package api

import (
	// "strconv"

	"github.com/Sirupsen/logrus"
	"github.com/cli-server/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := new(model.Merchant)
		if err = c.Bind(m); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError, NewJSON("OK", "内部错误", nil))
		}

		tx := c.Get("Tx").(*dbr.Tx)

		merchant := model.NewMerchant(m.ID, m.HotIndex, m.Name,
			m.Introduction, m.Logo, m.Address, m.Images, m.Tel, m.Password, m.OpenTime, m.Open, m.Score, m.Onsales)

		if err := merchant.Save(tx); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusBadRequest,
				NewJSON("OK", "创建/更新商家失败", nil))
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建/更改商家", merchant))
	}
}

func GetMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tel := c.Param("tel")

		tx := c.Get("Tx").(*dbr.Tx)

		merchant := new(model.Merchant)
		if _, err := merchant.Load(tx, tel); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "商家不存在", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家", merchant))
	}
}

func GetAllMerchant() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		merchants := new(model.Merchants)
		if _, err = merchants.Load(tx); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "商家不存在", nil))
		}

		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家列表", merchants))
	}
}
