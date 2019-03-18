package main

import (
	"flag"
	"fmt"
)

func main()  {
	// value代表的默认值
	wordPtr := flag.String("word","foo","a string")
	numPtr := flag.Int("numb",42,"an int")
	//对于布尔型而言，默认值如果为false，-fork就是true
	boolPtr := flag.Bool("fork",false,"a bool")

	var svar string
	flag.StringVar(&svar,"svar", "bar","a string var")

	flag.Parse()

	fmt.Println("word:",*wordPtr)
	fmt.Println("numb", *numPtr)
	fmt.Println("fork",*boolPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail:",flag.Args())

}
