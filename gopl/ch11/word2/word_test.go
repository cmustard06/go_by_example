package word2

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	//这种基于表的测试方式在go里面很常见
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := Ispalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//随机测试
//返回一个回文字符串，它的长度和内容都是随机数生成的
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r //随机字符最大\u0999
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	//初始化一个伪随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !Ispalindrome(p) {
			t.Errorf("%q==false", p)
		}
	}
}

//性能测试
//go test --bench=.
func BenchmarkIspalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ispalindrome("A man, a plan, a canal: Panama")
	}
}
