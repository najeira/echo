// +build go1.8

package echoutil

import "net/url"

func pathUnescape(s string) (string, error) {
	return url.PathUnescape(s)
}
