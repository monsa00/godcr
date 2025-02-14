package page

import (
	"gioui.org/layout"

	"github.com/planetdecred/godcr/ui/page/components"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	MaxWidth = components.MaxWidth
)

// todo: this method will be removed when the new modal implementation is used on the seedbackup page
func _modal(gtx layout.Context, body layout.Dimensions, modal layout.Dimensions) layout.Dimensions {
	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return body
		}),
		layout.Stacked(func(gtx C) D {
			return modal
		}),
	)
	return dims
}
