package buttons

import (
  "gioui.org/op"
  "gioui.org/app"
  "gioui.org/widget"
  "gioui.org/widget/material"
)

var sympleBtn widget.Clickable

func SetSympleBtn(
  theme *material.Theme,
  ops *op.Ops,
  event app.FrameEvent,
) {
  gtx := app.NewContext(ops, event)
  btn := material.Button(theme, &sympleBtn, "Symple")
  btn.Layout(gtx)
  event.Frame(gtx.Ops)
}
