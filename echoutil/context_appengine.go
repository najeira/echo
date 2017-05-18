// +build appengine

package echoutil

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

const (
	contextNetContextKey = "github.com/najeira/echo/echoutil/context"
)

// Context gets a net context from echo context.
func Context(c echo.Context) context.Context {
	obj := c.Get(contextNetContextKey)
	if obj != nil {
		if ctx, ok := obj.(context.Context); ok {
			return ctx
		}
	}
	ctx := appengine.NewContext(c.Request())
	c.Set(contextNetContextKey, ctx)
	return ctx
}
