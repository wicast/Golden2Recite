package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wicast/Golden2Recite/src/deduplicator"
	"github.com/wicast/Golden2Recite/src/digest"
	"github.com/wicast/Golden2Recite/src/file/reader"
	"github.com/wicast/Golden2Recite/src/file/writer"
	"github.com/wicast/Golden2Recite/src/formatter"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	app := cli.NewApp()
	app.Name = "golden2recite"
	app.Usage = "Convert the Goldendict history to Reciteword format"
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) {
		if c.Args()[0] != "" {
			words_path := c.Args()[0]
			words_name := filepath.Base(words_path)
			words_path_abs, err := filepath.Abs(words_path)
			words_dir := filepath.Dir(words_path_abs)

			if err != nil {
				panic("Word path is wrong")
			}
			origin_words := reader.ReadDic(words_path_abs)

			words := deduplicator.Deduplicate(origin_words)

			if c.Args()[1] != "" {
				full, errd := filepath.Abs(c.Args()[1])
				if errd != nil {
					panic("Dict path is wrong")
				}
				ext := filepath.Ext(full)
				stardict_path := full[:strings.LastIndex(full, ext)]

				checked_map, unchecked := digest.CheckWords(stardict_path, words)

				checked := formatter.Map2String(checked_map)

				writer.WriteFile(words_dir+"/recitable_"+words_name, checked)
				writer.WriteFile(words_dir+"/unchecked_"+words_name, unchecked)
			} else {

				fmt.Println("Word path is worng")
				return
			}
		} else {
			fmt.Println("Word path is worng")
			return
		}
	}
	app.Run(os.Args)

}
