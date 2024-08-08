package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
	"os"
)



func main(){
	reader := bufio.NewReader(os.Stdin)
	sentenceByte, _, _:= reader.ReadLine()
	sentence := string(sentenceByte)
	strings.TrimSpace(sentence)
	sentence = trimPunctuation(sentence)
	s := strings.ToLower(s)
	count:=  Counter(sentence)
	fmt.Println(count)
	
}

func Counter(s string) map[string]int {
	count := make(map[string]int)
	result := trimPunctuation(s)
	parts := strings.Split(result, " ")

	for _, word := range parts{
		count[word] ++
	}

	return count

}

func trimPunctuation(s string) string {
	var result strings.Builder
	for _, char  := range s{
		if !unicode.IsPunct(char){
			result.WriteRune(char)
		}
	}
	return result.String()
}

func chackPalidrome(s string) bool{
	left, right := 0, len(s)
	for left < right{
		if s[left] == s[right]{
			left += 1
			right -= 1
		}
		return false
	}
	return true
}
