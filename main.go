package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type dsLayout struct {
	node *canvas.Circle
	text *canvas.Text

	canvas fyne.CanvasObject
}

func (ds *dsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}

func (ds *dsLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	size = fyne.NewSize(diameter, diameter)
	centerPos := fyne.NewPos(size.Width/2, size.Height/2)

	ds.node.Resize(size)
	ds.node.Move(fyne.NewPos(centerPos.X-radius, centerPos.Y-radius))

	ds.text.Move(fyne.NewPos(radius-ds.node.StrokeWidth, radius-ds.node.StrokeWidth-ds.text.TextSize))
}

func (ds *dsLayout) render() *fyne.Container {
	ds.node = &canvas.Circle{
		StrokeColor: color.White,
		StrokeWidth: 5,
	}
	ds.text = &canvas.Text{
		Text:     "1",
		Color:    color.White,
		TextSize: 12,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}

	container := container.NewWithoutLayout(ds.node, ds.text)
	container.Layout = ds

	ds.canvas = container
	return container
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Node Visus")

	ds := &dsLayout{}
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
