package grifts

import (
	"github.com/dkeza/expensesapp/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
