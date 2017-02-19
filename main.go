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
	checkError(err)

	profile, err := t.Profile()
	checkError(err)

	fmt.Println("Your Profile:")
	fmt.Println("Name: " + profile.Name)
	fmt.Println("")

	var allRecs map[string]tindergo.RecsCoreUser
	var countRecs map[string]int

	for j := 0; j <= 3; j++ {
		recs, err := t.RecsCore()
		checkError(err)

		for _, elem := range recs {
			_, exist := allRecs[elem.ID]
			if exist {
				countRecs[elem.ID] = countRecs[elem.ID] + 1
			} else {
				countRecs[elem.ID] = 1
				allRecs[elem.ID] = elem
			}
		}
	}

	for i, e := range allRecs {
		if countRecs[i] > 2 {
			fmt.Println(e.Name, countRecs[i], float64((countRecs[i]*100)/4))
		}
	}
}

// checkError Panic application if has an error returned.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
