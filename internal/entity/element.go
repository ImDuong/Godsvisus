package entity

import (
	"errors"
	"reflect"

	"fyne.io/fyne/v2/canvas"
)

type (
	ElementWrapper struct {
		Data      *Node
		Component *canvas.Rectangle
		Text      *canvas.Text
	}

	ElementWrapperList struct {
		Nodes       []*ElementWrapper
		Connections []*canvas.Line
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
