package app

import (
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

const (
	visusWindowWidth  float32 = 1000
	visusWindowHeight float32 = 500
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

func LoadLayout(loadMethod loadLayout, data interface{}, title string) (fyne.Layout, error) {
	if visusInstance == nil {
		InitVisusApp()
	}
	visusWindow := visusInstance.app.NewWindow(title)

	visObj, visLay, err := loadMethod(visusWindow, data)
	if err != nil {
		return nil, err
	}

	visusWindow.SetContent(visObj)
	visusWindow.Resize(fyne.NewSize(visusWindowWidth, visusWindowHeight))
	visusWindow.Show()
	return visLay, err
}
