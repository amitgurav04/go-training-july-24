package main

import (
	"github.com/amitgurav04/interfaces/language"
)

func main() {
	e := language.EnglishLanguage{}
	h := language.HindiLanguage{}
	m := language.MarathiLanguage{}

	language.PrintGreeting(e)
	language.PrintGreeting(h)
	language.PrintGreeting(m)

}
