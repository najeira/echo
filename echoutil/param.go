package echoutil

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/najeira/conv"
)

var (
	getters = []getter{
		getPathParam,
		echo.Context.QueryParam,
		echo.Context.FormValue,
	}
)

type getter func(echo.Context, string) string

func getPathParam(c echo.Context, name string) string {
	v := echo.Context.Param(c, name)
	if len(v) <= 0 {
		return v
	}
	u, err := url.PathUnescape(v)
	if err != nil {
		return v
	}
	return u
}

func getParam(c echo.Context, name string) string {
	for _, getter := range getters {
		if v := getter(c, name); len(v) > 0 {
			return v
		}
	}
	return ""
}

func ParamString(c echo.Context, name string, args ...string) string {
	if v := getParam(c, name); len(v) > 0 {
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
