package linkedlist

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
	LinkedListLayout struct {
		component *entity.LinkedList
		canvas    fyne.CanvasObject
	}
)

func (lay *LinkedListLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *LinkedListLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// downsize 10 times
	size = fyne.NewSize(size.Width/10, size.Height/10)

	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)

	// hardcode hoziontal layout for a list
	distance := fyne.NewSize(radius+10, 0)

	curNode := lay.component.Root
	curIdx := 0
	for curNode != nil {
		accumulateDistance := fyne.NewSize(
			float32(curIdx)*distance.Width,
			float32(curIdx)*distance.Height,
		)

		curNode.Component.Resize(size)

		// get the position at the center of the node
		centerPos := fyne.NewPos(
			radius+accumulateDistance.Width,
			radius+accumulateDistance.Height,
		)

		// move the node to the right position: (center.X - radius, center.Y - radius)
		curNode.Component.Move(fyne.NewPos(
			centerPos.X-radius+accumulateDistance.Width,
			centerPos.Y-radius+accumulateDistance.Height,
		))

		// move the text to the center of the node
		curNode.Text.Move(fyne.NewPos(
			centerPos.X-curNode.Component.StrokeWidth-float32(len(curNode.Text.Text))*curNode.Text.TextSize/4+accumulateDistance.Width,
			centerPos.Y-curNode.Component.StrokeWidth-curNode.Text.TextSize/2+accumulateDistance.Height,
		))

		// add connecting line from the second node
		if curIdx > 0 {
			// attach connnecting line's head position to the previous node
			lay.component.Connections[curIdx-1].Position1 = fyne.NewPos(
				curNode.Prev.Component.Position2.X,
				centerPos.Y,
			)

			// attach connnecting line's tail position to the current node
			lay.component.Connections[curIdx-1].Position2 = fyne.NewPos(
				curNode.Component.Position1.X,
				centerPos.Y,
			)
		}
		curNode = curNode.Next
		curIdx++
	}
}

func (lay *LinkedListLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	curNode := lay.component.Root
	curIdx := 0
	for curNode != nil {
		// add connecting line from the second node
		if curIdx > 0 {
			line := canvas.Line{
				StrokeColor: theme.ForegroundColor(),
				StrokeWidth: 3,
			}
			lay.component.Connections = append(lay.component.Connections, &line)
			canvasObjs = append(canvasObjs, &line)
		}
		curNode.Component = &canvas.Circle{
			StrokeColor: color.White,
			StrokeWidth: 5,
		}
		curNode.Text = &canvas.Text{
			Text:     fmt.Sprintf("%v", curNode.Data.Value),
			Color:    color.White,
			TextSize: 12,
			TextStyle: fyne.TextStyle{
				Bold: true,
			},
		}
		canvasObjs = append(canvasObjs, curNode.Component, curNode.Text)

		curNode = curNode.Next
		curIdx++
	}

	container := container.NewWithoutLayout(canvasObjs...)
	container.Layout = lay

	lay.canvas = container
	return container
}

func Load(win fyne.Window) fyne.CanvasObject {
	rootNode := &entity.Node{
		Value: 12,
		Next: &entity.Node{
			Value: 3,
			Next: &entity.Node{
				Value: 69,
			},
		},
	}

	rootNodeWrapper := entity.NewNodeWrapper(rootNode)

	lay := &LinkedListLayout{
		component: &entity.LinkedList{
			Root: rootNodeWrapper,
		},
	}

	container := lay.render()
	return container
}
