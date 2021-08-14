package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	w := tabwriter.NewWriter(os.Stdout, 5, 5, 5, ' ',
		tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJhon|Smith\t")
	w.Flush()
}
