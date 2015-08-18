package digest

import (
	"github.com/wicast/Golden2Recite/src/formatter"
	"github.com/wicast/Golden2Recite/src/stardict"
	//"fmt"
	"runtime"
	"sync"
)

func CheckWords(dic_path string, words []string) (map[string]string, []string) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	stardict.PrepareIndex(dic_path)
	done := make(chan struct{})
	defer close(done)

	word_channel := walk_words(done, words)

	result_channel := make(chan word_set)
	var wg sync.WaitGroup
	const numDigesters = 10

	wg.Add(numDigesters)

	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, word_channel, result_channel)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(result_channel)
	}()

	checked_words := make(map[string]string)
	unchecked_words := []string{}

	for r := range result_channel {
		//fmt.Println(r.name, r.desc)
		if r.desc == "no such word" {
			//fmt.Println(r.name, r.desc)
			unchecked_words = append(unchecked_words, r.name+"\r\n")
		} else {
			checked_words[r.name] = r.desc
		}
	}

	return checked_words, unchecked_words

}

func walk_words(done <-chan struct{}, words []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, n := range words {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

type word_set struct {
	name string
	desc string
}

func digester(done <-chan struct{}, word_chan <-chan string, out chan<- word_set) {
	for word := range word_chan {
		str := stardict.Check(word)
		//fmt.Println(word)
		new_word_desc := formatter.Form(str)
		//fmt.Println("digest", new_word_desc)
		select {
		case out <- word_set{word, new_word_desc}:
		case <-done:
			return
		}
	}
}
