package entity

import (
	"errors"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type (
	ElementWrapper struct {
		Data        *Node
		Shape       *canvas.Rectangle
		Interaction *widget.Button
	}

	ElementWrapperList struct {
		Nodes []*ElementWrapper
	}
)

func NewElementWrapperList(data interface{}) (*ElementWrapperList, error) {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return nil, errors.New("input data is not a list")
	}
	dataSlice := reflect.ValueOf(data)
	eleList := ElementWrapperList{}

	for i := 0; i < dataSlice.Len(); i++ {
		eleList.Nodes = append(eleList.Nodes, &ElementWrapper{
			Data: &Node{
				Value: dataSlice.Index(i).Interface(),
			},
		})
	}
	return &eleList, nil
}

func (ew *ElementWrapper) Resize(s fyne.Size) {
	ew.Shape.Resize(s)
	ew.Interaction.Resize(s)
}

func (ew *ElementWrapper) Move(pos fyne.Position) {
	// move the shape
	ew.Shape.Move(pos)

	// move the button along with the shape
	ew.Interaction.Move(pos)
}
