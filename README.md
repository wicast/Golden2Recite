# Golden2Recite

Golden2Recite is a tool for convert [Goldendict](http://goldendict.org/) history to [生词本背单词](https://play.google.com/store/apps/details?id=zoz.reciteword) acceptable format.  

~~This project is mainly for golang practising.~~

*Notice that this tool is using stardict format to lookup explanation.*

### Usage
```
NAME:
   golden2recite - Convert the Goldendict history to Reciteword format

USAGE:
   golden2recite [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
The first argument is your new word which is about to recite,the second argument is your startdict files.A stardict-oxford-gb dictionary is already provided as a recomendation.
##### Usage Example
`$ ./Golden2Recite ./new_words_exp.txt ./stardict-oxford-gb-formated/oxford-gb-formated.dict`

This will create two new files with prefix "recitable_" as importable for 生词本背单词 and "unchecked_" which can't find its explaination.

### Dependency
http://github.com/codegangsta/cli —— Cli  
https://github.com/chuangbo/jianbing-dictionary-dns/blob/master/golang/jianbing-dns/stardict/stardict.go —— stardict api(modified)

### Build
This project use [gvp](https://github.com/pote/gvp) and [gpm](https://github.com/pote/gpm) as dependency manager.

Golang Compiler is required.

1. `source gvp in` —— Set local GOPATH
2. `gpm install` —— Install Dependency into GOPATH
3. `go install` —— Build project,the executable can be found in "bin" folder.

If you need cross complier for other platform check this [article](http://spf13.com/post/cross-compiling-go/) out.