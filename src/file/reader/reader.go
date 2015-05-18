package reader

import (
	"bufio"
	"bytes"
	"os"
)

func ReadDic(file_name string) (origin_words []string) {
	fi, err := os.Open(file_name)
	defer fi.Close()	
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)

	//origin_words := make([]string, 0)
	for {
		line, err := r.ReadBytes('\n')
		line = bytes.TrimRight(line, "\r\n")
		line = bytes.TrimLeft(line, "\ufeff")
		origin_words = append(origin_words, string(line))
		if err != nil {
			break
		}
	}
	return origin_words
}
