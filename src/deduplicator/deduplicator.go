package deduplicator

import (
	//"fmt"
	"strings"
)

type trie_node struct {
	routeName string
	word      string
	children  []*trie_node
	nextWord  *trie_node
}

var last_word *trie_node

func (this *trie_node) insert_word(word string) *trie_node {
	return this.insert(word, word)
}

func (this *trie_node) insert(word string, rest_word string) *trie_node {

	if word != "" {
		for _, v := range this.children {
			if v.routeName == rest_word[:1] {
				if v.word == word { //Same word
					return last_word
				} else if len(rest_word) == 1 { //Subword
					v.word = word
					last_word.nextWord = v
					return v
				} else { //Continue insert
					return v.insert(word, rest_word[1:])
				}
			}
		}
		return this.add_node(word, rest_word)
	} else {
		return nil
	}
}

func (this *trie_node) add_node(word string, rest_word string) *trie_node {
	if len(rest_word) == 1 { //Last character to insert.
		new_node := &trie_node{rest_word, word, nil, nil}
		last_word.nextWord = new_node
		//last_word = new_node

		this.children = append(this.children, new_node)
		return new_node
	} else if len(rest_word) > 1 { //Continue insert.
		new_node := trie_node{rest_word[:1], "", nil, nil}
		this.children = append(this.children, &new_node)
		n := len(this.children)
		return this.children[n-1].add_node(word, rest_word[1:])
	} else {
		return nil
	}
}

func (root trie_node) walk() []string {
	point := &root
	var dict []string

	for point.children != nil {
		point = point.children[0]
	}
	for point.nextWord != nil {

		dict = append(dict, point.word)
		point = point.nextWord
	}
	dict = append(dict, point.word)
	return dict
}

func Deduplicate(origin_word []string) (deduplicated_word []string) {
	TempTree := trie_node{}

	last_word = &trie_node{}

	for i, _ := range origin_word {
		last_word = TempTree.insert_word(strings.ToLower(origin_word[i]))
	}

	deduplicated_word = TempTree.walk()
	return deduplicated_word
}

//func main() {
//last_word = &trie_node{}
//root := trie_node{}
//	s := []string{"Red", "pink", "Blue", "Green", "yellow", "Yellow", "Pink"}

//s := "ab"
//last_word = root.insert_word(s)
//ss := "acb"
//last_word = root.insert_word(ss)
//sss := "hehe"
//last_word = root.insert_word(sss)

//fmt.Println(root.children[0].children[1].children[0])

//	dictionary := Deduplicate(s)
//	fmt.Println(dictionary)
//}
