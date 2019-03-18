package main

import "fmt"

import "time"

func main() {
	i := 2
	fmt.Print("Write ", i, " as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	var weekday = time.Now().Weekday()
	fmt.Println(weekday)
	// 常用的格式
	switch weekday{
	case time.Saturday,time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It's weekday")
	}

	//可以实现if/else逻辑
	t := time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}){
		switch t:= i.(type){   //该使用形式只能在switch语句块中使用
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an Int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("jj")

}
