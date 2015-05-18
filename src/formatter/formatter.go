package formatter

import (
	//"fmt"
	"strings"
)

func words2slice(decription string) []string {
	s := strings.Split(decription, "\n")
	s[0] = strings.Trim(s[0], "/*[]\r")
	if n := strings.Index(s[0], ";"); n != -1 {
		s[0] = s[0][:n]
	}
	return s
}

func Form(str string) (word_desc string) {
	word_desc = ""
	if str == "no such word" {
		//fmt.Println(str)
		return str
	} else {
		s := words2slice(str)
		for i, k := range s {
			if i != 0 {
				word_desc = word_desc + "#" + k + "\r\n"
			} else {
				word_desc = word_desc + "&" + k + "\r\n"
			}
		}
		word_desc = word_desc + "$1\n\r"
		return word_desc
	}
}

func Map2String(dict_map map[string]string) []string {
	s := []string{}
	for k, v := range dict_map {
		s = append(s, "+"+k+"\r\n"+v)
	}
	return s
}
