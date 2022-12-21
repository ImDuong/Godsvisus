package main

import (
	"github.com/ImDuong/godsvisus/internal/entity"
	"github.com/ImDuong/godsvisus/visualize/linkedlist"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gods Visus Linked Lists Comparison")

	visObj, visLay, err := linkedlist.Load(myWindow, []*entity.Node{
		{
			Value: 12,
			Next: &entity.Node{
				Value: 3,
				Next: &entity.Node{
					Value: 69,
				},
			},
		},
		{
			Value: 4,
			Next: &entity.Node{
				Value: 3,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	myWindow.SetContent(visObj)
	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.Show()
	visLay.(*linkedlist.LinkedListLayout).Compare()
	myApp.Run()
}
