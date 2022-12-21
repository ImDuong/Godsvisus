package array

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"reflect"

	"github.com/ImDuong/godsvisus/internal/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type (
	ArrayLayout struct {
		components []*entity.ElementWrapperList
		detail     *entity.NodeInfo
		canvas     fyne.CanvasObject
		isArranged bool
	}
)

func (lay *ArrayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (lay *ArrayLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	// avoid duplication of layout
	if lay.isArranged {
		return
	}

	// downsize 10 times
	var downRatio float32 = 10
	if size.Width > lay.MinSize(objs).Width*downRatio {
		size = fyne.NewSize(size.Width/downRatio, size.Height)
	}
	if size.Height > lay.MinSize(objs).Height*downRatio {
		size = fyne.NewSize(size.Width, size.Height/downRatio)
	}

	// hardcode hoziontal layout for an array
	distanceBtwEles := fyne.NewSize(size.Width, 0)
	distanceBtwLists := fyne.NewSize(0, size.Height*2)
	for compIdx := range lay.components {
		for i, node := range lay.components[compIdx].Nodes {
			node.Resize(size)

			// separate every element in a list by a distance of node.Shape.StrokeWidth
			node.Move(fyne.NewPos(
				float32(i)*(distanceBtwEles.Width+node.Shape.StrokeWidth*2)+float32(compIdx)*distanceBtwLists.Width,
				float32(i)*distanceBtwEles.Height+float32(compIdx)*distanceBtwLists.Height,
			))
		}
	}
	lay.isArranged = true
}

func (lay *ArrayLayout) render(data interface{}) (*fyne.Container, error) {
	// validate input
	dataKind := reflect.TypeOf(data).Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Array {
		return nil, errors.New("input data is not a list")
	}

	// parse input
	arrayList := reflect.ValueOf(data)

	// init canvas objs
	canvasObjs := []fyne.CanvasObject{}

	// init nodes
	lay.components = make([]*entity.ElementWrapperList, arrayList.Len())

	for compIdx := 0; compIdx < arrayList.Len(); compIdx++ {
		arrayComponent := arrayList.Index(compIdx)
		eleKind := reflect.TypeOf(arrayComponent.Interface()).Kind()
		if eleKind != reflect.Slice && eleKind != reflect.Array {
			return nil, errors.New("element of input data is not a list")
		}

		lay.components[compIdx] = &entity.ElementWrapperList{
			Nodes: make([]*entity.ElementWrapper, arrayComponent.Len()),
		}

		for i := 0; i < arrayComponent.Len(); i++ {
			// setup node data
			lay.components[compIdx].Nodes[i] = &entity.ElementWrapper{
				Data: arrayComponent.Index(i).Interface(),
			}
			if arrayComponent.Index(i).CanAddr() {
				lay.components[compIdx].Nodes[i].DataAddr = arrayComponent.Index(i).UnsafeAddr()
			}

			// setup node layout
			lay.components[compIdx].Nodes[i].Shape = &canvas.Rectangle{
				StrokeColor: color.White,
				StrokeWidth: 2,
			}
			mainText := fmt.Sprintf("%v", lay.components[compIdx].Nodes[i].Data)
			eleAddr := fmt.Sprintf("0x%x", lay.components[compIdx].Nodes[i].DataAddr)
			eleDetailJson, err := json.MarshalIndent(lay.components[compIdx].Nodes[i].Data, "", "\t")
			if err != nil {
				return nil, err
			}
			lay.components[compIdx].Nodes[i].Interaction = widget.NewButton(mainText, func() {
				lay.detail.SetInfo(eleAddr, string(eleDetailJson))
				lay.detail.Detail.Refresh()
			})

			canvasObjs = append(canvasObjs, lay.components[compIdx].Nodes[i].Shape, lay.components[compIdx].Nodes[i].Interaction)
		}
	}

	// register canvas objs with the container
	content := container.NewWithoutLayout(canvasObjs...)
	content.Layout = lay
	lay.canvas = content
	return content, nil
}

func Load(win fyne.Window, data interface{}) (fyne.CanvasObject, fyne.Layout, error) {
	lay := &ArrayLayout{}
	lay.detail = entity.NewNodeInfo()
	content, err := lay.render(data)
	if err != nil {
		return nil, nil, err
	}

	box := container.NewVBox(
		content,
		layout.NewSpacer(),
		lay.detail.Detail,
		layout.NewSpacer(),
	)
	return box, lay, nil
}
