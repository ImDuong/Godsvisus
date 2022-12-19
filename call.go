package godsvisus

import (
	"godsvisus/visualize/array"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func ShowArrays(data interface{}) error {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("Gods Visus: List")

	visObj, _, err := array.Load(visusWindow, data)
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()

	visusApp.Run()
	return nil
}

func CompareArrays(data interface{}) error {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("Gods Visus: Lists Comparison")

	visObj, visLay, err := array.Load(visusWindow, data)
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()
	visLay.(*array.ArrayLayout).Compare()
	visusApp.Run()
	return nil
}
