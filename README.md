# IsAltGo

定義 IsAlt 函式判斷輸入 string 是否為子母音交錯

## 我的解法

```golang
import "regexp"

func IsAlt(input string) bool {
	nonAltPattern := "[aeiou]{2}|[^aeiou]{2}"
	match, _ := regexp.MatchString(nonAltPattern, input)
	return !match
}

```