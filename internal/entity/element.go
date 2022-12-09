package entity

import "fyne.io/fyne/v2/canvas"

type (
	ElementWrapper struct {
		Data Node
		// Component *canvas.Rectangle
		Component *canvas.Circle
		Text      *canvas.Text
	}

	ElementWrapperList struct {
		Nodes       []*ElementWrapper
		Connections []*canvas.Line
	}
)
