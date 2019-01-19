package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/nielsdingsbums/dwb/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
