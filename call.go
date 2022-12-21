package godsvisus

import (
	visusapp "godsvisus/internal/app"
	"godsvisus/visualize/array"
	"godsvisus/visualize/linkedlist"

	"fyne.io/fyne/v2"
)

func Init() {
	visusapp.InitVisusApp()
}

func Run() {
	visusapp.Run()
}

func ShowArrays(data interface{}) (fyne.Layout, error) {
	visLay, err := visusapp.LoadLayout(array.Load, data, "Gods Visus: Lists")
	if err != nil {
		return nil, err
	}
	return visLay, nil
}

func CompareArrays(data interface{}) (fyne.Layout, error) {
	visLay, err := visusapp.LoadLayout(array.Load, data, "Gods Visus: Lists Comparison")
	if err != nil {
		return nil, err
	}
	return visLay, visLay.(*array.ArrayLayout).Compare()
}

func CompareDisplayedArrays(visLay fyne.Layout) error {
	return visLay.(*array.ArrayLayout).Compare()
}

func ShowLinkedLists(data interface{}) (fyne.Layout, error) {
	visLay, err := visusapp.LoadLayout(linkedlist.Load, data, "Gods Visus: Linked Lists")
	if err != nil {
		return nil, err
	}
	return visLay, nil
}

func CompareLinkedLists(data interface{}) (fyne.Layout, error) {
	visLay, err := visusapp.LoadLayout(linkedlist.Load, data, "Gods Visus: Linked Lists Comparison")
	if err != nil {
		return nil, err
	}
	return visLay, visLay.(*linkedlist.LinkedListLayout).Compare()
}

func CompareDisplayedLinkedLists(visLay fyne.Layout) error {
	return visLay.(*linkedlist.LinkedListLayout).Compare()
}
