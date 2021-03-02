package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/interfaces"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/pkg/firebase"
	"github.com/rhyth-me/backend/pkg/stripe"
)

func main() {
	godotenv.Load(".env")

	e := echo.New()

	auth := firebase.InitAuth()

	// Firebase auth - check login user
	// uid := c.(*model.CustomContext).UID
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := auth.VerifyIDToken(context.Background(), c.Request().Header.Get("X-Token"))
			if err != nil {
				cc := &model.CustomContext{Context: c}
				return next(cc)
			}

			if os.Getenv("STAGING") == "true" {
				token.UID = "STAGING_" + token.UID
			}

			cc := &model.CustomContext{Context: c, UID: token.UID}
			return next(cc)
		}
	})

	e.Debug = true
	e.Use(middleware.Recover())

	p := new(props.ControllerProps)

	// firebase init
	p.Firestore = firebase.InitFirestore()
	p.Stripe = stripe.Init()

	interfaces.Bootstrap(p, e, nil, os.Stdout)

	fmt.Println("All routes are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s: %s\n", r.Method, r.Path, r.Name)
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
