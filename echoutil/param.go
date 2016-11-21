package echoutil

import (
	"github.com/labstack/echo"
	"github.com/najeira/conv"
)

var (
	getters = []getter{
		echo.Context.Param,
		echo.Context.QueryParam,
		echo.Context.FormValue,
	}
)

type getter func(echo.Context, string) string

func getParam(c echo.Context, name string) string {
	for _, getter := range getters {
		if v := getter(c, name); len(v) > 0 {
			return v
		}
	}
	return ""
}

func ParamString(c echo.Context, name string, args ...string) string {
	v := getParam(c, name)
	if len(v) > 0 {
		return v
	} else if len(args) > 0 {
		return args[0]
	}
	return ""
}

func ParamInt(c echo.Context, name string, args ...int64) int64 {
	return conv.Int(getParam(c, name), args...)
}

func ParamFloat(c echo.Context, name string, args ...float64) float64 {
	return conv.Float(getParam(c, name), args...)
}
