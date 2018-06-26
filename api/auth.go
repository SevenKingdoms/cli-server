package api


import (
  // "fmt"
  "io/ioutil"

	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetOpenid() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

    code := c.QueryParam("code")
    if code == "" {
  		return c.JSON(fasthttp.StatusBadRequest,
        NewJSON("OK", "获取失败", "code: " + code))
    }
    resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx16dd2967089b83ab&secret=0f0aace6fd928cb70cb3a0338ed20670&js_code=" + code + "&grant_type=authorization_code")
    if err != nil {
  		return c.JSON(fasthttp.StatusBadRequest,
        NewJSON("OK", "获取失败", "code: " + code))
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
    }
		return c.JSON(fasthttp.StatusCreated,
      NewJSON("OK", "成功获取", string(body)))
	}
}
