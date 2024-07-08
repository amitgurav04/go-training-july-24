package language

type Language interface {
	GetGreeting() string
}
type EnglishLanguage struct{}
type HindiLanguage struct{}
type MarathiLanguage struct{}

func (e EnglishLanguage) GetGreeting() string {
	return "Hello!!!"
}
