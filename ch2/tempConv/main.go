// tempConv Does some temperature conversion
package main

import (
	"fmt"
)

// Celcius temp
type Celcius float64

// Fahrenheit temp
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	ZeroC         Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string { return fmt.Sprintf("%gC", c) }

// CToF Celcius to Fahrenheit conversion
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Fahrenheit conversion
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("32C = %gF\n", CToF(32))

	boilingF := CToF(BoilingC)

	fmt.Printf("BoilingC = %gF\n", boilingF)
	fmt.Println(FToC(boilingF).String())

}
