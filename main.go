package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("pepos")
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget
var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("img1.jpeg")
	btn1 = widget.NewButtonWithIcon("Weather", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})
	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})
	btn3 = widget.NewButtonWithIcon("Gallery", theme.StorageIcon(), func() {
		showGalleryApp(myWindow)
	})
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor(myWindow)
	})
	btn5 = widget.NewButtonWithIcon("Music Palyer", theme.MediaPlayIcon(), func() {
		MediaPayer()
	})
	btn6 := widget.NewButton("News app", func() {
		NewsApp()
	})
	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(
	7,
		DeskBtn,
		btn1,
		btn2,
		btn3,
		btn4,
		btn5,
		btn6,
	))

	myWindow.Resize(fyne.NewSize(1500, 780))
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	myWindow.ShowAndRun()

}
