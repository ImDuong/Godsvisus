package array

import (
	"fmt"
	"image/color"

	"godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

	// hardcode hoziontal layout for an array
	distance := fyne.NewSize(size.Width, 0)
	for i, node := range lay.component.Nodes {
		node.Resize(size)

		// separate every element by a distance of node.Shape.StrokeWidth
		node.Move(fyne.NewPos(
			float32(i)*(distance.Width+node.Shape.StrokeWidth*2),
			float32(i)*distance.Height,
		))
	}
}

func (lay *ArrayLayout) render() *fyne.Container {
	canvasObjs := []fyne.CanvasObject{}
	for i := range lay.component.Nodes {
		lay.component.Nodes[i].Shape = &canvas.Rectangle{
			StrokeColor: color.White,
			StrokeWidth: 2,
		}
		mainText := fmt.Sprintf("%v", lay.component.Nodes[i].Data.Value)
		lay.component.Nodes[i].Interaction = widget.NewButton(mainText, func() {
			fmt.Println(mainText)
		})
		canvasObjs = append(canvasObjs, lay.component.Nodes[i].Shape, lay.component.Nodes[i].Interaction)
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
