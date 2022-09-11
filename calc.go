package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

//changing name main -> showCalc
func showCalc(){
// func main(){
	/*
	a := app.New();
	w := a.NewWindow("calculator");
	w.Resize(fyne.NewSize(600, 400));
	*/
	

	input := widget.NewLabel("");
	isHistory := false;
	historyStirng := "";
	history := widget.NewLabel(historyStirng)
	output := "";

	var historyArray []string;

	historyBtn := widget.NewButton("History", func() {
		if isHistory==true {
			historyStirng = "";
		}else{
			for i:=len(historyArray)-1; i>=0; i--{
				historyStirng = historyStirng + historyArray[i];
				historyStirng += "\n";
			}
		}
		isHistory = !isHistory;
		history.SetText(historyStirng);
	});
	backBtn := widget.NewButton("Back",func() {
		if len(output) > 0 {
			output = output[:len(output)-1];
			input.SetText(output);
		}
	});

	clearBtn := widget.NewButton("clear", func() {
		output = "";
		input.SetText(output);
	});
	openBtn := widget.NewButton("(",func() {
		output = output + "(";
		input.SetText(output);
	});
	closeBtn := widget.NewButton(")", func() {
		output = output + ")";
		input.SetText(output);
	});
	devideBtn := widget.NewButton("/",func() {
		output = output + "/";
		input.SetText(output);
	});

	sevenBtn := widget.NewButton("7", func() {
		output = output + "7";
		input.SetText(output);
	});
	eightBtn := widget.NewButton("8",func() {
		output = output + "8";
		input.SetText(output);
	});
	nineBtn := widget.NewButton("9", func() {
		output = output + "9";
		input.SetText(output);
	});
	multyplyBtn := widget.NewButton("*",func() {
		output = output + "*";
		input.SetText(output);
	});

	fourBtn := widget.NewButton("4", func() {
		output = output + "4";
		input.SetText(output);
	});
	fiveBtn := widget.NewButton("5",func() {
		output = output + "5";
		input.SetText(output);
	});
	sixBtn := widget.NewButton("6", func() {
		output = output + "6";
		input.SetText(output);
	});
	minusBtn := widget.NewButton("-",func() {
		output = output + "-";
		input.SetText(output);
	});

	oneBtn := widget.NewButton("1", func() {
		output = output + "1";
		input.SetText(output);
	});
	twoBtn := widget.NewButton("2",func() {
		output = output + "2";
		input.SetText(output);
	});
	threeBtn := widget.NewButton("3", func() {
		output = output + "3";
		input.SetText(output);
	});
	plussBtn := widget.NewButton("+",func() {
		output = output + "+";
		input.SetText(output);
	});

	zeroBtn := widget.NewButton("0", func() {
		output = output + "0";
		input.SetText(output);
	});
	dotBtn := widget.NewButton(".",func() {
		output = output + ".";
		input.SetText(output);
	});
	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output);
		if err == nil{
			result, err := expression.Evaluate(nil);
			if err == nil{
				ans := strconv.FormatFloat(result.(float64),'f',-1,64);
				strToAppend := output+" = "+ans;
				historyArray = append(historyArray, strToAppend);
				output = ans;
			}else{
				output = "error";
			}
		}else{
			output = "error";
		}
		input.SetText(output);
		fmt.Println(historyArray);
	});
	

	equalBtn.Importance = widget.HighImportance;
	

	
	calcContainer := container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					historyBtn,
					backBtn,
				),
				container.NewGridWithColumns(4,
					clearBtn,
					openBtn,
					closeBtn,
					devideBtn,
				),
				container.NewGridWithColumns(4,
					sevenBtn,
					eightBtn,
					nineBtn,
					multyplyBtn,
				),
				container.NewGridWithColumns(4,
					fourBtn,
					fiveBtn,
					sixBtn,
					minusBtn,
				),
				container.NewGridWithColumns(4,
					oneBtn,
					twoBtn,
					threeBtn,
					plussBtn,
				),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
						zeroBtn,
						dotBtn,
					),
					equalBtn,
				),
			),
	);
	w:=myApp.NewWindow("Calc");
	w.Resize(fyne.NewSize(500,280));
	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calcContainer),
	)
	w.Show();
	

}