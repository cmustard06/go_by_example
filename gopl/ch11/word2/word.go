package word2

import "unicode"

//在前一版本的基础上，判断一个字符串是否是回文字符串
//忽略字母大小写，以及非字母字符

func Ispalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
