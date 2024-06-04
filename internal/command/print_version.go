package command

import (
	"fmt"

	"github.com/Hydoc/goo/internal/view"
)

type PrintVersion struct {
	view    view.View
	version string
}

func (cmd *PrintVersion) Execute() {
	cmd.view.RenderLine(fmt.Sprintf("goo version %s", cmd.version))
}

func NewPrintVersion(view view.View, version string) (Command, error) {
	return &PrintVersion{
		view:    view,
		version: version,
	}, nil
}
