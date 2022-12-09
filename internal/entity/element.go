package entity

import "fyne.io/fyne/v2/canvas"

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

func NewElementWrapperList(data []int) *ElementWrapperList {
	eleList := ElementWrapperList{}
	for i := range data {
		eleList.Nodes = append(eleList.Nodes, &ElementWrapper{
			Data: &Node{
				Value: data[i],
			},
		})
	}
	return &eleList
}
