package linkedlist

import (
	"errors"
	"fmt"
	"image/color"

	"godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type (
	LinkedListLayout struct {
		component *entity.LinkedList
		detail    *entity.NodeInfo
		canvas    fyne.CanvasObject
	}
)

func (lay *LinkedListLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *LinkedListLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// downsize 10 times
	var downRatio float32 = 10

	// calculate diameter, radius and new size for a node
	diameter := fyne.Min(size.Width, size.Height)
	if diameter > lay.canvas.MinSize().Width*downRatio {
		diameter = diameter / downRatio
	}
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)

	// hardcode the length of connecting lines between nodes
	var conLineLen float32 = 10
	distance := fyne.NewSize(radius+conLineLen, 0)

	curNode := lay.component.Root
	curIdx := 0
	for curNode != nil {
		curNode.Resize(size)

		accumulateDistance := fyne.NewSize(
			float32(curIdx)*distance.Width,
			float32(curIdx)*distance.Height,
		)

		// get the position at the center of the node
		centerPos := fyne.NewPos(
			radius+accumulateDistance.Width,
			radius+accumulateDistance.Height,
		)

		// move the node to the right position: (center.X - radius, center.Y - radius)
		curNode.Move(fyne.NewPos(
			centerPos.X-radius+accumulateDistance.Width,
			centerPos.Y-radius+accumulateDistance.Height,
		))

		// add connecting line from the second node
		if curIdx > 0 {
			// attach connnecting line's head position to the previous node
			lay.component.Connections[curIdx-1].Position1 = fyne.NewPos(
				curNode.Prev.Shape.Position2.X,
				centerPos.Y,
			)

			// attach connnecting line's tail position to the current node
			lay.component.Connections[curIdx-1].Position2 = fyne.NewPos(
				curNode.Shape.Position1.X,
				centerPos.Y,
			)
		}
		curNode = curNode.Next
		curIdx++
	}
}

func (lay *LinkedListLayout) render(data interface{}) (*fyne.Container, error) {
	// validate input
	dataNode, ok := data.(*entity.Node)
	if !ok {
		return nil, errors.New("input is not a valid node")
	}

	// init
	canvasObjs := []fyne.CanvasObject{}
	lay.component = &entity.LinkedList{}

	curNode := dataNode
	var prevNodeWrapper *entity.NodeWrapper

	for curNode != nil {
		// init a new node
		newNodeWrapper := &entity.NodeWrapper{
			Data: curNode,
		}

		// setup layout for new node
		newNodeWrapper.Shape = &canvas.Circle{
			StrokeColor: color.White,
			StrokeWidth: 2,
		}
		newNodeWrapper.Interaction = widget.NewButton(fmt.Sprintf("%v", curNode.Value), func() {
		})
		canvasObjs = append(canvasObjs, newNodeWrapper.Shape, newNodeWrapper.Interaction)

		// attach new node to old node
		newNodeWrapper.Prev = prevNodeWrapper

		if prevNodeWrapper == nil {
			// initialize the root
			lay.component.Root = newNodeWrapper
		} else {
			// attach previous node with new node
			prevNodeWrapper.Next = newNodeWrapper

			// add connecting line from the second node
			line := canvas.Line{
				StrokeColor: theme.ForegroundColor(),
				StrokeWidth: 2,
			}
			lay.component.Connections = append(lay.component.Connections, &line)
			canvasObjs = append(canvasObjs, &line)
		}

		// traverse to new node
		prevNodeWrapper = newNodeWrapper
		curNode = curNode.Next
	}

	lay.detail = entity.NewNodeInfo()
	content := container.NewWithoutLayout(canvasObjs...)
	content.Layout = lay
	lay.canvas = content
	return content, nil
}

func Load(win fyne.Window, data interface{}) (fyne.CanvasObject, error) {
	lay := &LinkedListLayout{}

	content, err := lay.render(data)
	if err != nil {
		return nil, err
	}
	box := container.NewVBox(
		content,
		layout.NewSpacer(),
		lay.detail.Detail,
		layout.NewSpacer(),
	)
	return box, nil
}
