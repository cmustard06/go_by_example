package main

import (
	"fmt"
	"os"
)

type point struct {
	x,y int
}

func main(){
	p := point{1,2}
	fmt.Printf("%v\n",p) // {1 2}
	fmt.Printf("%+v\n",p) //{x:1 y:2}
	fmt.Printf("%#v\n",p) //main.point{x:1, y:2}
	fmt.Printf("%T\n",p) //main.point

	fmt.Printf("%t\n", true)  //true
	fmt.Printf("%b\n",14) //1110
	fmt.Printf("%c\n",33) //!
	fmt.Printf("%x\n", 456) //1c8
	fmt.Printf("%e\n", 123400000.0) //1.234000e+08
	fmt.Printf("%E\n", 123400000.0) //1.234000E+08
	fmt.Printf("%q\n", "\"string\"")//"\"string\""
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p) //0xc000054080



	fmt.Printf("|%6d|%6d|\n",12,234) //|    12|   234|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //|  1.20|  3.45|
	fmt.Printf("|%6s|%6s|\n", "foo", "b") //|   foo|     b|

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //|foo   |b     |

	fmt.Fprintf(os.Stdout, "an %s\n", "error") //an error















}
