package naming

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitName_baseCases(t *testing.T) {
	assert.Equal(t, []string{""}, SplitName(""))
	assert.Equal(t, []string{""}, SplitName("_"))
	assert.Equal(t, []string{"a"}, SplitName("a"))
	assert.Equal(t, []string{"á"}, SplitName("á"))
}

func TestSplitName_wordBoundaries(t *testing.T) {
	const (
		upper   = "A"
		lower   = "a"
		title   = "ǲ"
		digit   = "0"
		word    = "Aa"
		acronym = "AA"
		bound   = "_"
	)

	assert.Equal(t, []string{"AA"}, SplitName(upper+upper))
	assert.Equal(t, []string{"Aa"}, SplitName(upper+lower))
	assert.Equal(t, []string{"A", "ǲ"}, SplitName(upper+title))
	assert.Equal(t, []string{"A0"}, SplitName(upper+digit))
	assert.Equal(t, []string{"A", "Aa"}, SplitName(upper+word))
	assert.Equal(t, []string{"AAA"}, SplitName(upper+acronym))
	assert.Equal(t, []string{"A"}, SplitName(upper+bound))

	assert.Equal(t, []string{"a", "A"}, SplitName(lower+upper))
	assert.Equal(t, []string{"aa"}, SplitName(lower+lower))
	assert.Equal(t, []string{"a", "ǲ"}, SplitName(lower+title))
	assert.Equal(t, []string{"a0"}, SplitName(lower+digit))
	assert.Equal(t, []string{"a", "Aa"}, SplitName(lower+word))
	assert.Equal(t, []string{"a", "AA"}, SplitName(lower+acronym))
	assert.Equal(t, []string{"a"}, SplitName(lower+bound))

	assert.Equal(t, []string{"ǲ", "A"}, SplitName(title+upper))
	assert.Equal(t, []string{"ǲa"}, SplitName(title+lower))
	assert.Equal(t, []string{"ǲ", "ǲ"}, SplitName(title+title))
	assert.Equal(t, []string{"ǲ0"}, SplitName(title+digit))
	assert.Equal(t, []string{"ǲ", "Aa"}, SplitName(title+word))
	assert.Equal(t, []string{"ǲ", "AA"}, SplitName(title+acronym))
	assert.Equal(t, []string{"ǲ"}, SplitName(title+bound))

	assert.Equal(t, []string{"0", "A"}, SplitName(digit+upper))
	assert.Equal(t, []string{"0a"}, SplitName(digit+lower))
	assert.Equal(t, []string{"0", "ǲ"}, SplitName(digit+title))
	assert.Equal(t, []string{"00"}, SplitName(digit+digit))
	assert.Equal(t, []string{"0", "Aa"}, SplitName(digit+word))
	assert.Equal(t, []string{"0", "AA"}, SplitName(digit+acronym))
	assert.Equal(t, []string{"0"}, SplitName(digit+bound))

	assert.Equal(t, []string{"Aa", "A"}, SplitName(word+upper))
	assert.Equal(t, []string{"Aaa"}, SplitName(word+lower))
	assert.Equal(t, []string{"Aa", "ǲ"}, SplitName(word+title))
	assert.Equal(t, []string{"Aa0"}, SplitName(word+digit))
	assert.Equal(t, []string{"Aa", "Aa"}, SplitName(word+word))
	assert.Equal(t, []string{"Aa", "AA"}, SplitName(word+acronym))
	assert.Equal(t, []string{"Aa"}, SplitName(word+bound))

	assert.Equal(t, []string{"AAA"}, SplitName(acronym+upper))
	assert.Equal(t, []string{"A", "Aa"}, SplitName(acronym+lower))
	assert.Equal(t, []string{"AA", "ǲ"}, SplitName(acronym+title))
	assert.Equal(t, []string{"AA0"}, SplitName(acronym+digit))
	assert.Equal(t, []string{"AA", "Aa"}, SplitName(acronym+word))
	assert.Equal(t, []string{"AAAA"}, SplitName(acronym+acronym))
	assert.Equal(t, []string{"AA"}, SplitName(acronym+bound))

	assert.Equal(t, []string{"A"}, SplitName(bound+upper))
	assert.Equal(t, []string{"a"}, SplitName(bound+lower))
	assert.Equal(t, []string{"ǲ"}, SplitName(bound+title))
	assert.Equal(t, []string{"0"}, SplitName(bound+digit))
	assert.Equal(t, []string{"Aa"}, SplitName(bound+word))
	assert.Equal(t, []string{"AA"}, SplitName(bound+acronym))
	assert.Equal(t, []string{""}, SplitName(bound+bound))
}

func TestToTitle(t *testing.T) {
	assert.Equal(t, "Word", ToTitle("WORD"))
	assert.Equal(t, "Word", ToTitle("word"))
	assert.Equal(t, "Word", ToTitle("Word"))
	assert.Equal(t, "ǲord", ToTitle("ǲord"))
	assert.Equal(t, "ǲord", ToTitle("ǳord"))
}

func TestCreateNaming(t *testing.T) {
	mapA := func(_ string) string { return "a" }
	mapB := func(_ string) string { return "b" }

	assert.Equal(t, "aaa", CreateNaming(mapA, mapA, "")("word_word_word"))
	assert.Equal(t, "abb", CreateNaming(mapA, mapB, "")("word_word_word"))

	assert.Equal(t, "a_a_a", CreateNaming(mapA, mapA, "_")("word_word_word"))
	assert.Equal(t, "a_b_b", CreateNaming(mapA, mapB, "_")("word_word_word"))
}

func ExampleSplitName() {
	for _, name := range []string{
		"The quick brown fox jumps over the lazy dog.",
		"authorization_code",
		"MarshalJSON",
		"PKCS1v15DecryptOptions",
		"IPAddr",
	} {
		fmt.Println(SplitName(name))
	}
	// Output:
	// [The quick brown fox jumps over the lazy dog]
	// [authorization code]
	// [Marshal JSON]
	// [PKCS1v15 Decrypt Options]
	// [IP Addr]
}

func Example_namingConventions() {
	for _, formatted := range []string{
		Flat("Alice-WasBeginning"),
		Upper("toGetVery"),
		Mixed("tired-OF sitting"),
		UpperMixed("by__herSister"),
		Snake("_ONThe bank,"),
		UpperSnake("andOfHaving"),
		Kebab("nothingTo do:"),
	} {
		fmt.Println(formatted)
	}
	// Output:
	// alicewasbeginning
	// TOGETVERY
	// tiredOfSitting
	// ByHerSister
	// on_the_bank
	// AND_OF_HAVING
	// nothing-to-do
}
