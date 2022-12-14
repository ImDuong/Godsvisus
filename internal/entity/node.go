package entity

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type (
	Node struct {
		Value interface{}
		Next  *Node
	}

	NodeWrapper struct {
		Data        *Node
		Shape       *canvas.Circle
		Interaction *widget.Button
		Next        *NodeWrapper
		Prev        *NodeWrapper
	}

	LinkedList struct {
		Root        *NodeWrapper
		Connections []*canvas.Line
	}
)

func (nw *NodeWrapper) Resize(s fyne.Size) {
	nw.Shape.Resize(s)
	nw.Interaction.Resize(s)
}

func (nw *NodeWrapper) Move(pos fyne.Position) {
	// move the shape
	nw.Shape.Move(pos)

	// move the button along with the shape
	nw.Interaction.Move(pos)
}
