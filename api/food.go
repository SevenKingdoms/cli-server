package api

import (
	"fmt"
	"strconv"
	// "fmt"
	"github.com/Sirupsen/logrus"
	"github.com/cli-server/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostFood() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		f := new(model.Food)
		fmt.Println(f)
		if err = c.Bind(f); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}

		tx := c.Get("Tx").(*dbr.Tx)

		food := model.NewFood(f.Id, f.Name, f.Image, f.Type, f.Price,
			f.HotIndex, f.Introduction, f.Merchant_id, f.InStock)

		if err := food.Save(tx); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusBadRequest,
				NewJSON("OK", "创建/更新食物失败", nil))
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建/更改食品", food))
	}
}

func GetFoodsByMerchantId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.QueryParam("merchant_id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		foods := new(model.Foods)
		count, err := foods.MerchantLoad(tx, number)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "商家不存在或者商家所拥有商品数量不足1", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家食品", foods))

	}
}
func GetFoodByFoodId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("food_id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		food := new(model.Food)
		count, err := food.Load(tx, number)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "食品不存在", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取食品", food))

	}
}

func DeleteFood() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("food_id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		food := new(model.Food)
		if err := food.FoodDelete(tx, number); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Food  does not exists.")
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功删除食品", food))

	}
}
