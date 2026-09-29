package main

import "fmt"

// Even though both have the same underlying type, float64, they are not the same type, so they cannot be compared or combined
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// The declaration below, in which the Celsius parameter c appears before the function name,
// associates with the Celsius type a method named String that returns c’s numeric value followed by °C:
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	t1c := Celsius(30)
	fmt.Printf("%g°C => %g°F\n", t1c, CToF(t1c))

	fmt.Printf("%v\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%v\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
}
