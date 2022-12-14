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
		Data        interface{}
		DataAddr    uintptr
		Shape       *canvas.Rectangle
		Interaction *widget.Button
	}

	ElementWrapperList struct {
		Nodes []*ElementWrapper
	}
)

func NewElementWrapperList(data interface{}) (*ElementWrapperList, error) {
	dataKind := reflect.TypeOf(data).Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return nil, errors.New("input data is not a slice")
	}
	dataSlice := reflect.ValueOf(data)

	eleList := ElementWrapperList{}
	eleList.Nodes = make([]*ElementWrapper, dataSlice.Len())

	for i := 0; i < dataSlice.Len(); i++ {
		eleList.Nodes[i] = &ElementWrapper{
			Data: dataSlice.Index(i).Interface(),
		}
		if dataSlice.Index(i).CanAddr() {
			eleList.Nodes[i].DataAddr = dataSlice.Index(i).UnsafeAddr()
		}
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
