package translator

import (
	"errors"
	"fmt"
	"golang.org/x/text/unicode/rangetable"
	"strings"
	"unicode"
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

	words := strings.Split(word, " ")

	var res []string
	for _, str := range words {
		encoded, _ := encode(str)
		res = append(res, encoded)
	}

	gopherWord := strings.Join(res, " ")
	Storage.Store(word, gopherWord)

	return gopherWord, nil
}

func encode(in string) (string, error) {
	in = strings.ToLower(in)
	vowel := rangetable.New('a', 'e', 'i', 'o', 'u', 'y')

	var vowelidx int

	for k, v := range in {
		if unicode.Is(vowel, v) {
			// weird exceptions for 'u'
			if v == 'u' && unicode.Is(vowel, rune(in[k+1])) {
				vowelidx++
			}

			if in[:k] == "xr" {
				return "ge" + in[vowelidx:] + in[:vowelidx], nil
			}

			if in[:k] == "qu" {
				return in[vowelidx:] + in[:vowelidx] + "quogo", nil
			}

			vowelidx += k
			break
		}
	}

	if vowelidx == -1 {
		return "", fmt.Errorf("'%s' has no vowels", in)
	}

	// add weird exeption for xray
	if vowelidx == 0 {
		return "g" + in, nil
	}

	suffix := ""
	if vowelidx != 0 {
		suffix = "ogo"
	}

	return in[vowelidx:] + in[:vowelidx] + suffix, nil
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
