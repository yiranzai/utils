package pinyin

import (
	"unicode"

	py "github.com/mozillazg/go-pinyin"
)

//GeneratePinyin 生成拼音
func GeneratePinyin(name string) (string, string) {
	var pinyin, firstLetter string
	a := py.NewArgs()
	a.Fallback = func(r rune, a py.Args) []string {
		return []string{string(r)}
	}
	a.Separator = ""
	pinyin = py.Slug(name, a)

	a.Style = py.FirstLetter
	name = resetString(name)
	firstLetter = py.Slug(name, a)

	return pinyin, firstLetter
}

//resetString 重置拼音
func resetString(name string) string {
	var strings []rune
	var isLetter bool
	isFirstLetter := true
	for _, n := range name {
		isLetter = IsLetter(n)
		if isLetter && isFirstLetter {
			strings = append(strings, n)
			isFirstLetter = false
		}
		if isLetter == false {
			isFirstLetter = true
			if unicode.Is(unicode.Scripts["Han"], n) {
				strings = append(strings, n)
			}
		}
	}
	return string(strings)
}

//IsLetter 是否英文字母
func IsLetter(r rune) bool {
	if r > 64 && r < 91 {
		return true
	}
	if r > 96 && r < 123 {
		return true
	}
	return false
}
