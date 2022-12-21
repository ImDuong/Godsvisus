package app

import (
	"godsvisus/visualize/array"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type (
	visusApp struct {
		app fyne.App
	}

	loadLayout func(fyne.Window, interface{}) (fyne.CanvasObject, fyne.Layout, error)
)

var visusInstance *visusApp
var once sync.Once

func InitVisusApp() {
	once.Do(func() {
		visusInstance = &visusApp{
			app: app.New(),
		}
	})
}

func Run() {
	visusInstance.app.Run()
}

func LoadLayout(loadMethod loadLayout, data interface{}, title string) error {
	visusWindow := visusInstance.app.NewWindow(title)

	visObj, _, err := array.Load(visusWindow, data)
	if err != nil {
		return err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(1000, 500))
	visusWindow.Show()
	return err
}
