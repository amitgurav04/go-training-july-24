package language

import "fmt"

func (h HindiLanguage) GetGreeting() string {
	return "Namaste!!!"
}

func (m MarathiLanguage) GetGreeting() string {
	return "Namaskar!!!"
}

func PrintGreeting(l Language) {
	fmt.Println(l.GetGreeting())
}
