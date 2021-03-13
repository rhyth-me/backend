package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/firebase/storage"
	"github.com/rhyth-me/backend/pkg/stripe"
)

func main() {
	_ = godotenv.Load(".env")

	// Init firebase
	auth.Init()
	firestore.Init()
	storage.Init()

	// Init stripe
	stripe.Init()

	// Init echo
	e := echo.New()
	e = initEchoSetting(e)

	// Init ControllerProps
	p := initControllerProps()

	apigen.Bootstrap(p, e, nil)

	fmt.Println("All endpoints are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s\n", r.Method, r.Path)
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
