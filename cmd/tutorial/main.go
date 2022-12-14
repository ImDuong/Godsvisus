package main

import (
	"fmt"
	"godsvisus/internal/entity"
	"godsvisus/visualize/array"
	"godsvisus/visualize/linkedlist"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type appInfo struct {
	name string
	icon fyne.Resource
	canv bool
	data interface{}
	run  func(fyne.Window, interface{}) (fyne.CanvasObject, error)
}
type teststruct struct {
	Value int
	Thing string
}

var test = []teststruct{
	{
		Value: 3,
		Thing: "kuro",
	},
	{
		Value: 2,
		Thing: "shit",
	},
	{
		Value: 1,
		Thing: "hehe",
	},
}

// var test = [3]int{10, 2, 69}

var apps = []appInfo{
	{"Linked List", nil, false, &entity.Node{
		Value: 12,
		Next: &entity.Node{
			Value: 3,
			Next: &entity.Node{
				Value: 69,
			},
		},
	}, linkedlist.Load},
	{"Array", nil, false, test, array.Load},
}

func main() {
	fmt.Printf("%p\n", &test)
	fmt.Printf("%p\n", &test[0])
	fmt.Printf("%p\n", &test[1])
	fmt.Printf("%p\n", &test[2])

	myApp := app.New()
	myWindow := myApp.NewWindow("Gods Visus Tutorial")

	content := container.NewMax()

	appList := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height
			icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
			return container.NewBorder(nil, nil, icon, nil,
				label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = apps[id].icon
			img.Refresh()
			text.SetText(apps[id].name)
		})
	appList.OnSelected = func(id widget.ListItemID) {
		canvasObj, err := apps[id].run(myWindow, apps[id].data)
		if err != nil {
			log.Fatal("Error when drawing gadget", err)
		} else {
			content.Objects = []fyne.CanvasObject{canvasObj}
		}
	}

	split := container.NewHSplit(appList, content)
	split.Offset = 0.1
	myWindow.SetContent(split)
	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.ShowAndRun()
}
