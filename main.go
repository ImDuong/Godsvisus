package main

import (
	"encoding/json"
	"errors"
	"image/color"
	"math/rand"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ref: https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1

const MEME_API = "https://api.imgflip.com/get_memes"

var myClient = &http.Client{Timeout: 10 * time.Second}

type (
	Meme struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	MemeData struct {
		Memes []Meme `json:"memes"`
	}

	MemeResponse struct {
		IsSuccess bool     `json:"success"`
		Data      MemeData `json:"data"`
	}
)

func getAMeme() (*Meme, error) {
	r, err := myClient.Get(MEME_API)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var memeResponse MemeResponse
	err = json.NewDecoder(r.Body).Decode(&memeResponse)
	if err != nil {
		return nil, err
	}
	if memeResponse.IsSuccess != true {
		return nil, errors.New("can not get meme")
	}
	// get a random index
	rand.Seed(time.Now().UnixNano())
	randIdx := 0 + rand.Intn(len(memeResponse.Data.Memes)-1-0+1)
	return &memeResponse.Data.Memes[randIdx], nil
}

func main() {
	visusApp := app.New()
	visusWindow := visusApp.NewWindow("GODSVISUS")

	// file menu
	fileMenu := fyne.NewMenu("File", fyne.NewMenuItem(
		"Quit", func() {
			visusApp.Quit()
		},
	))

	mainMenu := fyne.NewMainMenu(
		fileMenu,
	)

	visusWindow.SetMainMenu(mainMenu)

	// define a welcome text
	text := canvas.NewText("What do you expect?", color.White)
	text.Alignment = fyne.TextAlignCenter

	// get a meme
	meme, err := getAMeme()
	if err != nil {
		panic(err)
	}

	// define a resource holder
	resource, err := fyne.LoadResourceFromURLString(meme.URL)
	if err != nil {
		panic(err)
	}
	resourceHolder := canvas.NewImageFromResource(resource)
	resourceHolder.SetMinSize(fyne.Size{
		Width:  700,
		Height: 500,
	})

	// define a button to get a new meme
	getMemeBtn := widget.NewButton("Pandora", func() {
		meme, err = getAMeme()
		if err != nil {
			panic(err)
		}
		resource, err := fyne.LoadResourceFromURLString(meme.URL)
		if err != nil {
			panic(err)
		}
		resourceHolder.Resource = resource
		resourceHolder.Refresh()
	})
	getMemeBtn.Importance = widget.HighImportance

	// display a vertical box to contain text, image, button
	box := container.NewVBox(
		text,
		resourceHolder,
		getMemeBtn,
	)

	// display content
	visusWindow.SetContent(box)

	// attach Escape key to exit app
	visusWindow.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyEscape {
			visusApp.Quit()
		}
	})

	visusWindow.ShowAndRun()
}
