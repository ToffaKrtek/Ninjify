package forms

import "gioui.org/layout"

type Form struct {
  Children []layout.FlexChild
}

func (f *Form) Layout(gtx layout.Context) layout.Dimensions {
  return layout.Flex{Axis: layout.Vertical}.Layout(gtx, f.Children...)
}
