package main

type Printer interface {
	Print(name string) string
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "en":
		return &EnglishPrinter{}
	case "cn":
		return &ChinesePrinter{}
	default:
		return &EnglishPrinter{}
	}
}

type EnglishPrinter struct {
}

func (p *EnglishPrinter) Print(name string) string {
	return "Hello " + name
}

type ChinesePrinter struct {
}

func (p *ChinesePrinter) Print(name string) string {
	return "你好 " + name
}

func main() {
	p := NewPrinter("cn")
	println(p.Print("world"))
}

// Output:
// 你好 world
