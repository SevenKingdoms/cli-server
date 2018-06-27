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
	"github.com/cli-server/model"
  "github.com/gocraft/dbr"
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
    password := c.QueryParam("password")
    userType := c.QueryParam("type")

    if valid := ValidUser(c, username, password, userType); valid == true {
  		// Set claims
  		claims := &JWTCustomClaims{
  			username,
        password,
  			userType == "0",
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
  	} else {
      return c.JSON(fasthttp.StatusBadRequest,
        NewJSON("Not OK", "用户验证失败", nil))
    }

  	return echo.ErrUnauthorized
	}
}

func ValidUser(c echo.Context, username, password, userType string) bool {

	openId := username

	tx := c.Get("Tx").(*dbr.Tx)

  if userType == "0" { // admin
    // TODO:
  } else if userType == "1" { // user
		user := new(model.User)
    count, err := user.Load(tx, openId)
    if err != nil {
			logrus.Debug(err)
      return false
		}
    if count == 0 {
      return false
    }
    return true
  } else { // merchant
    // TODO:
  }
  return true
}
