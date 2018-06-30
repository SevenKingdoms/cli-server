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

func PostComment() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		f := new(model.Comment)
		fmt.Println(f)
		if err = c.Bind(f); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}

		tx := c.Get("Tx").(*dbr.Tx)

		comment := model.NewComment(f.Id, f.Description,f.Images,f.User_openId,f.Merchant_id,f.Score)

		if err := comment.Save(tx); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusBadRequest,
				NewJSON("OK", "创建/更新评论失败", nil))
		}
		return c.JSON(fasthttp.StatusCreated,
			NewJSON("OK", "成功创建/更改评论", comment))
	}
}

func GetCommentsByMerchantId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.QueryParam("merchant_id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		comments := new(model.Comments)
		count, err := comments.MerchantLoad(tx, number)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "商家不存在或者商家所拥有评论数为0", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取商家评论", comments))

	}
}

func GetCommentsByUserId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.QueryParam("User_openId")

		tx := c.Get("Tx").(*dbr.Tx)

		comments := new(model.Comments)
		count, err := comments.UserIdLoad(tx, id)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "用户不存在或者用户所拥有评论数为0", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取用户评论", comments))

	}
}
func GetCommentByCommentId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		comment := new(model.Comment)
		count, err := comment.Load(tx, number)
		if err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusInternalServerError,
				NewJSON("OK", "内部错误", nil))
		}
		if count == 0 {
			return c.JSON(fasthttp.StatusOK,
				NewJSON("OK", "评论不存在", nil))
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功获取评论", comment))

	}
}

func DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.Param("id")
		number, _ := strconv.ParseInt(id, 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		comment := new(model.Comment)
		if err := comment.CommentDelete(tx, number); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "评论  does not exists.")
		}
		return c.JSON(fasthttp.StatusOK,
			NewJSON("OK", "成功删除评论", comment))

	}
}
