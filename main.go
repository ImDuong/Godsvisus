package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type (
	Node struct {
		Value interface{}
		Next  *Node
	}

	NodeWraper struct {
		Data      Node
		Component *canvas.Circle
		Text      *canvas.Text
	}

	NodeWrapperList struct {
		Nodes       []*NodeWraper
		Connections []*canvas.Line
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
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)

	// hardcode hoziontal layout for a list
	distance := fyne.NewSize(radius+10, 0)
	for i, node := range ds.component.Nodes {
		accumulateDistance := fyne.NewSize(
			float32(i)*distance.Width,
			float32(i)*distance.Height,
		)

		node.Component.Resize(size)

		// get the position at the center of the node
		centerPos := fyne.NewPos(
			radius+accumulateDistance.Width,
			radius+accumulateDistance.Height,
		)

		// move the node to the right position: (center.X - radius, center.Y - radius)
		node.Component.Move(fyne.NewPos(
			centerPos.X-radius+accumulateDistance.Width,
			centerPos.Y-radius+accumulateDistance.Height,
		))

		// move the text to the center of the node
		node.Text.Move(fyne.NewPos(
			centerPos.X-node.Component.StrokeWidth-float32(len(node.Text.Text))*node.Text.TextSize/4+accumulateDistance.Width,
			centerPos.Y-node.Component.StrokeWidth-node.Text.TextSize/2+accumulateDistance.Height,
		))

		// add connecting line from the second node
		if i > 0 {
			// attach connnecting line's head position to the previous node
			ds.component.Connections[i-1].Position1 = fyne.NewPos(
				ds.component.Nodes[i-1].Component.Position2.X,
				centerPos.Y,
			)

			// attach connnecting line's tail position to the current node
			ds.component.Connections[i-1].Position2 = fyne.NewPos(
				node.Component.Position1.X,
				centerPos.Y,
			)
		}
	}
}

func (ds *dsLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	for i := range ds.component.Nodes {
		// add connecting line from the second node
		if i > 0 {
			line := canvas.Line{
				StrokeColor: theme.ForegroundColor(),
				StrokeWidth: 3,
			}
			ds.component.Connections = append(ds.component.Connections, &line)
			canvasObjs = append(canvasObjs, &line)
		}
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
			Nodes: []*NodeWraper{
				{
					Data: Node{
						Value: 10,
					},
				},
				{
					Data: Node{
						Value: 12,
					},
				},
				{
					Data: Node{
						Value: 69,
					},
				},
				{
					Data: Node{
						Value: 13,
					},
				},
				{
					Data: Node{
						Value: 5,
					},
				}, {
					Data: Node{
						Value: 753,
					},
				},
				{
					Data: Node{
						Value: 14,
					},
				},
				{
					Data: Node{
						Value: 28,
					},
				}, {
					Data: Node{
						Value: 11,
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
