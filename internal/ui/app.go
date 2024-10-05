package ui

import (
	"log"

	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/ToffaKrtek/Ninjify/internal/ui/components"
	// "github.com/ToffaKrtek/Ninjify/internal/ui/components/buttons"
)

type (
  AppState struct {
    theme *material.Theme
    welcomeLabel  string
    usernameField widget.Editor
    passwordField widget.Editor
    loginButton widget.Clickable
  }
  )

func Start() {
  go func() {
    window := new(app.Window)
    state := AppState{
			theme: material.NewTheme(),
			welcomeLabel: "Добро пожаловать!",
		}
    err := run(window, &state)
    if err != nil {
      log.Fatal(err)
    }
  }()
  app.Main()
}

func run(window *app.Window,  state *AppState) error {
  theme := material.NewTheme()
  var ops op.Ops
  for {
    switch e := window.Event().(type) {
      case app.DestroyEvent:
        return e.Err
      // case app.FrameEvent:
      //   gtx := app.NewContext(&ops, e)
      //   title := material.H1(theme, "Start")
      //   maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
      //   title.Color = maroon
      //   title.Alignment = text.Middle
      //   title.Layout(gtx)
      //
      //   e.Frame(gtx.Ops)
      case app.FrameEvent:
        // buttons.SetSympleBtn(theme, &ops, e)
        components.SetMainComponent(theme, &ops, e)
        // gtx := app.NewContext(&ops, e)
        // // Обновляем окно
        // layoutUI(gtx, state)
        // e.Frame(gtx.Ops)
        window.Invalidate()
		}
  }
}
func layoutUI(gtx layout.Context, state *AppState) layout.Dimensions {
  paintBackground(gtx)
	// Центрируем форму
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.H6(state.theme, state.welcomeLabel).Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				state.usernameField.SingleLine = true
				editor := material.Editor(state.theme, &state.usernameField, "Имя пользователя")
        editor.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255} // Цвет текста
				// editor.Background = color.NRGBA{R: 0, G: 0, B: 0, A: 0} // Прозрачный фон
				return editor.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				state.passwordField.SingleLine = true
				state.passwordField.Mask = '*'
				editor := material.Editor(state.theme, &state.passwordField, "Пароль")
        editor.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255} // Цвет текста
				// editor.Background = color.NRGBA{R: 0, G: 0, B: 0, A: 0} // Прозрачный фон
				return editor.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				button := material.Button(state.theme, &state.loginButton, "Войти")
				button.Background = color.NRGBA{R: 0, G: 0, B: 0, A: 0} // Прозрачный фон
        button.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255} // Цвет текста
				return button.Layout(gtx)
			}),
		)
	})
}
// func gradientAnimation(gtx layout.Context) {
// 	// Пример анимации градиента (можно улучшить)
// 	rand.Seed(time.Now().UnixNano())
// 	// Здесь можно добавить логику для изменения цвета фона
// }

func (state *AppState) Layout(gtx layout.Context) layout.Dimensions {
	return layoutUI(gtx, state)
}
func paintBackground(gtx layout.Context) {
  ops := gtx.Ops
	paint.Fill(ops, color.NRGBA{R: 30, G: 30, B: 30, A: 255}) // Темный фон
  paint.PaintOp{}.Add(ops)
}

