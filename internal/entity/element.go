package entity

import (
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

func (ew *ElementWrapper) Refresh() {
	ew.Shape.Refresh()
	ew.Interaction.Refresh()
}

func (ewl *ElementWrapperList) Refresh() {
	for i := range ewl.Nodes {
		ewl.Nodes[i].Refresh()
	}
}
