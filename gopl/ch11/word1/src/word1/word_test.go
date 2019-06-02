package word1

import "testing"

//检查isPalindrome是否针对单个输入参数给出正确的结果
//go test在没有指定包参数的情况下，默认会以当前目录所在的包为参数

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated")=false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPanlindrome("kayak")=false`)
	}

}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`=true`)
	}
}
