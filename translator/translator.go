package translator

import (
	"errors"
	"fmt"
	"strings"
)

// TranslateWord translates a single word to gophers' language.
func TranslateWord(word string) (string, error) {
	if len(word) == 0 {
		return "", errors.New("empty word")
	}

	if strings.ContainsAny(word, "’'") {
		return "", errors.New("gophers can not understand shortened versions of words or apostrophes")
	}

	word = strings.ToLower(word)

	//if the word is already translated in history return it directly instead of translating it again.
	wh := Storage.Load(word)
	if wh != "" {
		return fmt.Sprintf("%v", wh), nil
	}

	prefix := "g"
	vowelIndex := strings.Index(word, "xr")
	if vowelIndex == 0 {
		prefix = "ge"
	} else {
		vowelIndex = strings.IndexAny(word, "aeiou")
		if vowelIndex == -1 {
			vowelIndex = strings.Index(word, "y")
		}

		if vowelIndex >= 2 && word[vowelIndex-1:vowelIndex+1] == "qu" {
			vowelIndex++
		}
	}

	if vowelIndex == -1 {
		return "", fmt.Errorf("'%s' has no vowels", word)
	}

	var builder strings.Builder

	if vowelIndex == 0 {
		builder.WriteString(prefix)
	}
	builder.WriteString(word[vowelIndex:len(word)])
	builder.WriteString(word[0:vowelIndex])
	if vowelIndex != 0 {
		builder.WriteString("ogo")
	}

	gopherWord := builder.String()
	Storage.Store(word, gopherWord)

	return gopherWord, nil
}

func punctuation(word string) (string, string) {
	var p string
	if strings.LastIndexAny(word, ",.?!") == len(word)-1 {
		ln := len(word)
		p = word[ln-1 : ln]
		word = word[:ln-1]
	}
	return word, p
}

// TranslateSentence translates a whole sentence in gopher
func TranslateSentence(sentence string) (string, error) {
	english := strings.Split(sentence, " ")
	var gopher []string

	for _, word := range english {
		word, punctuation := punctuation(word)
		translated, err := TranslateWord(word)
		if err != nil {
			return "", err
		}

		gopher = append(gopher, translated+punctuation)
	}

	return strings.Join(gopher, " "), nil
}
