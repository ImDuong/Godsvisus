package main

import (
	"godsvisus/visualize/array"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gods Visus List Comparison")

	visObj, visLay, err := array.Load(myWindow, [][]int{
		{6, 3, 100, 55, 87, -4},
		{6, 9, 100, 55, 12, 33},
	})
	if err != nil {
		panic(err)
	}

	myWindow.SetContent(visObj)
	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.Show()
	visLay.(*array.ArrayLayout).Compare()
	myApp.Run()
}
