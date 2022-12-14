package array

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"reflect"

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
		detail    *entity.NodeInfo
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

func (lay *ArrayLayout) render(data interface{}) (*fyne.Container, error) {
	// validate input
	dataKind := reflect.TypeOf(data).Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return nil, errors.New("input data is not a list")
	}

	// parse input
	dataList := reflect.ValueOf(data)

	// init nodes & canvas objs
	lay.component = &entity.ElementWrapperList{
		Nodes: make([]*entity.ElementWrapper, dataList.Len()),
	}
	canvasObjs := []fyne.CanvasObject{}

	for i := 0; i < dataList.Len(); i++ {
		// setup node data
		lay.component.Nodes[i] = &entity.ElementWrapper{
			Data: dataList.Index(i).Interface(),
		}
		if dataList.Index(i).CanAddr() {
			lay.component.Nodes[i].DataAddr = dataList.Index(i).UnsafeAddr()
		}

		// setup node layout
		lay.component.Nodes[i].Shape = &canvas.Rectangle{
			StrokeColor: color.White,
			StrokeWidth: 2,
		}
		mainText := fmt.Sprintf("%v", lay.component.Nodes[i].Data)
		eleAddr := fmt.Sprintf("0x%x", lay.component.Nodes[i].DataAddr)
		eleDetailJson, err := json.MarshalIndent(lay.component.Nodes[i].Data, "", "\t")
		if err != nil {
			return nil, err
		}
		lay.component.Nodes[i].Interaction = widget.NewButton(mainText, func() {
			lay.detail.SetInfo(eleAddr, string(eleDetailJson))
			lay.detail.Detail.Refresh()
		})

		// register canvas objs
		canvasObjs = append(canvasObjs, lay.component.Nodes[i].Shape, lay.component.Nodes[i].Interaction)
	}
	lay.detail = entity.NewNodeInfo()
	content := container.NewWithoutLayout(canvasObjs...)
	content.Layout = lay
	lay.canvas = content
	return content, nil
}

func Load(win fyne.Window, data interface{}) (fyne.CanvasObject, error) {
	lay := &ArrayLayout{}
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
