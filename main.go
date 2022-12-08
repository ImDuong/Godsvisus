package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type (
	Node struct {
		Value interface{}
	}

	NodeWraper struct {
		Data      Node
		Component *canvas.Circle
		Text      *canvas.Text
	}

	NodeWrapperList struct {
		Nodes []NodeWraper
	}

	dsLayout struct {
		component NodeWrapperList
		canvas    fyne.CanvasObject
	}
)

func (ds *dsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (ds *dsLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	for _, node := range ds.component.Nodes {
		diameter := fyne.Min(size.Width, size.Height)
		radius := diameter / 2
		size = fyne.NewSize(diameter, diameter)
		centerPos := fyne.NewPos(size.Width/2, size.Height/2)

		node.Component.Resize(size)
		node.Component.Move(fyne.NewPos(centerPos.X-radius, centerPos.Y-radius))

		node.Text.Move(fyne.NewPos(radius-node.Component.StrokeWidth, radius-node.Component.StrokeWidth-node.Text.TextSize))
	}

}

func (ds *dsLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	for i := range ds.component.Nodes {
		ds.component.Nodes[i].Component = &canvas.Circle{
			StrokeColor: color.White,
			StrokeWidth: 5,
		}
		ds.component.Nodes[i].Text = &canvas.Text{
			Text:     fmt.Sprintf("%v", ds.component.Nodes[i].Data.Value),
			Color:    color.White,
			TextSize: 12,
			TextStyle: fyne.TextStyle{
				Bold: true,
			},
		}
		canvasObjs = append(canvasObjs, ds.component.Nodes[i].Component, ds.component.Nodes[i].Text)
	}

	container := container.NewWithoutLayout(canvasObjs...)
	container.Layout = ds

	ds.canvas = container
	return container
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Node Visus")

	ds := &dsLayout{
		component: NodeWrapperList{
			Nodes: []NodeWraper{
				{
					Data: Node{
						Value: 1,
					},
				},
				{
					Data: Node{
						Value: 2,
					},
				},
			},
		},
	}
	content := ds.render()

	text := canvas.NewText("Hello world to visus", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: true}
	box := container.NewVBox(
		text,
		content,
	)
	myWindow.SetContent(box)
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}
