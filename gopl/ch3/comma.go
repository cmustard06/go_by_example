package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/* 打印从命令行获取到的数字，并将每三个数字用，分开
 */

func comma(v int64) string {
	sign := ""

	// Min int64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		n, _ := strconv.ParseInt(os.Args[i], 10, 64)
		fmt.Printf(" %s\n", comma(n))
	}
}
