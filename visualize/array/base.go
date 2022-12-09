package array

import (
	"fmt"
	"image/color"

	"godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type (
	ArrayLayout struct {
		component *entity.ElementWrapperList
		canvas    fyne.CanvasObject
	}
)

func (lay *ArrayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *ArrayLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// downsize 10 times
	size = fyne.NewSize(size.Width/10, size.Height/10)

	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)

	// hardcode hoziontal layout for a list
	distance := fyne.NewSize(radius, 0)
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
	}
}

func (lay *ArrayLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	for i := range lay.component.Nodes {
		lay.component.Nodes[i].Component = &canvas.Rectangle{
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

func Load(win fyne.Window, data interface{}) (fyne.CanvasObject, error) {
	eleList, err := entity.NewElementWrapperList(data)
	if err != nil {
		return nil, err
	}

	lay := &ArrayLayout{
		component: eleList,
	}

	container := lay.render()
	return container, nil
}
