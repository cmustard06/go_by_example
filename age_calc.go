package main

import (
	"fmt"
	"time"
)

func GetTimeFromStrDate(date string) (year, month, day int) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date) //输入的日期，格式需要满足2006-01-02
	if err != nil {
		fmt.Println("出生日期错误")
		return 0, 0, 0
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}

func GetZodiac(year int) (zodiac string) {
	//获取生肖
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}

	if x == 0 {
		zodiac = "牛"
	}

	if x == 11 || x == -1 {
		zodiac = "虎"
	}

	if x == 10 || x == -2 {
		zodiac = "兔"
	}

	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func main() {
	y, m, d := GetTimeFromStrDate("1567-09-11")
	fmt.Printf("%d-%d-%d\n", y, m, d)
	fmt.Println(GetZodiac(y))
}
