package godsvisus

import (
	"godsvisus/visualize/array"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func ShowArray(data interface{}) error {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("Gods Visus Lists Comparison")

	visObj, _, err := array.Load(visusWindow, []interface{}{
		data,
	})
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()

	go visusApp.Run()
	return nil
}
