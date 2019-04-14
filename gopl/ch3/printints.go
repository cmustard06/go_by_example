package main

import (
	"bytes"
	"fmt"
)

/* 将int数组转换成[1,2,3]这种格式输出的字符串，用到了byte.Buffer

 */
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
