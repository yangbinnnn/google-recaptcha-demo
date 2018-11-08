package api

import (
	"io/ioutil"
	"net/http"

	"github.com/yangbinnnn/google-recaptcha-demo/core"

	"github.com/labstack/echo"
)

func verify(c echo.Context) error {
	resp, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	result := core.Verify(string(resp))
	return c.String(http.StatusOK, string(result))
}
