package array

import (
	"fmt"
	"image/color"

	"godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type (
	ArrayLayout struct {
		component entity.ElementWrapperList
		canvas    fyne.CanvasObject
	}
)

func (lay *ArrayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *ArrayLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)

	// hardcode hoziontal layout for a list
	distance := fyne.NewSize(radius+10, 0)
	for i, node := range lay.component.Nodes {
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
			lay.component.Connections[i-1].Position1 = fyne.NewPos(
				lay.component.Nodes[i-1].Component.Position2.X,
				centerPos.Y,
			)

			// attach connnecting line's tail position to the current node
			lay.component.Connections[i-1].Position2 = fyne.NewPos(
				node.Component.Position1.X,
				centerPos.Y,
			)
		}
	}
}

func (lay *ArrayLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	for i := range lay.component.Nodes {
		// add connecting line from the second node
		if i > 0 {
			line := canvas.Line{
				StrokeColor: theme.ForegroundColor(),
				StrokeWidth: 3,
			}
			lay.component.Connections = append(lay.component.Connections, &line)
			canvasObjs = append(canvasObjs, &line)
		}
		lay.component.Nodes[i].Component = &canvas.Circle{
			StrokeColor: color.White,
			StrokeWidth: 5,
		}
		lay.component.Nodes[i].Text = &canvas.Text{
			Text:     fmt.Sprintf("%v", lay.component.Nodes[i].Data.Value),
			Color:    color.White,
			TextSize: 12,
			TextStyle: fyne.TextStyle{
				Bold: true,
			},
		}
		canvasObjs = append(canvasObjs, lay.component.Nodes[i].Component, lay.component.Nodes[i].Text)
	}

	container := container.NewWithoutLayout(canvasObjs...)
	container.Layout = lay

	lay.canvas = container
	return container
}

func Load(win fyne.Window) fyne.CanvasObject {
	lay := &ArrayLayout{
		component: entity.ElementWrapperList{
			Nodes: []*entity.ElementWrapper{
				{
					Data: entity.Node{
						Value: 10,
					},
				},
				{
					Data: entity.Node{
						Value: 2,
					},
				},
				{
					Data: entity.Node{
						Value: 3,
					},
				},
			},
		},
	}

	container := lay.render()
	return container
}
