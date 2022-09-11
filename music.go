package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false

func MediaPayer() {
	go func(msg string) {
		fmt.Println(msg)
		if streamer != nil {
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")
	time.Sleep(time.Second)
	a := app.New()
	w := a.NewWindow("audio  Player")
	w.Resize(fyne.NewSize(400, 400))

	logo := canvas.NewImageFromFile("logo.jpeg")
	logo.FillMode = canvas.ImageFillOriginal

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
	)
	label := widget.NewLabel("Audio Mp3")
	// label.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel("play Mp3")
	// label2.Alignment = fyne.TextAlignCenter

	browser_file := widget.NewButton("Browse...", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	c := container.NewVBox(label, browser_file, label2, toolbar)
	w.SetContent(
		container.NewBorder(logo, nil, nil, nil, c),
	)
	w.Show()
}
