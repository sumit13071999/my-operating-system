package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//for calling in main os change its name main -> showWeatherApp
func showWeatherApp(w fyne.Window){
// func main(){
	/*
	a := app.New();
	w := a.NewWindow("wheather app");
	w.Resize(fyne.NewSize(600,400));
	*/
	

	//api work
	res,err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=haridwar&APPID=a4554fcd0c4725b34d27438684db2d4e");
	if err != nil {
		fmt.Println(err);
	}
	defer res.Body.Close();
	body , err := ioutil.ReadAll(res.Body);
	if err!=nil{
		fmt.Println(err);
	}

	weather , err := UnmarshalWelcome(body);
	if err!=nil{
		fmt.Println(err);
	}
	//api work finish

	img := canvas.NewImageFromFile("weather.png");
	img.FillMode = canvas.ImageFillOriginal;

	label1 := canvas.NewText("weather deatails",color.White);
	label1.TextStyle = fyne.TextStyle{Bold: true};
	
	label2 := canvas.NewText(fmt.Sprintf("Contry %s", weather.Sys.Country),color.Black);
	label3 := canvas.NewText(fmt.Sprintf("windSpeed %.2f", weather.Wind.Speed),color.Black);
	label4 := canvas.NewText(fmt.Sprintf("Temprature %.2f", weather.Main.Temp),color.Black);
	label5 := canvas.NewText(fmt.Sprint("Humidity ", weather.Main.Humidity),color.Black);
	label6 := canvas.NewText(fmt.Sprint("Pressure",weather.Main.Temp),color.Black);
/*
	w.SetContent(
		container.NewVBox(
			label1,
			img,
			label2,
			label3,
			label4,
			label5,
			label6,
		),
	)
	w.ShowAndRun();
	*/
	
	weatherContainer := container.NewVBox(
		label1,
		img,
		label2,
		label3,
		label4,
		label5,
		label6,
		container.NewGridWithColumns(1,
		),
	);

	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,weatherContainer),);
	w.Show();
	

}





// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()



func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}