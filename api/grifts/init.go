package grifts

import (
	"github.com/YAWAL/workshops/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
