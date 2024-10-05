package buttons

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Button struct {
	Widget *widget.Clickable
	Theme  *material.Theme
	Label  string
}

func NewButton(th *material.Theme, label string) *Button {
  return &Button{
    Widget: new(widget.Clickable),
    Theme: th,
    Label: label,
  }
}

func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
  return material.Button(b.Theme, b.Widget, b.Label).Layout(gtx)
}
