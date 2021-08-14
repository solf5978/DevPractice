package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	message.Set(language.English, "%d items to do",
		plural.Selectf(1, "%d", "=0", "no items to do",
			plural.One, "one item to do",
			"<100", "%[1]d items to do",
			plural.Ohter, "lot of items to do",
		))
	message.Set(language.English, "The average is %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "The avg. is zero",
			"=1", "The avg. is one",
			plural.Other, "The avg. is %[1]f",
		))
	prt := message.NewPrinter(language.English)
	prt.Printf("%d items to do\n", 0)
	prt.Printf("%d items to do\n", 1)
	prt.Printf("%d items to do\n", 10)
	prt.Printf("%d items to do\n", 1000)

	prt.Printf("The avg. is %.2f\n", 0.8)
	prt.Printf("The avg. is %.2f\n", 1.0)
	prt.Printf("The avg. is %.2f\n", 10.0)
}
