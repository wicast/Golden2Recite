package writer

import (
	"os"
)

func WriteFile(file_name string, text []string) {
	fi, err := os.Create(file_name)
	defer fi.Close()
	if err != nil {
		panic(err)
	}

	for _, v := range text {
		fi.WriteString(v)
	}
}
