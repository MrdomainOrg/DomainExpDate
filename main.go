package main

import (
	"flag"
	"fmt"
	irnicDate "github.com/MrdomainOrg/DomainExpDate/irnicDate"
	"os"
	"time"
)

type DateFlag struct {
	*time.Time
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}
	var (
		dateFlag = flag.String("d", "", "Date in YYYY/MM/DD or YYYY-MM-DD format")
	)
	flag.Parse()
	if *dateFlag != "" {

	} else {
		fmt.Printf("Release Date for today is: %s\n", irnicDate.GetReleaseDate())
	}
}
