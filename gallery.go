package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/internal/widget"
	// "fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/internal/widget"
)

func showGalleryApp(w fyne.Window){
// func main(){
	/*
	a := app.New();
	w := a.NewWindow("Imgage Viewer");
	w.Resize(fyne.NewSize(800,600));
	*/
	
    root_src := "C:\\Users\\Sumit Kumar\\OneDrive\\Pictures";
	files, err := ioutil.ReadDir(root_src);
	if err != nil{
		log.Fatal(err);
	}
	var picsArray []string;
	tabs := container.NewAppTabs();
	for _, file := range files{
		fmt.Println(file.Name(), file.IsDir());
		if !file.IsDir(){
			extension := strings.Split(file.Name(), ".");
			if extension[1] == "png" || extension[1] == "jpg" || extension[1] == "jpeg"{
				picsArray = append(picsArray,);
				image := canvas.NewImageFromFile(root_src+"/"+file.Name());
				image.FillMode = canvas.ImageFillContain;
				tabs.Append(container.NewTabItem(file.Name(),image));
			}
		}
	}
	

	tabs.SetTabLocation(container.TabLocationLeading);

	// w.SetContent(tabs);
	// w.ShowAndRun();
	galleryContainer := tabs;

	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,galleryContainer),);
	w.Show();
}










