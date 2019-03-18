package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	// 0.0 <= f < 1.0
	fmt.Print((rand.Float64()*5)+5, ",") //7.1885709359349015,
	fmt.Print((rand.Float64()*5)+5) //7.123187485356329
	fmt.Println()

	//设置种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100),",")
	fmt.Println(r1.Intn(100))

	//如果source设置的是相同的种子，那么产生的随机数也是相同的
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Println(r2.Intn(100))


	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))


}
