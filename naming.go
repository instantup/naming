// Package naming provides functions for name splitting and converting to
// common naming conventions.
package naming

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// A Naming function formats names according to a naming convention.
type Naming func(name string) string

var (
	// Flat converts names to flatcase.
	Flat = CreateNaming(strings.ToLower, strings.ToLower, "")
	// Upper converts names to UPPERCASE.
	Upper = CreateNaming(strings.ToUpper, strings.ToUpper, "")
	// Mixed converts names to mixedCase
	Mixed = CreateNaming(strings.ToLower, ToTitle, "")
	// UpperMixed converts names to MixedCase
	UpperMixed = CreateNaming(ToTitle, ToTitle, "")
	// Snake converts names to snake_case
	Snake = CreateNaming(strings.ToLower, strings.ToLower, "_")
	// UpperSnake converts names to SNAKE_CASE
	UpperSnake = CreateNaming(strings.ToUpper, strings.ToUpper, "_")
	// Kebab converts names to kebab-case
	Kebab = CreateNaming(strings.ToLower, strings.ToLower, "-")
)

// CreateNaming creates a Naming function. The function mapFirst is called to
// format the first word in a name. The function mapRest is called to format
// the second and following words. The separator string sep is placed between
// words in the resulting name.
func CreateNaming(mapFirst func(string) string, mapRest func(string) string, sep string) Naming {
	return func(name string) string {
		words := SplitName(name)
		var b strings.Builder
		b.Grow(len(name))
		b.WriteString(mapFirst(words[0]))
		for _, word := range words[1:] {
			b.WriteString(sep)
			b.WriteString(mapRest(word))
		}
		return b.String()
	}
}

// ToTitle returns word with the first Unicode character mapped to title case.
// Letters following the first character are mapped to lower case.
func ToTitle(word string) string {
	firstRune, firstRuneSize := utf8.DecodeRuneInString(word)
	var b strings.Builder
	b.Grow(len(word))
	b.WriteRune(unicode.ToTitle(firstRune))
	b.WriteString(strings.ToLower(word[firstRuneSize:]))
	return b.String()
}

// SplitName splits name into distinct words. Unicode letters and digits are
// considered to be word characters. All other characters are thrown out.
//
// Words are split according to the following rules:
//   - Word characters separated by non-word characters.
//     e.g. word_word => word word
//
//   - A lower case or digit character followed by an upper case character.
//     e.g. wordWord => word Word
//
//   - When two or more upper case characters are followed by a lower case
//     character, a split is made before the last upper case character.
//     e.g. WORDWord => WORD Word
func SplitName(name string) []string {
	if name == "" {
		return []string{""}
	}

	if utf8.RuneCountInString(name) == 1 {
		firstRune, _ := utf8.DecodeRuneInString(name)
		if unicode.IsLetter(firstRune) || unicode.IsDigit(firstRune) {
			return []string{name}
		}
		return []string{""}
	}

	words := new(splitState).splitName(name)
	if words == nil {
		return []string{""}
	}
	return words
}

type splitState struct {
	mode    splitMode
	words   []string
	curWord strings.Builder
}

type splitMode int

const (
	boundSplitMode splitMode = iota
	wordStartSplitMode
	wordSplitMode
	acronymSplitMode
)

func (s *splitState) splitName(name string) []string {
	for _, r := range name {
		switch s.mode {
		case boundSplitMode:
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				s.startWord(r)
			}
		case wordStartSplitMode:
			if unicode.IsUpper(r) || unicode.IsLower(r) || unicode.IsDigit(r) {
				s.writeWordRune(r)
			} else {
				s.finishWord(r)
			}
		case wordSplitMode:
			if unicode.IsLower(r) || unicode.IsDigit(r) {
				s.writeWordRune(r)
			} else {
				s.finishWord(r)
			}
		case acronymSplitMode:
			if unicode.IsUpper(r) || unicode.IsDigit(r) {
				s.writeWordRune(r)
			} else {
				s.finishWord(r)
			}
		}
	}
	s.finishWord(-1)
	return s.words
}

func (s *splitState) startWord(r rune) {
	s.curWord.WriteRune(r)
	if unicode.IsUpper(r) {
		s.mode = wordStartSplitMode
	} else {
		s.mode = wordSplitMode
	}
}

func (s *splitState) writeWordRune(r rune) {
	s.curWord.WriteRune(r)
	if unicode.IsUpper(r) {
		s.mode = acronymSplitMode
	} else {
		s.mode = wordSplitMode
	}
}

func (s *splitState) finishWord(r rune) {
	switch {
	case unicode.IsUpper(r):
		s.addWord(s.curWord.String())
		s.curWord.WriteRune(r)
		s.mode = wordStartSplitMode
	case unicode.IsLower(r):
		word := s.curWord.String()
		lastRune, lastRuneSize := utf8.DecodeLastRuneInString(word)
		s.addWord(word[:len(word)-lastRuneSize])
		s.curWord.WriteRune(lastRune)
		s.curWord.WriteRune(r)
		s.mode = wordSplitMode
	case unicode.IsTitle(r):
		s.addWord(s.curWord.String())
		s.curWord.WriteRune(r)
		s.mode = wordSplitMode
	default:
		s.addWord(s.curWord.String())
		s.mode = boundSplitMode
	}
}

func (s *splitState) addWord(word string) {
	if s.curWord.Len() > 0 {
		s.words = append(s.words, word)
		s.curWord.Reset()
	}
}
