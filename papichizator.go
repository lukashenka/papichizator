package papichizator

import (
	"strings"
	"time"
	"math/rand"
	"regexp"
	"fmt"
)

type Papichizable interface {
	Papichize(text string) (string)
}

type Papichizator struct {
}

func (p *Papichizator) Papichize(text string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	reg := regexp.MustCompile("([А-я]+)")
	indexes := reg.FindAllStringIndex(text, -1)
	papichSays := make([]string, len(indexes)*2)
	lastEnd := 0
	for k, index := range indexes {
		start := index[0]
		betweenSpecialRange := make([]int, 2)
		betweenSpecialRange[0] = start
		betweenSpecialRange[1] = lastEnd
		end := index[1]
		lastEnd = end
		specialWord := text[betweenSpecialRange[1]:betweenSpecialRange[0]]
		word := text[start:end]
		papichWord := ""
		translate, ok := translate(word)
		if ok == true {
			papichWord = translate

		} else {
			if len(word) > 4 {
				papichWord =ichizator(word)
			} else {
				papichWord = word
			}
			if rand.Intn(5) == 0 {
				papichWord = curse(papichWord)
			} else {
				papichWord = papichWord
			}
		}
		papichSays[k] = papichWord + specialWord

		fmt.Println(index)
	}

	return strings.Join(papichSays, " ")
}

func ichizator(word string) string {
	lastSymb := word[len(word) -2:]
	fmt.Println(lastSymb)
	hardEndings := map[string]bool {
		"б":true,
		"г":true,
		"з":true,
		"к":true,
		"л":true,
		"п":true,
		"ц":true,
		"ч":true,
		"щ":true,
		"в":false,
		"ж":false,
		"д":false,
		"м":false,
		"н":false,
		"р":false,
		"с":false,
		"т":false,
		"ф":false,
		"х":false,
		"ш": false,
	}


	isI, isHard := hardEndings[lastSymb]
	if isHard {
		if isI {
			return word + "ич"
		} else {
			return word + "ыч"
		}
	} else {
		return ichizator(word[:len(word) -2])
	}

	return "что блять за слово"
}
func curse(word string) string {

	curses:= []string{"блять","тварына","уебище","мать сдохла","ныыыаааа"}
	return word + " " + curses[rand.Intn(len(curses))] + " "
}

func translate(word string) (string, bool) {

	translateWords, ok := translateWords[strings.ToLower(word)]
	if ok {

		translate := translateWords[rand.Intn(len(translateWords))]
		return translate, true
	} else {
		return "", false
	}
}

var translateWords = map[string][]string{
	"привет": {"кулити", "хех, здарова", "добры почанток!"},
}
