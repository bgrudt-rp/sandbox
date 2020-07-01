package main

import "fmt"

type methodLoopExample struct {
	tempC    float64 `header:"Celsius"`
	tempF    float64 `header:"Fahrenheit"`
	thoughts string  `header:"Ask Brandon"`
}

type methodLoopSlice []methodLoopExample

func testMethodLoop() {
	var outputTemps methodLoopSlice
	inputTemps := []float64{25, 21.1, 12.5, 33.3, 0.2, -27.7}

	for _, temp := range inputTemps {
		outputTemps.analyze(temp)
	}

	fmt.Println(outputTemps)

}

func (l *methodLoopSlice) analyze(c float64) {
	var o methodLoopExample

	o.tempC = c
	o.tempF = c*1.8 + 32

	switch {
	case o.tempF < 5:
		o.thoughts = "Okay, yes it is actually cold."
	case o.tempF >= 5 && o.tempF < 32:
		o.thoughts = "Probably should wear a coat."
	case o.tempF >= 32 && o.tempF < 60:
		o.thoughts = "Hoodie and shorts."
	case o.tempF >= 60 && o.tempF < 72:
		o.thoughts = "This is football weather."
	case o.tempF >= 72 && o.tempF < 77:
		o.thoughts = "Nice if there's a breeze."
	case o.tempF >= 77 && o.tempF < 90:
		o.thoughts = "It's just too hot out."
	case o.tempF >= 90:
		o.thoughts = "Literal hell on earth."
	}

	*l = append(*l, o)

}
