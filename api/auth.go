package api


import (
  // "fmt"
  "time"
  "io/ioutil"

	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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
		return c.JSON(fasthttp.StatusOK,
      NewJSON("OK", "成功获取", string(body)))
	}
}

func GetJWT() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

    username := c.QueryParam("username")
    // password := c.QueryParam("password")
    userType := c.QueryParam("type")

    valid := true

  	if valid {
  		// Set claims
  		claims := &JWTCustomClaims{
  			username,
  			userType == "1",
  			jwt.StandardClaims{
  				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
  			},
  		}

  		// Create token with claims
  		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  		// Generate encoded token and send it as response.
  		t, err := token.SignedString([]byte("secret"))
  		if err != nil {
  			return err
  		}
      return c.JSON(fasthttp.StatusOK,
        NewJSON("OK", "验证成功", echo.Map{
    			"token": t,
    		}))
  	}

  	return echo.ErrUnauthorized
	}
}
