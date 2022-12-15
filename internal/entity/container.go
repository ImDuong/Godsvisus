package entity

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	NodeInfo struct {
		Detail  *fyne.Container
		address *canvas.Text
		info    *widget.Label
	}
)

func NewNodeInfo() *NodeInfo {
	ni := NodeInfo{}
	nodeAddrLabel := canvas.Text{
		Text:     "Address: ",
		Color:    color.White,
		TextSize: 12,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	ni.address = &canvas.Text{
		Color:    color.White,
		TextSize: 12,
	}
	nodeInfoLabel := canvas.Text{
		Text:     "Info: ",
		Color:    color.White,
		TextSize: 12,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	ni.info = &widget.Label{}
	ni.Detail = container.NewVBox(
		container.NewHBox(
			&nodeAddrLabel,
			ni.address,
		),
		container.NewHBox(
			&nodeInfoLabel,
			ni.info,
		),
	)
	return &ni
}

func (ni *NodeInfo) SetInfo(address, info string) {
	ni.address.Text = address
	ni.info.Text = info
}
