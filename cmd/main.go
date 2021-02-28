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
	"github.com/rhyth-me/backend/pkg/stripe"
	"github.com/rhyth-me/backend/utils"
)

func main() {
	godotenv.Load(".env")

	e := echo.New()

	auth := utils.InitAuth()

	// Firebase auth - login user check
	// uid := c.(*model.CustomContext).AuthUser.UID
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := auth.VerifyIDToken(context.Background(), c.Request().Header.Get("X-Token"))
			if err != nil {
				cc := &model.CustomContext{c, model.AuthUser{}}
				return next(cc)
			}

			cc := &model.CustomContext{c, model.AuthUser{
				UID: token.UID,
			}}
			return next(cc)
		}
	})

	e.Debug = true
	e.Use(middleware.Recover())

	p := new(props.ControllerProps)

	// firebase init
	p.Firestore = utils.InitFirestore()
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
