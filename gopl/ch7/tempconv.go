package main

import (
	"flag"
	"fmt"
)

/*
自己实现一个类似与flag.Duration的方法
./tempconv -cel 22C
*/

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

/*
type Value interface {
	String() string
	Set(string) error
}
*/
//首先需要实现Value接口中所有的方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//注册到flag中，以供调用
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	c := CelsiusFlag("cel", Celsius(1.0), "test")
	flag.Parse()
	fmt.Println(*c)

}
