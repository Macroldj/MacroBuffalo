package grifts

import (
	"github.com/chenghung/buffalo_api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
