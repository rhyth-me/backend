package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhyth-me/backend/interfaces"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/pkg/authority"
	"github.com/rhyth-me/backend/pkg/firebase"
	"github.com/rhyth-me/backend/pkg/validator"
	"github.com/stripe/stripe-go/v72/client"
)

func initEchoSetting(e *echo.Echo) {
	e.Debug = true
	e.HideBanner = true
	e.Validator = validator.NewValidator()
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	// Firebase auth - check login user
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := authority.Identify(c, c.Request().Header.Get("X-Token"))
			return next(cc)
		}
	})

	// CORS config
	origins := []string{"https://rhyth.me"}
	if os.Getenv("STAGING") == "true" {
		origins = []string{"*"}
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func initControllerProps() *props.ControllerProps {

	p := new(props.ControllerProps)

	// firebase init
	p.Firestore = firebase.InitFirestore()
	p.Auth = firebase.InitAuth()

	// stripe init
	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_API_KEY"), nil)
	p.Stripe = sc

	return p
}

func main() {
	godotenv.Load(".env")

	e := echo.New()
	initEchoSetting(e)

	props := initControllerProps()
	interfaces.Bootstrap(props, e, nil, os.Stdout)

	fmt.Println("All endpoints are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s\n", r.Method, strings.TrimRight(r.Path, "/"))
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
