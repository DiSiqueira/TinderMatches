package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/disiqueira/tindergo"
	"strconv"
)

type App struct {
	token   *string
	tinder  *tindergo.TinderGo
	profile tindergo.Profile
}

func (a *App) createValidateToken() error {
	return a.validateTokenFlag()
}

func (a *App) createTokenFlag() {
	a.token = flag.String("token", "", "Your Facebook Token.")
	flag.Parse()
}

func (a *App) validateTokenFlag() error {
	if *a.token == "" {
		return errors.New("You must provide a valid Facebook Token.")
	}
	return nil
}

func (a *App) createTinderAuthenticate() error {
	a.createTinderInstance()
	return a.authenticateTinderUsingToken()
}

func (a *App) createTinderInstance() {
	a.tinder = tindergo.New()
}

func (a *App) authenticateTinderUsingToken() error {
	return a.tinder.Authenticate(*a.token)
}

func (a *App) getTinderProfile() (err error) {
	a.profile, err = a.tinder.Profile()
	return
}

func (a *App) printBasicInfo() {
	a.printProfileName()
	a.printBlankLine()
}

func (a *App) printProfileName() {
	fmt.Println("You:")
	fmt.Println("Name:" + a.profile.Name)
}

func (a *App) printBlankLine() {
	fmt.Println("")
}

func (a *App) getPrintMatches(requests int) (err error) {
	matchList, err := a.getMatches(requests)
	if err != nil {
		return
	}
	a.printMatchesHeader()
	a.printMatches(matchList, requests)
	return
}

type matchList map[string]match

type match struct {
	user  tindergo.RecsCoreUser
	count int
}

func (a *App) getMatches(requests int) (matchAll matchList, err error) {
	matchAll = make(matchList)
	for j := 0; j <= requests; j++ {
		recs, err := a.tinder.RecsCore()
		if err != nil {
			return matchAll, err
		}

		fmt.Println("NEW REQUEST")

		for _, elem := range recs {
			fmt.Println(elem.ID)
			m := match{
				user:  elem,
				count: 1,
			}
			_, exist := matchAll[elem.ID]
			if exist {
				m.count += matchAll[elem.ID].count
			}
			matchAll[elem.ID] = m
		}
	}
	return
}

func (a *App) printMatchesHeader() {
	fmt.Printf("|%40s|%10s|\n", "Your Matches", "Accuracy")
}

func (a *App) printMatches(matches matchList,requests int) {
	for _, e := range matches {
		if e.count > 1 {
			fmt.Printf("|%40s|%10s|\n", e.user.Name, strconv.FormatFloat(float64((e.count*100)/requests), 'f', 0, 64)+"%")
		}
	}
}