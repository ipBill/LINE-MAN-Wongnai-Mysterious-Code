package main

import (
	b64 "encoding/base64"
	"fmt"
	"unicode"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sDec, _ := b64.StdEncoding.DecodeString(secret)
	var whatIsIt = string(sDec)
	fmt.Println(backwardString(whatIsIt))
}

func backwardString(input string) (result string) {
	for _, v := range input {
		if unicode.IsLetter(v) {
			result = string(v) + result
		} else {
			result = " " + result
		}
	}
	return
}
