package array

import (
	"encoding/json"
	"fmt"
	"image/color"

	"godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type (
	ArrayLayout struct {
		component *entity.ElementWrapperList
		detail    *canvas.Text
		canvas    fyne.CanvasObject
	}
)

func (lay *ArrayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *ArrayLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// downsize 10 times
	var downRatio float32 = 10
	if size.Width > lay.MinSize(objs).Width*downRatio {
		size = fyne.NewSize(size.Width/downRatio, size.Height)
	}
	if size.Height > lay.MinSize(objs).Height*downRatio {
		size = fyne.NewSize(size.Width, size.Height/downRatio)
	}

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
		eleDetailJson, err := json.Marshal(lay.component.Nodes[i].Data)
		if err != nil {
			panic(err)
		}
		lay.component.Nodes[i].Interaction = widget.NewButton(mainText, func() {
			lay.detail.Text = string(eleDetailJson)
			lay.detail.Refresh()
		})
		canvasObjs = append(canvasObjs, lay.component.Nodes[i].Shape, lay.component.Nodes[i].Interaction)
	}
	lay.detail = &canvas.Text{
		Color:    color.White,
		TextSize: 12,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	content := container.NewWithoutLayout(canvasObjs...)
	content.Layout = lay

	lay.canvas = content
	return content
}

func Load(win fyne.Window, data interface{}) (fyne.CanvasObject, error) {
	eleList, err := entity.NewElementWrapperList(data)
	if err != nil {
		return nil, err
	}

	lay := &ArrayLayout{
		component: eleList,
	}

	content := lay.render()

	box := container.NewVBox(
		content,
		layout.NewSpacer(),
		lay.detail,
		layout.NewSpacer(),
	)
	return box, nil
}
