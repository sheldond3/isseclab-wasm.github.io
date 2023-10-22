package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

// random num gen
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// defines random color
func color() string {
	letters := "0123456789ABCDEF"
	color := make([]byte, 6)
	//creates new random number gen with random seed

	for i := 0; i < 6; i++ {
		color[i] += letters[int(math.Floor(r.Float64()*16))]
	}
	return "#" + string(color)
} //end color

func changeColAndRotate(this js.Value, args []js.Value) interface{} {

	element := js.Global().Get("document").Call("getElementById", args[0].String()).Get("style")
	//change color
	divColor := color()
	element.Set("background-color", divColor)

	//rotate element
	rotVal := int(math.Floor(r.Float64() * 360))
	element.Set("transform", "rotate("+fmt.Sprintf("%d", rotVal)+"deg)")

	return true
} //end changeColAndRotate

func resetAll(this js.Value, args []js.Value) interface{} {

	element := args[0].Get("style")
	divColor := "white"
	element.Set("background-color", divColor)

	element.Set("transform", "rotate(0deg)")

	return true
} //end resetAll

func display(this js.Value, args []js.Value) interface{} {

	word := args[0].String()
	var divString string
	var elementID string

	for i := 0; i < len(word); i++ {

		divString += "<div id=\""

		//if balnk space add to string as such
		if string(word[i]) == " " {
			elementID = "space" + strconv.Itoa(i)
			divString += elementID + "\" class = \"displayBox\"" +
				"onclick=\"CaR(this.id)" + "\">" +
				"&nbsp;" + "</div>"
		} else {
			elementID = string(word[i]) + strconv.Itoa(i)
			divString += elementID + "\" class = \"displayBox\" " +
				"onclick=\"CaR(this.id)" + "\">" +
				string(word[i]) + "</div>"
		}

	}

	return divString
} //end display

func main() {

	ch := make(chan struct{}, 0)
	fmt.Printf("Hello Web Assembly from Go!\n")

	js.Global().Set("goDisplay", js.FuncOf(display))
	js.Global().Set("goCaR", js.FuncOf(changeColAndRotate))
	js.Global().Set("goReset", js.FuncOf(resetAll))
	<-ch

} //end main
