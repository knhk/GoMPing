package gomping

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	mText1 = `Manual`
	mText2 = `[green]       Ctrl+M: [white]Open manual page.(this page)`
	mText3 = `[green]       Ctrl+Q: [white]Return to first page.       `
	mText4 = `[green]    ←, Ctrl+P: [white]Move to previous page.      `
	mText5 = `[green]    →, Ctrl+N: [white]Move to next page.          `
	mText6 = `[green]  ↑, wheel-up: [white]Scroll up for ping list.    `
	mText7 = `[green]↓, wheel-down: [white]Scroll down for ping list.  `
)

func DrawManual() tview.Primitive {

	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(mText1, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(mText2, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mText3, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mText4, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mText5, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mText6, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mText7, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	tl := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(frame, 0, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tl, 0, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)

	return layout
}
