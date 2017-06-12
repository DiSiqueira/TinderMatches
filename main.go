package main

var (
	app App
)

const (
	numRequests = 4
)

func init() {
	app = App{}
	app.createTokenFlag()
}

func main() {
	err := app.validateTokenFlag()
	checkError(err)

	err = app.createTinderAuthenticate()
	checkError(err)

	err = app.getTinderProfile()
	checkError(err)

	app.printBasicInfo()

	err = app.getPrintMatches(numRequests)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}