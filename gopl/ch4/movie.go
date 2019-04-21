package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Moive struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //c成员标签，omitempty表示如果这个成员的值是零值或者为空，则不输出这个成员到json中
	Actors []string
}

var moives = []Moive{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	//编码为json格式
	data, err := json.Marshal(moives)
	if err != nil {
		log.Fatal("Json marshaling faild:%s", err)
	}
	fmt.Printf("%s\n", data)
	//MarshalIndent
	data, err = json.MarshalIndent(moives, "", "	") //格式化输出json， prefix代表每行第一个字符，
	if err != nil {
		log.Fatal("Json marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)
	/*
		output
		[
			{
				"Title": "Casablanca",
				"released": 1942,
				"Actors": [
					"Humphrey Bogart",
					"Ingrid Bergman"
				]
			},
			{
				"Title": "Cool Hand Luke",
				"released": 1967,
				"color": true,
				"Actors": [
					"Paul Newman"
				]
			},
			{
				"Title": "Bullitt",
				"released": 1968,
				"color": true,
				"Actors": [
					"Steve McQueen",
					"Jacqueline Bisset"
				]
			}
		]

	*/

	// Unmarshall
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("JSON unmarshaling failed:%s", err)
	}
	fmt.Println(titles) //[{Casablanca} {Cool Hand Luke} {Bullitt}]

}
