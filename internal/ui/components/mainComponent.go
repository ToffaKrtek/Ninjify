package components

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/ToffaKrtek/Ninjify/internal/ui/components/buttons"
	"github.com/ToffaKrtek/Ninjify/internal/ui/components/fields"
	"github.com/ToffaKrtek/Ninjify/internal/ui/components/forms"
	"github.com/ToffaKrtek/Ninjify/internal/ui/components/blocks"
)

func SetMainComponent(
  theme *material.Theme,
  ops *op.Ops,
  event app.FrameEvent,
) {
  var inputField *fields.InputField
  var submitButton *buttons.Button
  var form *forms.Form

  inputField = fields.NewInputField(theme, "Enter something:")
	submitButton = buttons.NewButton(theme, "Submit")

  form = &forms.Form{
			Children: []layout.FlexChild{
				layout.Rigid(inputField.Layout),
				layout.Rigid(submitButton.Layout),
			},
		}
  gtx := app.NewContext(ops, event)
  layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		card := &blocks.CardBlock{
			Title:    "My Form",
			Children: []layout.FlexChild{layout.Rigid(form.Layout)},
		}
	  return card.Layout(theme, gtx)
	})
  event.Frame(gtx.Ops)
}
