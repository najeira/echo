package echoutil

import (
	"github.com/labstack/echo"
	"github.com/najeira/conv"
)

func GetString(c echo.Context, name string, args ...string) string {
	return conv.String(c.Get(name), args...)
}

func GetInt(c echo.Context, name string, args ...int64) int64 {
	return conv.Int(c.Get(name), args...)
}

func GetFloat(c echo.Context, name string, args ...float64) float64 {
	return conv.Float(c.Get(name), args...)
}
