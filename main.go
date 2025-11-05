package main

import (
	"Friendy/routes"
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// hooks.BindFriendsHooks(app)
	app.OnServe().BindFunc(routes.Routes)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
