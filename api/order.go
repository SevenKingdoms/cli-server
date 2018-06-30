package api

import (
	// "strconv"
	// "fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/cli-server/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostOrder() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := new(model.Order)
		if err = c.Bind(m); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}

		tx := c.Get("Tx").(*dbr.Tx)

		order := model.NewOrder(m.Id, m.NumOfPeople, m.DeskId, m.Remark, m.Paid, m.User_openId, m.Merchant_id, m.Foods, m.Merchant_name, m.Merchant_tel)

		if err := order.Save(tx); err != nil {

			logrus.Debug(err)
			return c.JSON(fasthttp.StatusBadRequest,
				NewJSON("OK", "创建/更新订单失败", nil))
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建/更新订单", order))
	}
}

func GetOrderByOrderId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		order_id := c.Param("order_id")
		//order_id  int64
		number, _ := strconv.ParseInt(order_id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		order := new(model.Order)
		count, err := order.Load(tx, number)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError, NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "订单不存在", nil))
		}
		// fix the missing of openId
		//user.OpenId = openId
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取订单", order))
	}
}

func GetOrdersBySomeId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
    merchant_id := c.QueryParam("merchant_id")
    open_id := c.QueryParam("open_id")
    tx := c.Get("Tx").(*dbr.Tx)

    if merchant_id != "" {
      // 商家
  		id := merchant_id
  		number, _ := strconv.ParseInt(id, 0, 64)
  		orders := new(model.Orders)
  		count, err := orders.MerchantIdLoad(tx, number)

  		if err != nil {
  			logrus.Debug(err)
  			return c.JSON(fasthttp.StatusInternalServerError,
  				NewJSON("OK", "内部错误", nil))
  		}
  		if count == 0 {
  			return c.JSON(fasthttp.StatusOK,
  				NewJSON("OK", "商家还未有订单", nil))
  		}
  		return c.JSON(fasthttp.StatusOK,
  			NewJSON("OK", "成功获取商家订单列表", orders))
    } else {
      // 用户
  		id := open_id

  		orders := new(model.Orders)
  		count, err := orders.OpenIdLoad(tx, id)

  		if err != nil {
  			logrus.Debug(err)
  			return c.JSON(fasthttp.StatusInternalServerError,
  				NewJSON("OK", "内部错误", nil))
  		}
  		if count == 0 {
  			return c.JSON(fasthttp.StatusOK,
  				NewJSON("OK", "用户还未有订单", nil))
  		}
  		return c.JSON(fasthttp.StatusOK,
  			NewJSON("OK", "成功获取用户订单列表", orders))
    }
	}
}

func DeleteOrderByOrderId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("order_id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		order := new(model.Order)
		if err := order.OrderDelete(tx, number); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "订单不存在", order))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功删除订单", order))
	}
}
