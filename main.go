package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/disiqueira/tindergo"
)

func main() {
	token := flag.String("token", "", "Your Facebook Token.")

	flag.Parse()

	if *token == "" {
		fmt.Println("You must provide a valid Facebook Token.")
		os.Exit(2)
	}

	t := tindergo.New()

	err := t.Authenticate(*token)

	profile, err := t.Profile()

	fmt.Println("Your Profile:")
	fmt.Println("Name: " + profile.Name)
	fmt.Println("")

	var allRecs []tindergo.RecsCoreUser

	for j := 0; j <= 3; j++ {
		recs, err := t.RecsCore()

		allRecs = append(allRecs, recs)
	}

}