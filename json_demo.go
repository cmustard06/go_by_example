package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int  `json:"page"`
	Fruits []string `json:"fruits"` //自定义json编码后的key名称为fruits
}

func main(){
	bolB,_:= json.Marshal(true)
	fmt.Println(string(bolB))

	intB,_:= json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD) //["apple","peach","pear"]
	fmt.Println(string(slcB))

	res1D := &response1{
		Page:1,
		Fruits:[]string{"apple","peach","pear"},
	}
	res1B,_ := json.Marshal(res1D)
	fmt.Println(string(res1B))  //{"Page":1,"Fruits":["apple","peach","pear"]}

	mapD := map[string]int{"apple":5,"lettuce":7}
	mapB,_ := json.Marshal(mapD)
	fmt.Println(string(mapB))  //{"apple":5,"lettuce":7}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) //{"page":1,"fruits":["apple","peach","pear"]}

	//现在将json数据解码为go识别的类型

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}  //注意这里的键值必须为接口类型

	if err:= json.Unmarshal(byt, &dat);err!=nil{
		panic(err)
	}
	fmt.Println(dat)  //map[num:6.13 strs:[a b]]
	//num := dat["num"].(float64)
	//fmt.Printf("%T, %f\n", num,num) //float64, 6.130000
	fmt.Printf("%T\n", dat["num"])  //float64a

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Printf("%s\n",str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res) //{1 [apple peach]}

	fmt.Println(res.Fruits[0])
	//还可以将JSON编码直接流式传输到os.Writers，如os.Stdout甚至HTTP响应体
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d) //{"apple":5,"lettuce":7}


}
