package godsvisus

import (
	visusapp "godsvisus/internal/app"
	"godsvisus/visualize/array"
	"godsvisus/visualize/linkedlist"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Init() error {
	visusapp.InitVisusApp()
	return nil
}

func Run() {
	visusapp.Run()
}

func ShowArrays(data interface{}) error {
	err := visusapp.LoadLayout(array.Load, data, "Gods Visus: Lists")
	if err != nil {
		return err
	}
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

func ShowLinkedLists(data interface{}) error {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("Gods Visus: Linked Lists")

	visObj, _, err := linkedlist.Load(visusWindow, data)
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()

	visusApp.Run()
	return nil
}

func CompareLinkedLists(data interface{}) error {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("Gods Visus: Lists Comparison")

	visObj, visLay, err := linkedlist.Load(visusWindow, data)
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()
	visLay.(*linkedlist.LinkedListLayout).Compare()
	visusApp.Run()
	return nil
}
