package echoutil

import (
	"github.com/labstack/echo"
	"github.com/najeira/conv"
)

func getParams(c echo.Context, name string) []string {
	if vs := c.QueryParams(); len(vs) > 0 {
		if ps, ok := vs[name]; ok {
			return ps
		}
	}
	if vs, _ := c.FormParams(); len(vs) > 0 {
		if ps, ok := vs[name]; ok {
			return ps
		}
	}
	return nil
}

func ParamStrings(c echo.Context, name string) []string {
	return getParams(c, name)
}

func ParamInts(c echo.Context, name string) []int64 {
	ps := getParams(c, name)
	if len(ps) <= 0 {
		return nil
	}
	rv := make([]int64, len(ps))
	for i, p := range ps {
		rv[i] = conv.Int(p)
	}
	return rv
}

func ParamFloats(c echo.Context, name string) []float64 {
	ps := getParams(c, name)
	if len(ps) <= 0 {
		return nil
	}
	rv := make([]float64, len(ps))
	for i, p := range ps {
		rv[i] = conv.Float(p)
	}
	return rv
}
