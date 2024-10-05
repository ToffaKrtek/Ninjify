package fields

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type InputField struct {
  Widget *widget.Editor
  Theme *material.Theme
  Label string
}

func NewInputField(th *material.Theme, label string) *InputField {
  return &InputField{
    Widget: new(widget.Editor),
    Theme: th,
    Label: label,
  }
}

func (f *InputField) Layout(gtx layout.Context) layout.Dimensions {
  return layout.Flex{Axis: layout.Vertical}.Layout(gtx, 
      layout.Rigid(func(gtx layout.Context) layout.Dimensions {
      return material.Label(f.Theme, 16, f.Label).Layout(gtx)
    }),
      layout.Rigid(func(gtx layout.Context) layout.Dimensions {
      return material.Editor(f.Theme, f.Widget, "").Layout(gtx)
    }),
  )
}
