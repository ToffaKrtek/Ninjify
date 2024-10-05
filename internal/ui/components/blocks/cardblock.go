package blocks

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type CardBlock struct {
  Title string
  Children []layout.FlexChild
}

func (b *CardBlock) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
  return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
      return material.H6(th, b.Title).Layout(gtx)
    }),
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
      return layout.Flex{Axis: layout.Vertical}.Layout(gtx, b.Children...)
    }),
  )
}
