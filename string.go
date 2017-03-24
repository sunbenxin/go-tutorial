package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var a = []string{
	"孙本新",
}

func main() {
	str := "中文文文文"
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$")
	fmt.Println(hzRegexp.MatchString(str))
	fmt.Println(IsChineseChar(str))
	fmt.Println(utf8.RuneCountInString(a[0]))
}
