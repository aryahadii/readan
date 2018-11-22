package simplifier

import (
	"net/url"
)

type Simplifier interface {
	SimplifyHTML(url *url.URL) (string, error)
}
